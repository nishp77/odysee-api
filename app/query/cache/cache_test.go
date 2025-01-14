package cache

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc"
)

func TestCache(t *testing.T) {
	var (
		res jsonrpc.RPCResponse
		req jsonrpc.RPCRequest
	)

	rreq := `{"jsonrpc":"2.0","method":"resolve","params":{"urls":["one", "two", "three"]},"id":1555013448981}`
	err := json.Unmarshal([]byte(rreq), &req)
	if err != nil {
		t.Fatal(err)
	}

	absPath, _ := filepath.Abs("./testdata/resolve_response.json")
	rres, err := ioutil.ReadFile(absPath)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(rres, &res)
	require.NoError(t, err)

	cfg := DefaultConfig()
	cfg.ristrettoMetrics = true
	c, err := New(cfg)
	require.NoError(t, err)

	retrievals := 0
	wg := &sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			cached, err := c.Retrieve("resolve", req.Params, func() (interface{}, error) {
				retrievals++
				time.Sleep(500 * time.Millisecond)
				return res, nil
			})
			assert.NoError(t, err)
			assert.Equal(t, res, cached)
			wg.Done()
		}()
	}
	wg.Wait()

	assert.EqualValues(t, 1, c.count())
	assert.EqualValues(t, 1, retrievals)
}
