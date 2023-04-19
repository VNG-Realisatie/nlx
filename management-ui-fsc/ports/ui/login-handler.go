// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

func (s *Server) loginHandler(w http.ResponseWriter, _ *http.Request) {
	page := loginPage{
		BasePage:            s.basePage,
		IsUserAuthenticated: false,
	}

	err := page.render(w)
	if err != nil {
		s.logger.Error("could not render login page", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
