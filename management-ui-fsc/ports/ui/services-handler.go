// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

func (s *Server) servicesHandler(w http.ResponseWriter, _ *http.Request) {
	page := servicesPage{
		BasePage: s.basePage,
		BaseAuthenticatedPage: BaseAuthenticatedPage{
			PrimaryNavigationActivePath: PathServicesPage,
			Title:                       s.i18n.Translate("Services"),
			Description:                 "",
			Username:                    "admin",
			OrganizationName:            "Gemeente Stijns",
			OrganizationSerialNumber:    "12345678901234567890",
		},
		AmountOfServices: 0,
		Services: []*servicesPageService{
			{
				Name: "Service A",
			},
		},
	}

	err := page.render(w)
	if err != nil {
		s.logger.Error("could not render services page", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
