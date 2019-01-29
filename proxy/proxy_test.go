package proxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/lbryio/lbryweb.go/config"
	"github.com/stretchr/testify/assert"
	"github.com/ybbus/jsonrpc"
)

func TestForwardCallWithHTTPError(t *testing.T) {
	defaultDaemonURL := DaemonURL
	DaemonURL = "http://localhost:59999"
	query := jsonrpc.NewRequest("account_balance")
	queryBody, _ := json.Marshal(query)
	response, err := ForwardCall(queryBody)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "rpc call account_balance() on http://localhost:59999"))
	assert.True(t, strings.HasSuffix(err.Error(), "connect: connection refused"))
	DaemonURL = defaultDaemonURL
}

func TestForwardCallWithLbrynetError(t *testing.T) {
	var response jsonrpc.RPCResponse
	query := jsonrpc.NewRequest("crazy_method")
	queryBody, _ := json.Marshal(query)
	rawResponse, err := ForwardCall(queryBody)
	json.Unmarshal(rawResponse, &response)
	assert.Nil(t, err)
	assert.NotNil(t, response.Error)
	// TODO: Uncomment after lbrynet 0.31 release
	// assert.Equal(t, "Invalid method requested: crazy_method.", response.Error.Message)
	assert.Equal(t, "Method Not Found", response.Error.Message)
}

// A jsonrpc server that responds to every query with an error
func launchGrumpyServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(jsonrpc.RPCResponse{Error: &jsonrpc.RPCError{Message: "your ways are wrong"}})
		w.Write(response)
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:59999", nil))
}

func TestForwardCallWithClientError(t *testing.T) {
	var response jsonrpc.RPCResponse

	defaultDaemonURL := DaemonURL
	DaemonURL = "http://localhost:59999"

	go launchGrumpyServer()

	query := jsonrpc.NewRequest("get")
	queryBody, _ := json.Marshal(query)
	rawResponse, err := ForwardCall(queryBody)
	json.Unmarshal(rawResponse, &response)
	assert.Nil(t, err)
	assert.NotNil(t, response.Error)
	assert.Equal(t, "your ways are wrong", response.Error.Message)

	DaemonURL = defaultDaemonURL
}

func TestForwardCall(t *testing.T) {
	var err error
	var query *jsonrpc.RPCRequest
	var response jsonrpc.RPCResponse
	var rawResponse []byte
	var result map[string]interface{}
	var queryBody []byte

	_, err = ForwardCall([]byte("yo"))
	assert.NotNil(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "client json parse error: invalid character 'y'"))

	query = &jsonrpc.RPCRequest{Method: "account_balance", ID: 123}
	queryBody, _ = json.Marshal(query)
	rawResponse, err = ForwardCall(queryBody)
	json.Unmarshal(rawResponse, &response)
	if err != nil {
		t.Errorf("failed with an unexpected error: %v", err)
		return
	} else if response.Error != nil {
		t.Errorf("daemon unexpectedly errored: %v", response.Error.Message)
	} else if response.Result != "0.0" {
		t.Errorf("unexpected result from daemon: %q", response.Result)
	}
	if response.ID != 123 {
		t.Errorf("daemon response ID mismatch: %v != 123", response.ID)
	}

	query = jsonrpc.NewRequest("get", map[string]string{"uri": "what"})
	queryBody, _ = json.Marshal(query)
	rawResponse, err = ForwardCall(queryBody)
	json.Unmarshal(rawResponse, &response)
	if err != nil {
		t.Errorf("failed with an unexpected error: %v", err)
	} else if response.Error != nil {
		t.Errorf("daemon errored: %v", response.Error.Message)
	}

	response.GetObject(&result)
	expectedPath := fmt.Sprintf(
		"%s%s/%s", config.Settings.GetString("BaseContentURL"), "what", result["outpoint"])
	assert.Equal(t, expectedPath, result["download_path"])

	outpoint := result["outpoint"]
	query = jsonrpc.NewRequest("file_list", map[string]string{"outpoint": outpoint.(string)})
	queryBody, _ = json.Marshal(query)
	rawResponse, err = ForwardCall(queryBody)
	json.Unmarshal(rawResponse, &response)
	var resultArray []map[string]interface{}
	response.GetObject(&resultArray)
	assert.Nil(t, err)
	assert.Nil(t, response.Error)
	if len(resultArray) == 0 {
		t.Errorf("not enough results, daemon responded with %v", response.Result)
		return
	}

	expectedPath = fmt.Sprintf(
		"%s%s/%s/%s", config.Settings.GetString("BaseContentURL"), "outpoints",
		outpoint, resultArray[0]["file_name"])
	if resultArray[0]["download_path"] != expectedPath {
		t.Errorf("expected result.0.download_path to be %v but got %v", expectedPath, resultArray[0]["download_path"])
	}
}
