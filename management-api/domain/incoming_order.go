package domain

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IncomingOrderService struct {
	IncomingOrderID uint
	Service         string
	Organization    string
}

type IncomingOrder struct {
	// id           uint
	reference string
}

var ErrInvalidReference = errors.New("invalid reference")

func NewIncomingOrder(reference string) (*IncomingOrder, error) {
	err := validation.Validate(reference, validation.Required)
	if err != nil {
		return nil, fmt.Errorf("reference: %s", err)
	}

	return &IncomingOrder{
		reference: reference,
	}, nil
}

func (i *IncomingOrder) Reference() string {
	return i.reference
}
