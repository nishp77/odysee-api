package proxy

// Package proxy handles incoming JSON-RPC requests from UI client (lbry-desktop or any other),
// forwards them to the sdk and returns its response to the client.
// The purpose of it is to expose the SDK over a publicly accessible http interface,
// fixing aspects of it which normally would prevent SDK from being shared between multiple
// remote clients.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lbryio/lbrytv/app/auth"
	"github.com/lbryio/lbrytv/app/query"
	"github.com/lbryio/lbrytv/app/query/cache"
	"github.com/lbryio/lbrytv/app/rpcerrors"
	"github.com/lbryio/lbrytv/app/sdkrouter"
	"github.com/lbryio/lbrytv/internal/audit"
	"github.com/lbryio/lbrytv/internal/errors"
	"github.com/lbryio/lbrytv/internal/ip"
	"github.com/lbryio/lbrytv/internal/lbrynext"
	"github.com/lbryio/lbrytv/internal/metrics"
	"github.com/lbryio/lbrytv/internal/monitor"
	"github.com/lbryio/lbrytv/internal/responses"
	"github.com/lbryio/lbrytv/models"
	"github.com/sirupsen/logrus"

	"github.com/ybbus/jsonrpc"
)

var logger = monitor.NewModuleLogger("proxy")

const (
	orgOdysee  = "odysee"
	orgLbrytv  = "lbrytv"
	orgAndroid = "android"
	orgiOS     = "ios"
)

// observeFailure requires metrics.MeasureMiddleware middleware to be present on the request
func observeFailure(d float64, method, kind string) {
	metrics.ProxyE2ECallDurations.WithLabelValues(method).Observe(d)
	metrics.ProxyE2ECallFailedDurations.WithLabelValues(method, kind).Observe(d)
	metrics.ProxyE2ECallCounter.WithLabelValues(method).Inc()
	metrics.ProxyE2ECallFailedCounter.WithLabelValues(method, kind).Inc()
}

// observeSuccess requires metrics.MeasureMiddleware middleware to be present on the request
func observeSuccess(d float64, method string) {
	metrics.ProxyE2ECallDurations.WithLabelValues(method).Observe(d)
	metrics.ProxyE2ECallCounter.WithLabelValues(method).Inc()
}

func writeResponse(w http.ResponseWriter, b []byte) {
	w.Write(b)
}

// Handle forwards client JSON-RPC request to proxy.
func Handle(w http.ResponseWriter, r *http.Request) {
	responses.AddJSONContentType(w)
	origin := getDevice(r)

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, rpcerrors.NewJSONParseError(errors.Err("empty request body")).JSON())

		observeFailure(metrics.GetDuration(r), "", metrics.FailureKindClient)
		logger.Log().Debugf("empty request body")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, rpcerrors.NewJSONParseError(errors.Err("error reading request body")).JSON())

		observeFailure(metrics.GetDuration(r), "", metrics.FailureKindClient)
		logger.Log().Debugf("error reading request body: %v", err.Error())

		return
	}

	var rpcReq *jsonrpc.RPCRequest
	err = json.Unmarshal(body, &rpcReq)
	if err != nil {
		writeResponse(w, rpcerrors.NewJSONParseError(err).JSON())

		observeFailure(metrics.GetDuration(r), "", metrics.FailureKindClientJSON)
		logger.Log().Debugf("error unmarshaling request body: %v", err)

		return
	}

	logger.Log().Tracef("call to method %s", rpcReq.Method)

	user, err := auth.FromRequest(r)
	if query.MethodRequiresWallet(rpcReq.Method, rpcReq.Params) {
		authErr := GetAuthError(user, err)
		if authErr != nil {
			writeResponse(w, rpcerrors.ErrorToJSON(authErr))
			observeFailure(metrics.GetDuration(r), rpcReq.Method, metrics.FailureKindAuth)

			return
		}
	}

	var userID int
	if query.MethodAcceptsWallet(rpcReq.Method) && user != nil {
		userID = user.ID
	}

	sdkAddress := sdkrouter.GetSDKAddress(user)
	if sdkAddress == "" {
		rt := sdkrouter.FromRequest(r)
		sdkAddress = rt.RandomServer().Address
	}

	var qCache *cache.Cache
	if cache.IsOnRequest(r) {
		qCache = cache.FromRequest(r)
	}
	c := query.NewCaller(sdkAddress, userID)

	remoteIP := ip.FromRequest(r)
	// Logging remote IP with query
	c.AddPostflightHook("wallet_", func(_ *query.Caller, hctx *query.HookContext) (*jsonrpc.RPCResponse, error) {
		hctx.AddLogField("remote_ip", remoteIP)
		return nil, nil
	}, "")
	c.AddPostflightHook(query.MethodWalletSend, func(_ *query.Caller, hctx *query.HookContext) (*jsonrpc.RPCResponse, error) {
		audit.LogQuery(userID, remoteIP, query.MethodWalletSend, body)
		return nil, nil
	}, "")

	lbrynext.InstallHooks(c)
	c.Cache = qCache

	rpcRes, err := c.Call(rpcReq)
	metrics.ProxyCallDurations.WithLabelValues(rpcReq.Method, c.Endpoint(), origin).Observe(c.Duration)
	metrics.ProxyCallCounter.WithLabelValues(rpcReq.Method, c.Endpoint(), origin).Inc()

	if err != nil {
		monitor.ErrorToSentry(err, map[string]string{"request": fmt.Sprintf("%+v", rpcReq), "response": fmt.Sprintf("%+v", rpcRes)})
		writeResponse(w, rpcerrors.ToJSON(err))

		logger.Log().Errorf("error calling lbrynet: %v, request: %+v", err, rpcReq)
		observeFailure(metrics.GetDuration(r), rpcReq.Method, metrics.FailureKindNet)
		metrics.ProxyCallFailedDurations.WithLabelValues(rpcReq.Method, c.Endpoint(), origin, metrics.FailureKindNet).Observe(c.Duration)
		metrics.ProxyCallFailedCounter.WithLabelValues(rpcReq.Method, c.Endpoint(), origin, metrics.FailureKindNet).Inc()
		return
	}

	serialized, err := responses.JSONRPCSerialize(rpcRes)
	if err != nil {
		monitor.ErrorToSentry(err)

		writeResponse(w, rpcerrors.NewInternalError(err).JSON())

		logger.Log().Errorf("error marshaling response: %v", err)
		observeFailure(metrics.GetDuration(r), rpcReq.Method, metrics.FailureKindRPCJSON)

		return
	}

	if rpcRes.Error != nil {
		observeFailure(metrics.GetDuration(r), rpcReq.Method, metrics.FailureKindRPC)
		metrics.ProxyCallFailedDurations.WithLabelValues(rpcReq.Method, c.Endpoint(), origin, metrics.FailureKindRPC).Observe(c.Duration)
		metrics.ProxyCallFailedCounter.WithLabelValues(rpcReq.Method, c.Endpoint(), origin, metrics.FailureKindRPC).Inc()

		logger.WithFields(logrus.Fields{
			"method":   rpcReq.Method,
			"endpoint": sdkAddress,
			"response": rpcRes.Error,
		}).Errorf("proxy handler got rpc error: %v", rpcRes.Error)
	} else {
		observeSuccess(metrics.GetDuration(r), rpcReq.Method)
	}

	writeResponse(w, serialized)
}

func GetAuthError(user *models.User, err error) error {
	if err == nil && user != nil {
		return nil
	}

	if errors.Is(err, auth.ErrNoAuthInfo) {
		return rpcerrors.NewAuthRequiredError()
	} else if err != nil {
		return rpcerrors.NewForbiddenError(err)
	} else if user == nil {
		return rpcerrors.NewForbiddenError(errors.Err("must authenticate"))
	}

	return errors.Err("unknown auth error")
}

func getDevice(r *http.Request) string {
	rf := r.Header.Get("referer")
	ua := r.Header.Get("user-agent")
	if strings.HasSuffix(rf, "odysee.com/") {
		return orgOdysee
	}
	if strings.HasSuffix(rf, "lbry.tv/") {
		return orgLbrytv
	}
	if strings.Contains(ua, "okhttp") {
		return orgAndroid
	}
	if strings.Contains(strings.ToLower(ua), "odysee") {
		return orgiOS
	}
	return ""
}
