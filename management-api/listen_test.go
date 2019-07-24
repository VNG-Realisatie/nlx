package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/stretchr/testify/assert"
)

func TestManagementAPIListen(t *testing.T) {
	api := API{
		mux: runtime.NewServeMux(),
	}

	mockServer := httptest.NewServer(api)
	defer mockServer.Close()

	client := http.Client{}

	req, err := http.NewRequest("GET", mockServer.URL, nil)
	if err != nil {
		t.Error("can not construct request", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error("can not perform request", err)
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
