package certportal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewCertPortal(t *testing.T) {
	certPortal := NewCertPortal(zap.NewNop(), "host")
	assert.NotNil(t, certPortal)
	assert.Equal(t, certPortal.caHost, "host")
	assert.NotNil(t, certPortal.router)
	assert.Equal(t, 3, len(certPortal.GetRouter().Routes()))
}

func TextCreateSigner(t *testing.T) {
	certPortal := NewCertPortal(zap.NewNop(), "host")
	signer, err := certPortal.createSigner()
	assert.NoError(t, err)
	assert.NotNil(t, signer)
}
