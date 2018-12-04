package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewEncrypterServeMux(t *testing.T) {
	var logBuff bytes.Buffer
	logger := log.New(&logBuff, "", 0)
	srb := storeRequestBody{
		Id:[]byte("qwert"),
		Payload:[]byte("yuiop"),
	}
	srv := httptest.NewServer(NewEncrypterServeMux(logger))

	var key []byte
	t.Run("store", func(t *testing.T) {
		rb, err := json.Marshal(srb)
		assert.NoError(t, err)
		assert.NotNil(t, rb)

		reqBod := bytes.NewBuffer(rb)
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/store", srv.URL), reqBod)
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)

		key, err = ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
	})

	var payload []byte
	t.Run("retrieve", func(t *testing.T) {
		rrb := retrieveRequestBody{
			Id:srb.Id,
			Key:key,
		}

		rb, err := json.Marshal(rrb)
		assert.NoError(t, err)
		assert.NotNil(t, rb)

		reqBod := bytes.NewBuffer(rb)
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/retrieve", srv.URL), reqBod)
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)

		payload, err = ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
	})
	
	assert.Equal(t, srb.Payload, payload, logBuff.String())
}
