// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package http

import (
	"net/http"

	"go.nlx.io/nlx/common/httperrors"
)

func WriteError(w http.ResponseWriter, location httperrors.Location, nlxErr *httperrors.NLXNetworkError) {
	nlxErr.Location = location
	nlxErr.Source = httperrors.Inway
	httperrors.WriteError(w, nlxErr)
}
