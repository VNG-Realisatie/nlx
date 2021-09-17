// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package certportal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	certportal "go.nlx.io/nlx/ca-certportal"
)

func Test_GenerateDemoSerialNumber(t *testing.T) {
	got, err := certportal.GenerateDemoSerialNumber()
	assert.NoError(t, err)
	assert.Len(t, got, 20)
}
