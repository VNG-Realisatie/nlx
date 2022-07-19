package http

import (
	"net/http"

	"go.nlx.io/nlx/common/httperrors"
)

func WriteError(w http.ResponseWriter, location httperrors.Location, code httperrors.Code, message string) {
	httperrors.WriteError(w, httperrors.Outway, location, code, message)
}
