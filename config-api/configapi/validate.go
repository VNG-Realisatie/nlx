package configapi

import (
	"errors"
	"fmt"
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

	return nil
}
