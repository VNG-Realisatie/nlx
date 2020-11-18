package api

import (
	"errors"
	fmt "fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// nolint:gocritic // this is a valid regex pattern
var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)

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

// Validate the service, check if all fields are valid
func (service *Service) Validate() error {
	return validation.ValidateStruct(
		service,
		validation.Field(&service.Name, validation.Required, validation.Match(nameRegex)),
		validation.Field(&service.EndpointURL, validation.Required, is.URL),
		validation.Field(&service.DocumentationURL, is.URL),
		validation.Field(&service.ApiSpecificationURL, is.URL),
	)
}

// Validate the service when creating a new one, check if all fields are valid
func (s *CreateServiceRequest) Validate() error {
	if s.Name == "" {
		return errors.New("invalid service name")
	}

	if s.EndpointURL == "" {
		return fmt.Errorf("invalid endpoint URL for service %s", s.Name)
	}

	return nil
}

// Validate the service when updating it, check if all fields are valid
func (s *UpdateServiceRequest) Validate() error {
	if s.EndpointURL == "" {
		return fmt.Errorf("invalid endpoint URL for service %s", s.Name)
	}

	return nil
}
