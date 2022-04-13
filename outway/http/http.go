package http

import (
	"fmt"
	"net/http"

	"go.nlx.io/nlx/common/httperrors"
)

func WriteError(w http.ResponseWriter, message string) {
	http.Error(w, fmt.Sprintf("nlx-outway: %s", message), httperrors.StatusNLXNetworkError)
}
