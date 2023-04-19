// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

type inwaysPage struct {
	*BasePage
	BaseAuthenticatedPage

	AmountOfInways  uint
	AmountOfOutways uint
	Inways          []*inwaysPageInway
}

type inwaysPageInway struct {
	Name        string
	Hostname    string
	SelfAddress string
	Services    uint
	Version     string
}

func (p *inwaysPage) render(w http.ResponseWriter) error {
	baseTemplate := p.TemplateWithHelpers()

	t, err := baseTemplate.
		ParseFS(
			tplFolder,
			"templates/base-authenticated.html",
			"templates/inways.html",
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
