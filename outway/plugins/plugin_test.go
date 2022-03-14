// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"net/http"
	"net/http/httptest"

	"go.uber.org/zap"
)

func fakeContext(dest *Destination) *Context {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", http.NoBody)

	return &Context{
		Destination: dest,
		Request:     request,
		Response:    recorder,
		Logger:      zap.NewNop(),
		LogData:     map[string]string{},
	}
}

func nopServeFunc(context *Context) error {
	return nil
}
