// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"testing"

	"go.nlx.io/nlx/common/delegation"

	"github.com/stretchr/testify/assert"
)

func TestStripHeadersPlugin(t *testing.T) {
	headers := []string{
		"X-NLX-Requester-User",
		"X-NLX-Requester-Claims",
		delegation.HTTPHeaderDelegator,
		delegation.HTTPHeaderOrderReference,
		"X-NLX-Request-Subject-Identifier",
		"X-NLX-Request-Application-Id",
		"X-NLX-Request-User-Id",
		"X-NLX-Logrecord-ID",
		"X-NLX-Request-Data-Subject",
		"Proxy-Authorization",
	}

	unsafeHeaders := []string{
		"X-NLX-Requester-User",
		"X-NLX-Request-User-Id",
		"X-NLX-Requester-Claims",
		delegation.HTTPHeaderDelegator,
		delegation.HTTPHeaderOrderReference,
		"X-NLX-Request-Application-Id",
		"X-NLX-Request-Subject-Identifier",
	}

	safeHeaders := []string{
		"X-NLX-Logrecord-ID",
		"X-NLX-Request-Data-Subject",
	}

	tests := map[string]struct {
		name                 string
		receiverOrganization string
		expectHeaders        []string
		disallowedHeaders    []string
	}{
		"Different Organization": {
			receiverOrganization: "00000000000000000002",
			expectHeaders:        safeHeaders,
			disallowedHeaders:    unsafeHeaders,
		},
		"Same Organization": {
			receiverOrganization: "00000000000000000001",
			expectHeaders:        append(safeHeaders, unsafeHeaders...),
			disallowedHeaders:    nil,
		},
		"Do not pass Proxy-Authorization": {
			receiverOrganization: "00000000000000000002",
			expectHeaders:        nil,
			disallowedHeaders:    []string{"Proxy-Authorization"},
		},
		"Never pass Proxy-Authorization": {
			receiverOrganization: "00000000000000000001",
			expectHeaders:        nil,
			disallowedHeaders:    []string{"Proxy-Authorization"},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&Destination{
				OrganizationSerialNumber: "00000000000000000001",
			})

			for _, header := range headers {
				context.Request.Header.Add(header, header)
			}

			plugin := NewStripHeadersPlugin(tt.receiverOrganization)

			err := plugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			if tt.expectHeaders != nil {
				for _, header := range tt.expectHeaders {
					assert.Equal(t, header, context.Request.Header.Get(header))
				}
			}

			if tt.disallowedHeaders != nil {
				for _, header := range tt.disallowedHeaders {
					assert.Equal(t, "", context.Request.Header.Get(header))
				}
			}
		})
	}
}
