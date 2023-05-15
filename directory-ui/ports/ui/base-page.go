// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

type BasePage struct {
	staticPath   string
	svgTemplates *template.Template
	assetsMap    assets
}

func NewBasePage(staticPath string) (*BasePage, error) {
	if staticPath == "" {
		return nil, fmt.Errorf("static path is required")
	}

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

	assetsMap, err := getAssetsMap(staticPath, "parcel-manifest.json")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to setup assets map")
	}

	return &BasePage{
		staticPath:   staticPath,
		assetsMap:    assetsMap,
		svgTemplates: svgTemplates,
	}, nil
}

func (b *BasePage) TemplateWithHelpers() *template.Template {
	funcMap := template.FuncMap{
		"svg": func(name, class string) template.HTML {
			svgData := new(strings.Builder)
			err := b.svgTemplates.ExecuteTemplate(svgData, fmt.Sprintf("%s.svg", name), struct {
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
		"asset": func(filePath string) template.HTML {
			result := filePath

			val, ok := b.assetsMap[filePath]
			if ok {
				result = val
			}

			// nolint:gosec // we are the owners of the HTML input, so the Cross-site Scripting is not applicable
			return template.HTML(result)
		},
	}

	return template.
		New("").
		Funcs(funcMap)
}

type assets map[string]string

func getAssetsMap(staticPath, manifestFileName string) (assets, error) {
	result := assets{}

	content, err := os.ReadFile(path.Join(staticPath, manifestFileName))
	if err != nil {
		return result, errors.Wrapf(err, "failed to read manifest file '%s'", manifestFileName)
	}

	var raw map[string]interface{}

	err = json.Unmarshal(content, &raw)
	if err != nil {
		return result, errors.Wrapf(err, "error while unmarshalling translation file '%s'", manifestFileName)
	}

	for key, value := range raw {
		result[key] = fmt.Sprintf("%s", value)
	}

	return result, nil
}
