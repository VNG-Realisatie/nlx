package api

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	common_is "go.nlx.io/nlx/common/validation/is"
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
		validation.Field(&req.PublicKeyPEM, validation.Required),
		validation.Field(&req.SelfAddressAPI, validation.Required),
		validation.Field(&req.Version, validation.Required),
	)
}

// Validate the service when creating a new one, check if all fields are valid
func (req *CreateServiceRequest) Validate() error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&req.EndpointURL, validation.Required, is.URL),
		validation.Field(&req.DocumentationURL, is.URL),
		validation.Field(&req.ApiSpecificationURL, is.URL),
	)
}

// Validate the service when updating it, check if all fields are valid
func (req *UpdateServiceRequest) Validate() error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&req.EndpointURL, validation.Required, is.URL),
		validation.Field(&req.DocumentationURL, is.URL),
		validation.Field(&req.ApiSpecificationURL, is.URL),
	)
}

func (req *RetrieveClaimForOrderRequest) Validate() error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.OrderReference, validation.Required),
		validation.Field(&req.OrderOrganizationSerialNumber, validation.Required, common_is.SerialNumber),
		validation.Field(&req.ServiceName, validation.Required),
		validation.Field(&req.ServiceOrganizationSerialNumber, validation.Required, common_is.SerialNumber),
	)
}
