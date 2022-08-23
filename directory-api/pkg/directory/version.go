// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package directory

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func (h *DirectoryService) GetVersion(context.Context, *emptypb.Empty) (*directoryapi.GetVersionResponse, error) {
	h.logger.Info("rpc request GetVersion")

	return &directoryapi.GetVersionResponse{
		Version: h.version,
	}, nil
}
