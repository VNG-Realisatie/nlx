package tlsconfig

import "crypto/tls"

// Defaults returns a tls.Config that has sane defaults for intra-NLX traffic.
func Defaults() *tls.Config {
	c := &tls.Config{}
	ApplyDefaults(c)
	return c
}

// ApplyDefaults sets tls.Config values to sane defaults for intra-NLX traffic.
func ApplyDefaults(c *tls.Config) {
	c.MinVersion = tls.VersionTLS12
}
