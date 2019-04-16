package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// AuthorizationMode holds the authorization details of a API
type AuthorizationMode struct {
	Mode          string   `yaml:"mode"`
	Organizations []string `yaml:"organizations"`
}

type Config struct {
	Kind   string       `json:"kind"`
	Config *InwayConfig `json:"config"`
}

// InwayConfig is the top-level for the service configuration file.
type InwayConfig struct {
	Name              string       `yaml:"name"`
	ListenAddress     string       `yaml:"listenAddress"`
	SelfAddress       string       `yaml:"selfAddress"`
	DisableLogging    bool         `yaml:"disableLogging"`
	DirectoryAddress  string       `yaml:"directoryAddress"`
	TransactionLogDSN string       `yaml:"transactionLogDSN"`
	APIS              []APIDetails `yaml:"APIs"`
}

// APIDetails holds the details for a single service definition.
type APIDetails struct {
	Name                        string             `yaml:"name"`
	EndpointURL                 string             `yaml:"endpointURL"`
	DocumentationURL            string             `yaml:"documentationURL"`
	APISpecificationDocumentURL string             `yaml:"apiSpecificationURL"`
	CACertPath                  string             `yaml:"caCertPath"`
	PublicSupportContact        string             `yaml:"publicSupportContact"`
	TechSupportContact          string             `yaml:"techSupportContact"`
	Internal                    bool               `yaml:"internal"`
	Authorization               *AuthorizationMode `yaml:"authorizationSettings"`
}

// ParseConfig reads the service config from bytes and returns.
func ParseConfig(logger *zap.Logger, configBytes []byte) (*Config, error) {
	inwayConfig := &Config{}
	err := yaml.Unmarshal(configBytes, inwayConfig)
	if err != nil {
		return nil, err
	}
	return inwayConfig, nil
}
