// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

type loginPage struct {
	*BasePage

	IsUserAuthenticated bool
}

func (p *loginPage) render(w http.ResponseWriter) error {
	baseTemplate := p.TemplateWithHelpers()

	t, err := baseTemplate.
		ParseFS(
			tplFolder,
			"templates/base-unauthenticated.html",
			"templates/login.html",
		)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "base-unauthenticated.html", p)
	if err != nil {
		return err
	}

	return nil
}
