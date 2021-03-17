// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package plugins

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogRecordID(t *testing.T) {
	l, _ := NewLogRecordID()
	expected := 32
	size := len(l.String())

	assert.Equal(t, expected, size, "The string size of the LogRecordID")
}

var LogRecordIDResult string

func BenchmarkLogRecordID(b *testing.B) {
	var r string

	for i := 0; i < b.N; i++ {
		l, _ := NewLogRecordID()
		r = l.String()
	}

	LogRecordIDResult = r
}

func TestCreateRecordData(t *testing.T) {
	headers := http.Header{}
	headers.Add("X-NLX-Request-Process-Id", "process-id")
	headers.Add("X-NLX-Request-Data-Elements", "data-elements")
	headers.Add("X-NLX-Requester-User", "user")
	headers.Add("X-NLX-Requester-Claims", "claims")
	headers.Add("X-NLX-Request-User-Id", "user-id")
	headers.Add("X-NLX-Request-Application-Id", "application-id")
	headers.Add("X-NLX-Request-Subject-Identifier", "subject-identifier")

	recordData := createRecordData(headers, "/path")

	assert.Equal(t, "process-id", recordData["doelbinding-process-id"])
	assert.Equal(t, "data-elements", recordData["doelbinding-data-elements"])
	assert.Equal(t, "user", recordData["doelbinding-user"])
	assert.Equal(t, "claims", recordData["doelbinding-claims"])
	assert.Equal(t, "user-id", recordData["doelbinding-user-id"])
	assert.Equal(t, "application-id", recordData["doelbinding-application-id"])
	assert.Equal(t, "subject-identifier", recordData["doelbinding-subject-identifier"])
	assert.Equal(t, "/path", recordData["request-path"])
}
