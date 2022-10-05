// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package directory

import (
	"context"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func (h *DirectoryService) GetTermsOfService(_ context.Context, _ *directoryapi.GetTermsOfServiceRequest) (*directoryapi.GetTermsOfServiceResponse, error) {
	h.logger.Info("rpc request GetTermsOfService")

	return &directoryapi.GetTermsOfServiceResponse{
		Enabled: h.termsOfServiceURL != "",
		Url:     h.termsOfServiceURL,
	}, nil
}
