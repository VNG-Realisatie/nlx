// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package config

import (
	"strings"

	"github.com/ktr0731/toml"
	"github.com/pkg/errors"
)

// serviceConfigVersion is a minimal variant of the service configuration file to be able to select the correct parser.
type serviceConfigVersion struct {
	Version int `toml:"version"`
}

// ServiceConfigV1 is a previous version of the service configuration file.
type serviceConfigV1 struct {
	serviceConfigVersion
	Services map[string]serviceDetailsV1
}

// ServiceConfig is the top-level for the service configuration file.
type ServiceConfig struct {
	serviceConfigVersion
	Services map[string]ServiceDetails
}

type AuthorizationModel string

const (
	AuthorizationmodelNone      AuthorizationModel = "none"
	AuthorizationmodelWhitelist AuthorizationModel = "whitelist"
	ServiceConfigVersionV1      int                = 1
	ServiceConfigVersionV2      int                = 2
)

// AuthorizationWhitelistItem holds the criteria for the whitelist.
type AuthorizationWhitelistItem struct {
	OrganizationName string `toml:"organization-name"`
	PublicKey        string `toml:"public-key"`
}

// ServiceDetails holds the details for a single service definition.
type ServiceDetailsBase struct {
	EndpointURL                 string             `toml:"endpoint-url"`
	AuthorizationModel          AuthorizationModel `toml:"authorization-model"`
	DocumentationURL            string             `toml:"documentation-url"`              // Config parameter will be moved to directory admin interface
	APISpecificationDocumentURL string             `toml:"api-specification-document-url"` // Config parameter will be moved to directory admin interface
	InsightAPIURL               string             `toml:"insight-api-url"`                // Config parameter will be moved to directory admin interface
	IrmaAPIURL                  string             `toml:"irma-api-url"`                   // Config parameter will be moved to directory admin interface
	CACertPath                  string             `toml:"ca-cert-path"`
	PublicSupportContact        string             `toml:"public-support-contact"`
	TechSupportContact          string             `toml:"tech-support-contact"`
	Internal                    bool               `toml:"internal"`
}

type serviceDetailsV1 struct {
	ServiceDetailsBase
	AuthorizationWhitelist []string `toml:"authorization-whitelist"`
}

type ServiceDetails struct {
	ServiceDetailsBase
	AuthorizationWhitelist []AuthorizationWhitelistItem `toml:"authorization-whitelist"`
}

// LoadServiceConfig reads the service config from disk and returns.
func LoadServiceConfig(serviceConfigLocation string) (*ServiceConfig, error) {
	serviceConfigVersion := &serviceConfigVersion{}

	tomlMetaData, err := toml.DecodeFile(serviceConfigLocation, serviceConfigVersion)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load service config")
	}

	if !tomlMetaData.IsDefined("version") || serviceConfigVersion.Version == ServiceConfigVersionV1 { // Original spec, ServiceConfigV1
		serviceConfigV1, err := loadServiceConfigV1(serviceConfigLocation)
		if err != nil {
			return nil, err
		}

		return convertServiceConfigV1ToV2(serviceConfigV1), errors.Errorf("%s uses deprected version %d, please upgrade to %d", serviceConfigLocation, ServiceConfigVersionV1, ServiceConfigVersionV2)
	}

	if serviceConfigVersion.Version == ServiceConfigVersionV2 {
		return loadServiceConfigV2(serviceConfigLocation)
	}

	return nil, errors.Errorf("unsupported version(%v) in toml", serviceConfigVersion.Version)
}

func loadServiceConfigV2(serviceConfigLocation string) (*ServiceConfig, error) {
	serviceConfig := &ServiceConfig{}

	err := loadServiceConfigFile(serviceConfigLocation, serviceConfig)
	if err != nil {
		return nil, err
	}

	return serviceConfig, nil
}

func loadServiceConfigV1(serviceConfigLocation string) (*serviceConfigV1, error) {
	serviceConfig := &serviceConfigV1{}

	err := loadServiceConfigFile(serviceConfigLocation, serviceConfig)
	if err != nil {
		return nil, err
	}

	return serviceConfig, nil
}

func convertServiceConfigV1ToV2(serviceConfigV1 *serviceConfigV1) *ServiceConfig {
	serviceConfig := &ServiceConfig{}

	if len(serviceConfigV1.Services) > 0 {
		serviceConfig.Services = make(map[string]ServiceDetails)
	}

	for key := range serviceConfigV1.Services {
		serviceDetailsV1 := serviceConfigV1.Services[key]

		var authorizationWhitelist []AuthorizationWhitelistItem
		for _, organizationName := range serviceDetailsV1.AuthorizationWhitelist {
			authorizationWhitelist = append(authorizationWhitelist, AuthorizationWhitelistItem{
				OrganizationName: organizationName,
			})
		}

		serviceConfig.Services[key] = ServiceDetails{ServiceDetailsBase: serviceDetailsV1.ServiceDetailsBase, AuthorizationWhitelist: authorizationWhitelist}
	}

	return serviceConfig
}

func loadServiceConfigFile(serviceConfigLocation string, serviceConfig interface{}) error {
	tomlMetaData, err := toml.DecodeFile(serviceConfigLocation, serviceConfig)
	if err != nil {
		return err
	}

	if len(tomlMetaData.Undecoded()) > 0 {
		return errors.Errorf("unsupported values in toml. key: %s", strings.Join(tomlMetaData.Undecoded()[0], ">"))
	}

	return nil
}
