// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

type Service struct {
	Grants                      []*Grant
	Name                        string
	EndpointURL                 string
	DocumentationURL            string
	APISpecificationDocumentURL string
	InsightAPIURL               string
	IrmaAPIURL                  string
	PublicSupportContact        string
	TechSupportContact          string
	OneTimeCosts                int32
	MonthlyCosts                int32
	RequestCosts                int32
	Internal                    bool
}

type Grant struct {
	OrganizationName     string
	PublicKeyPEM         string
	PublicKeyFingerprint string
}
