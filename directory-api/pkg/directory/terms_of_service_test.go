// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func TestGetTermsOfService(t *testing.T) {
	tests := map[string]struct {
		termsOfServiceURL string
		expectedResponse  *directoryapi.GetTermsOfServiceResponse
		expectedError     error
	}{
		"happy_flow_enabled": {
			termsOfServiceURL: "https://mock.mock/terms-of-service",
			expectedResponse:  &directoryapi.GetTermsOfServiceResponse{Enabled: true, Url: "https://mock.mock/terms-of-service"},
			expectedError:     nil,
		},
		"happy_flow_disabled": {
			termsOfServiceURL: "",
			expectedResponse:  &directoryapi.GetTermsOfServiceResponse{Enabled: false},
			expectedError:     nil,
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _ := newService(t, tt.termsOfServiceURL)

			got, err := service.GetTermsOfService(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
