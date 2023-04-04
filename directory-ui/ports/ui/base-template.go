// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"fmt"
	"html/template"
	"log"
	"strings"
)

func templateWithSVGHelper() *template.Template {
	svgTemplates := template.Must(template.ParseFS(
		tplFolder,
		"templates/svg/icon-close.svg",
		"templates/svg/icon-external-link.svg",
		"templates/svg/news.svg",
		"templates/svg/icon-search.svg",
		"templates/svg/state-down.svg",
		"templates/svg/state-up.svg",
		"templates/svg/icon-home.svg",
		"templates/svg/icon-participants.svg",
		"templates/svg/icon-support.svg",
		"templates/svg/icon-chevron-right.svg",
		"templates/svg/nlx-logo.svg",
		"templates/svg/vng-logo.svg",
	))

	funcMap := template.FuncMap{
		"svg": func(name, class string) template.HTML {
			svgData := new(strings.Builder)
			err := svgTemplates.ExecuteTemplate(svgData, fmt.Sprintf("%s.svg", name), struct {
				Class string
			}{
				Class: class,
			})
			if err != nil {
				log.Printf("unexpected error: %s\n", err)
				return "INVALID SVG NAME PROVIDED"
			}

			// nolint:gosec // we are the owners of the HTML input, so the Cross-site Scripting is not applicable
			return template.HTML(svgData.String())
		},
	}

	return template.
		New("").
		Funcs(funcMap)
}
