package plugins

type Service struct {
	Name                        string
	EndpointURL                 string
	Grants                      []Grant
	Internal                    bool
	DocumentationUrl            string
	ApiSpecificationDocumentUrl string
	InsightApiUrl               string
	IrmaApiUrl                  string
	PublicSupportContact        string
	TechSupportContact          string
	OneTimeCosts                int
	MonthlyCosts                int
	RequestCosts                int
}

type Grant struct {
	OrganizationName     string
	PublicKeyPEM         string
	PublicKeyFingerprint string
}
