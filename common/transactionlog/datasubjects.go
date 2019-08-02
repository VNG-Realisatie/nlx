// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var regExpKey = regexp.MustCompile("^[a-z0-9]{4,100}$")
var regExpValue = regexp.MustCompile("^[a-zA-Z0-9.-]{1,100}$")

// ParseDataSubjectHeader parses a data subject header value and returns a key/value map, or an error.
func ParseDataSubjectHeader(header string) (map[string]string, error) {
	if header == "" {
		return make(map[string]string), nil
	}
	parts := strings.Split(header, ",")
	dataSubjects := make(map[string]string)
	for i, part := range parts {
		kv := strings.Split(part, "=")
		if len(kv) != 2 {
			return nil, errors.Errorf("invalid datasubject in header, %dth subject is not a correct key=value format", i+1)
		}
		key, value := strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1])
		if !regExpKey.MatchString(key) {
			return nil, errors.Errorf("invalid datasubject in header, %dth subject key '%s' is not valid", i+1, key)
		}

		if !regExpValue.MatchString(value) {
			return nil, errors.Errorf("invalid datasubject in header, %dth subject value '%s' is not valid", i+1, value)
		}
		dataSubjects[key] = value
	}
	return dataSubjects, nil
}
