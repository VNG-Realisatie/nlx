// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package i18n

type I18n interface {
	Translate(key string) string
	Locale() string
}
