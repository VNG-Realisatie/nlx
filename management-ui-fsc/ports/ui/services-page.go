// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

type servicesPage struct {
	*BasePage
	BaseAuthenticatedPage

	AmountOfServices uint
	Services         []*servicesPageService
}

type servicesPageService struct {
	Name string
}

func (p *servicesPage) render(w http.ResponseWriter) error {
	baseTemplate := p.TemplateWithHelpers()

	t, err := baseTemplate.
		ParseFS(
			tplFolder,
			"templates/base-authenticated.html",
			"templates/services.html",
		)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "base-authenticated.html", p)
	if err != nil {
		return err
	}

	return nil
}
