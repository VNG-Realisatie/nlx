package transactionlog

import (
	"strings"

	"github.com/pkg/errors"
)

// ParseDataSubjectHeader parses a data subject header value and returns a key/value map, or an error.
func ParseDataSubjectHeader(header string) (map[string]string, error) {
	parts := strings.Split(header, ",")
	dataSubjects := make(map[string]string)
	for i, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, errors.Errorf("invalid datasubject in header, %dth subject is empty", i+1)
		}
		//TODO: Add regex to validate key and value
		// key: [a-z-]{4,100}
		// value: all printable characters? {1,100}?
		kv := strings.Split(part, "=")
		if len(kv) != 2 {
			return nil, errors.Errorf("invalid datasubject in header, %dth subject is not a correct key=value format", i+1)
		}
		key, value := kv[0], kv[1]
		dataSubjects[key] = value
	}
	return dataSubjects, nil
}
