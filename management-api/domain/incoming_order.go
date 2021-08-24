package domain

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IncomingOrderService struct {
	service      string
	organization string
}

type IncomingOrder struct {
	reference   string
	description string
	delegator   string
	revokedAt   *time.Time
	validFrom   time.Time
	validUntil  time.Time
	services    []IncomingOrderService
}

var ErrInvalidReference = errors.New("invalid reference")

const (
	descriptionMinLength = 1
	descriptionMaxLength = 100
)

// nolint:gocritic // these are valid regex patterns
var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-. _\s]{1,100}$`)

func NewIncomingOrder(reference, description, delegator string, revokedAt *time.Time, validFrom, validUntil time.Time, services []IncomingOrderService) (*IncomingOrder, error) {
	err := validation.Validate(reference, validation.Required)
	if err != nil {
		return nil, fmt.Errorf("reference: %s", err)
	}

	err = validation.Validate(description, validation.Required, validation.Length(descriptionMinLength, descriptionMaxLength))
	if err != nil {
		return nil, fmt.Errorf("description: %s", err)
	}

	err = validation.Validate(delegator, validation.Required, validation.Match(organizationNameRegex))
	if err != nil {
		return nil, fmt.Errorf("delegator: %s", err)
	}

	return &IncomingOrder{
		reference:   reference,
		description: description,
		delegator:   delegator,
		revokedAt: revokedAt,
		validFrom: validFrom,
		validUntil: validUntil,
		services: services,
	}, nil
}

func (i *IncomingOrder) Reference() string {
	return i.reference
}

func (i *IncomingOrder) Description() string {
	return i.description
}

func (i *IncomingOrder) Delegator() string {
	return i.delegator
}

func (i *IncomingOrder) RevokedAt() *time.Time {
	return i.revokedAt
}

func (i *IncomingOrder) ValidFrom() time.Time {
	return i.validFrom
}

func (i *IncomingOrder) ValidUntil() time.Time {
	return i.validUntil
}

func (i *IncomingOrder) Services() []IncomingOrderService {
	return i.services
}

func NewIncomingOrderService(service, organization string) IncomingOrderService {
	return IncomingOrderService{
		service: service,
		organization: organization,
	}
}

func (s *IncomingOrderService) Service() string {
	return s.service
}

func (s *IncomingOrderService) Organization() string {
	return s.organization
}