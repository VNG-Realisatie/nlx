// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package strings

import (
	"regexp"
	"strings"
)

func ToSnakeCase(s string) string {
	capitalsRegex := regexp.MustCompile("([A-Z]+)")
	duplicateUnderscoreRegex := regexp.MustCompile("_+")

	replaceCapitals := capitalsRegex.ReplaceAllString(s, "_$1")
	replaceDuplicateUnderscores := duplicateUnderscoreRegex.ReplaceAllString(replaceCapitals, "_")

	lower := strings.ToLower(replaceDuplicateUnderscores)
	trimmed := strings.Trim(lower, "_")

	return trimmed
}
