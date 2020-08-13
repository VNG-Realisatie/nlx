package inway

import (
	"net/http"

	common_tls "go.nlx.io/nlx/common/tls"
)

// setupClient create a test client with certificates
func setupClient(cert *common_tls.CertificateBundle) http.Client {
	config := cert.TLSConfig()
	config.InsecureSkipVerify = true

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := http.Client{
		Transport: tr,
	}
	return client
}
