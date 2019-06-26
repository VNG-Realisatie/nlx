// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package irma_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/insight-api/irma"
)

func TestNewJWTHandler(t *testing.T) {
	jwtHandler := irma.NewJWTGenerator()
	assert.NotNil(t, jwtHandler)
}
