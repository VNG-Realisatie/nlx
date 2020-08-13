// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	common_tls "go.nlx.io/nlx/common/tls"
)

func TestNewCertPoolFromFile(t *testing.T) {
	file := filepath.Join(pkiDir, "ca-root.pem")

	p, c, err := common_tls.NewCertPoolFromFile(file)

	assert.NoError(t, err)

	s := p.Subjects()
	assert.Len(t, s, 1)
	assert.Equal(t, c.RawSubject, s[0])
}
