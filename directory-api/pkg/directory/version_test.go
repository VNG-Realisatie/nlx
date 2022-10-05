// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package directory_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func TestGetVersion(t *testing.T) {
	tests := map[string]struct {
		version      string
		wantResponse *directoryapi.GetVersionResponse
	}{
		"happy_flow_version_not_set": {
			version:      "",
			wantResponse: &directoryapi.GetVersionResponse{Version: "unknown"},
		},
		"happy_flow": {
			version:      testNlxVersion128,
			wantResponse: &directoryapi.GetVersionResponse{Version: testNlxVersion128},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _ := newService(t, tt.version, "", &testClock{
				timeToReturn: time.Now(),
			})

			got, err := service.GetVersion(context.Background(), &directoryapi.GetVersionRequest{})
			assert.NoError(t, err)

			assert.Equal(t, tt.wantResponse, got)
		})
	}
}
