// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

type Service struct {
	Grants                      []*Grant `json:"grants"`
	Name                        string   `json:"name"`
	EndpointURL                 string   `json:"endpoint_url"`
	DocumentationURL            string   `json:"documentation_url"`
	APISpecificationDocumentURL string   `json:"api_specification_document_url"`
	PublicSupportContact        string   `json:"public_support_url"`
	TechSupportContact          string   `json:"tech_support_contact"`
	OneTimeCosts                int32    `json:"one_time_costs"`
	MonthlyCosts                int32    `json:"monthly_costs"`
	RequestCosts                int32    `json:"request_costs"`
	Internal                    bool     `json:"internal"`
}

type Grant struct {
	OrganizationSerialNumber string `json:"organization_serial_number"`
	PublicKeyPEM             string `json:"public_key_pem"`
	PublicKeyFingerprint     string `json:"public_key_fingerprint"`
}
