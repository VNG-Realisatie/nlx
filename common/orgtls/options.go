// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package orgtls

// TLSOptions defines the TLS options for a common NLX component.
type TLSOptions struct {
	NLXRootCert string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem" required:"true"`
	OrgCertFile string `long:"tls-org-cert" env:"TLS_ORG_CERT" description:"Absolute or relative path to the Organization cert .pem" required:"true"`
	OrgKeyFile  string `long:"tls-org-key" env:"TLS_ORG_KEY" description:"Absolute or relative path to the Organization key .pem" required:"true"`
}
