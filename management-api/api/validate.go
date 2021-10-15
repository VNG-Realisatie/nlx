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

// Validate the service when creating a new one, check if all fields are valid
func (s *CreateServiceRequest) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&s.EndpointURL, validation.Required, is.URL),
		validation.Field(&s.DocumentationURL, is.URL),
		validation.Field(&s.ApiSpecificationURL, is.URL),
	)
}

// Validate the service when updating it, check if all fields are valid
func (s *UpdateServiceRequest) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&s.EndpointURL, validation.Required, is.URL),
		validation.Field(&s.DocumentationURL, is.URL),
		validation.Field(&s.ApiSpecificationURL, is.URL),
	)
}
