// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package directory

import (
	"context"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func (h *DirectoryService) GetVersion(context.Context, *directoryapi.GetVersionRequest) (*directoryapi.GetVersionResponse, error) {
	h.logger.Info("rpc request GetVersion")

	return &directoryapi.GetVersionResponse{
		Version: h.version,
	}, nil
}
