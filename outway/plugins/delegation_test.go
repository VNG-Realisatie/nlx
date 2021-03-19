// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelegationPlugin(t *testing.T) {
	tests := map[string]struct {
		wantErr    bool
		setHeaders func(*http.Request)
	}{
		"missing_order_reference_returns_an_errors": {
			wantErr: true,
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "TestOrg")
			},
		},

		"missing_delegator_returns_an_errors": {
			wantErr: true,
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-OrderReference", "test-ref-123")
			},
		},

		"required_headers_returns_ok": {
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "TestOrg")
				r.Header.Add("X-NLX-Request-OrderReference", "test-ref-123")
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&Destination{})

			tt.setHeaders(context.Request)

			plugin := NewDelegationPlugin()

			err := plugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()

			defer response.Body.Close()

			if tt.wantErr {
				assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
			} else {
				assert.Equal(t, http.StatusOK, response.StatusCode)
				assert.Equal(t, "TestOrg", context.LogData["delegator"])
				assert.Equal(t, "test-ref-123", context.LogData["orderReference"])
			}
		})
	}
}
