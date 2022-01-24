// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func (h *DirectoryService) GetTermsOfService(ctx context.Context, _ *emptypb.Empty) (*directoryapi.GetTermsOfServiceResponse, error) {
	h.logger.Info("rpc request GetTermsOfService")

	return &directoryapi.GetTermsOfServiceResponse{
		Enabled: h.termsOfServiceURL != "",
		Url:     h.termsOfServiceURL,
	}, nil
}
