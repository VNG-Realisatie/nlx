// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package outway

import (
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
