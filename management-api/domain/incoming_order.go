package domain

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	common_tls "go.nlx.io/nlx/common/tls"
)

type incomingOrderServiceOrganization struct {
	serialNumber string
	name         string
}

type IncomingOrderService struct {
	service      string
	organization incomingOrderServiceOrganization
}

func NewIncomingOrderService(service, organizationSerialNumber, organizationName string) IncomingOrderService {
	return IncomingOrderService{
		service: service,
		organization: incomingOrderServiceOrganization{
			serialNumber: organizationSerialNumber,
			name:         organizationName,
		},
	}
}

func (s *IncomingOrderService) Service() string {
	return s.service
}

func (s *IncomingOrderService) OrganizationSerialNumber() string {
	return s.organization.serialNumber
}

func (s *IncomingOrderService) OrganizationName() string {
	return s.organization.name
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

type NewIncomingOrderArgs struct {
	Reference   string
	Description string
	Delegator   string
	RevokedAt   *time.Time
	ValidFrom   time.Time
	ValidUntil  time.Time
	Services    []IncomingOrderService
}

const (
	descriptionMinLength = 1
	descriptionMaxLength = 100
)

// Usage is documented in /docs/docs/reference-information/data-validation.md
// nolint:gocritic // these are valid regex patterns
var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-. _\s]{1,100}$`)
var serviceNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)

func validateOrganizationSerialNumber(value interface{}) error {
	valueAsString, _ := value.(string)

	err := common_tls.ValidateSerialNumber(valueAsString)
	if err != nil {
		return fmt.Errorf("organization serial number must be in a valid format: %s", err)
	}

	return err
}

func NewIncomingOrder(args *NewIncomingOrderArgs) (*IncomingOrder, error) {
	err := validation.Validate(args.Reference, validation.Required)
	if err != nil {
		return nil, fmt.Errorf("reference: %s", err)
	}

	err = validation.Validate(args.Description, validation.Required, validation.Length(descriptionMinLength, descriptionMaxLength))
	if err != nil {
		return nil, fmt.Errorf("description: %s", err)
	}

	err = validation.Validate(args.Delegator, validation.Required, validation.Match(organizationNameRegex))
	if err != nil {
		return nil, fmt.Errorf("delegator: %s", err)
	}

	err = validation.Validate(args.ValidUntil, validation.Required, validation.Min(args.ValidFrom).Error("order can not expire before the start date"))
	if err != nil {
		return nil, fmt.Errorf("valid from: %s", err)
	}

	err = validation.Validate(args.Services, validation.Each(validation.By(func(value interface{}) error {
		orderService, ok := value.(IncomingOrderService)
		if !ok {
			return errors.New("expecting an order-service")
		}

		err = validation.Validate(orderService.organization.name, validation.Match(organizationNameRegex).Error("organization must be in a valid format"))
		if err != nil {
			return fmt.Errorf("organization: %s", err)
		}

		err = validation.Validate(orderService.organization.serialNumber, validation.By(validateOrganizationSerialNumber))
		if err != nil {
			return fmt.Errorf("organization: %s", err)
		}

		err = validation.Validate(orderService.service, validation.Match(serviceNameRegex).Error("service must be in a valid format"))
		if err != nil {
			return fmt.Errorf("name: %s", err)
		}

		return nil
	})))
	if err != nil {
		return nil, fmt.Errorf("services: %s", err)
	}

	return &IncomingOrder{
		reference:   args.Reference,
		description: args.Description,
		delegator:   args.Delegator,
		revokedAt:   args.RevokedAt,
		validFrom:   args.ValidFrom,
		validUntil:  args.ValidUntil,
		services:    args.Services,
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
