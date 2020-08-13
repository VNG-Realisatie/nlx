// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package cmd

type TLSOptions struct {
	RootCertFile string `long:"tls-root-cert" env:"TLS_ROOT_CERT" description:"Absolute or relative path to the CA root cert .pem" required:"false"`
	CertFile     string `long:"tls-cert" env:"TLS_CERT" description:"Absolute or relative path to the cert .pem" required:"false"`
	KeyFile      string `long:"tls-key" env:"TLS_KEY" description:"Absolute or relative path to the key .pem" required:"false"`
}

type TLSOrgOptions struct {
	NLXRootCert string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem" required:"true"`
	OrgCertFile string `long:"tls-org-cert" env:"TLS_ORG_CERT" description:"Absolute or relative path to the Organization cert .pem" required:"true"`
	OrgKeyFile  string `long:"tls-org-key" env:"TLS_ORG_KEY" description:"Absolute or relative path to the Organization key .pem" required:"true"`
}
