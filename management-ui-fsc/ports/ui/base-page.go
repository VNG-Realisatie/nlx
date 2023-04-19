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

	"go.nlx.io/nlx/management-ui-fsc/ports/ui/i18n"
)

const PathServicesPage = "/services"

type BasePage struct {
	staticPath       string
	svgTemplates     *template.Template
	assetsMap        assets
	I18n             i18n.I18n
	PathServicesPage string
}

func NewBasePage(staticPath string, translations i18n.I18n) (*BasePage, error) {
	if staticPath == "" {
		return nil, fmt.Errorf("static path is required")
	}

	svgTemplates := template.Must(template.ParseFS(
		tplFolder,
		"templates/svg/icon-inways-and-outways.svg",
		"templates/svg/icon-services.svg",
		"templates/svg/icon-error.svg",
		"templates/svg/icon-plus.svg",
		"templates/svg/icon-chevron-down.svg",
		"templates/svg/icon-chevron-right.svg",
		"templates/svg/icon-settings.svg",
		"templates/svg/icon-inway.svg",
		"templates/svg/icon-outway.svg",
		"templates/svg/nlx-logo.svg",
		"templates/svg/nlx-management-logo.svg",
	))

	assetsMap, err := getAssetsMap(staticPath, "parcel-manifest.json")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to setup assets map")
	}

	return &BasePage{
		staticPath:       staticPath,
		assetsMap:        assetsMap,
		svgTemplates:     svgTemplates,
		I18n:             translations,
		PathServicesPage: PathServicesPage,
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
		"i18n": func(key string) template.HTML {
			// nolint:gosec // we are the owners of the HTML input, so the Cross-site Scripting is not applicable
			return template.HTML(b.I18n.Translate(key))
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
