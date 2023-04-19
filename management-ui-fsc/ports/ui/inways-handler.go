// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

func (s *Server) inwaysHandler(w http.ResponseWriter, _ *http.Request) {
	inways := []*inwaysPageInway{
		{
			Name:        "gemeente-stijns-nlx-inway",
			Hostname:    "gemeente-stijns-nlx-inway-7f8ccfb4dd-9qpgs",
			SelfAddress: "gemeente-stijns-nlx-inway:443",
			Services:    0,
			Version:     "0.147.2-acc-a631ad6e",
		},
	}

	page := inwaysPage{
		BasePage: s.basePage,
		BaseAuthenticatedPage: BaseAuthenticatedPage{
			PrimaryNavigationActivePath: "/inways-and-outways",
			Title:                       s.i18n.Translate("Inways and Outways"),
			Description:                 s.i18n.Translate("Gateways to provide (Inways) or consume (Outways) services."),
			Username:                    "admin",
			OrganizationName:            "Gemeente Stijns",
			OrganizationSerialNumber:    "12345678901234567890",
		},
		AmountOfInways:  uint(len(inways)),
		AmountOfOutways: 0,
		Inways:          inways,
	}

	err := page.render(w)
	if err != nil {
		s.logger.Error("could not render inways page", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
