// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package testingutils

import (
	"fmt"
	"path/filepath"

	common_tls "go.nlx.io/nlx/common/tls"
)

type CertificateBundleOrganizationName string

const (
	OrgNLXTest             CertificateBundleOrganizationName = "org-nlx-test"
	OrgNLXTestB            CertificateBundleOrganizationName = "org-nlx-test-b"
	OrgWithoutName         CertificateBundleOrganizationName = "org-without-name"
	OrgWithoutSerialNumber CertificateBundleOrganizationName = "org-without-serial-number"
)

func GetCertificateBundle(pkiDir string, name CertificateBundleOrganizationName) (*common_tls.CertificateBundle, error) {
	return common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, fmt.Sprintf("%s-chain.pem", name)),
		filepath.Join(pkiDir, fmt.Sprintf("%s-key.pem", name)),
		filepath.Join(pkiDir, "ca-root.pem"),
	)
}
