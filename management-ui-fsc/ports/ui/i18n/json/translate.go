// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package jsoni18n

func (s *I18n) Translate(key string) string {
	val, ok := s.translations[key]
	if !ok {
		return key
	}

	return val
}
