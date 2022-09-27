// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package api

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// nolint:gocritic // this is a valid regex pattern
var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-.]{1,100}$`)

// Validate the inway, check if all fields are valid
func (inway *Inway) Validate() error {
	return validation.ValidateStruct(
		inway,
		validation.Field(&inway.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&inway.Hostname, is.Host),
		validation.Field(&inway.IpAddress, is.IP),
		validation.Field(&inway.SelfAddress, is.URL),
	)
}

// Validate the outway, check if all fields are valid
func (req *RegisterOutwayRequest) Validate() error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&req.PublicKeyPem, validation.Required),
		validation.Field(&req.SelfAddressApi, validation.Required),
		validation.Field(&req.Version, validation.Required),
	)
}

// Validate the service when creating a new one, check if all fields are valid
func (req *CreateServiceRequest) Validate() error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&req.EndpointUrl, validation.Required, is.URL),
		validation.Field(&req.DocumentationUrl, is.URL),
		validation.Field(&req.ApiSpecificationUrl, is.URL),
	)
}

// Validate the service when updating it, check if all fields are valid
func (req *UpdateServiceRequest) Validate() error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&req.EndpointUrl, validation.Required, is.URL),
		validation.Field(&req.DocumentationUrl, is.URL),
		validation.Field(&req.ApiSpecificationUrl, is.URL),
	)
}
