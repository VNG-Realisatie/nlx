package configapi

import (
	"errors"
	"fmt"
)

const (
	authorizationModeWhitelist = "whitelist"
	authorizationModeNone      = "none"
)

// Validate the inway, check if all fields are valid
func (i *Inway) Validate() error {
	if i.Name == "" {
		return errors.New("invalid inway name")
	}

	return nil
}

// Validate the service, check if all fields are valid
func (s *Service) Validate() error {
	if s.Name == "" {
		return errors.New("invalid service name")
	}

	if s.EndpointURL == "" {
		return fmt.Errorf("invalid endpoint URL for service %s", s.Name)
	}

	if s.AuthorizationSettings == nil {
		return fmt.Errorf("invalid authorization settings for service %s", s.Name)
	}

	if s.AuthorizationSettings.Mode != authorizationModeWhitelist && s.AuthorizationSettings.Mode != authorizationModeNone {
		return fmt.Errorf("invalid authorization mode for service %s, expected whitelist or none, got %s", s.Name, s.AuthorizationSettings.Mode)
	}

	if s.AuthorizationSettings.Mode == authorizationModeWhitelist && s.AuthorizationSettings.Authorizations == nil {
		return fmt.Errorf("invalid list of authorizations for service %s, expected list of authorizations, got nil", s.Name)
	}

	return nil
}
