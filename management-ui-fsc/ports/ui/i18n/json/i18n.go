// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package jsoni18n

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"go.nlx.io/nlx/management-ui-fsc/ports/ui/i18n"
)

const LocaleEn = "en"
const LocaleNl = "nl"

type I18n struct {
	locale       string
	translations map[string]string
}

//go:embed source/**
var i18nFolder embed.FS

func New(locale string) (i18n.I18n, error) {
	if !slices.Contains([]string{LocaleEn, LocaleNl}, locale) {
		return nil, fmt.Errorf("invalid locale '%s' provided", locale)
	}

	filename := fmt.Sprintf("source/%s.json", locale)

	content, err := i18nFolder.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading %s", filename)
	}

	var raw map[string]interface{}

	err = json.Unmarshal(content, &raw)
	if err != nil {
		return nil, errors.Wrapf(err, "error while unmarshalling translation file %s", content)
	}

	translations := map[string]string{}

	for key, value := range raw {
		translations[key] = fmt.Sprintf("%s", value)
	}

	return &I18n{
		translations: translations,
		locale:       locale,
	}, nil
}
