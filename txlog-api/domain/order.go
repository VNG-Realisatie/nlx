// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Order struct {
	delegator string
	reference string
}

type NewOrderArgs struct {
	Delegator string
	Reference string
}

func NewOrder(args *NewOrderArgs) (*Order, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Delegator, validation.When(len(args.Reference) > 0, validation.Required)),
		validation.Field(&args.Reference, validation.When(len(args.Delegator) > 0, validation.Required)),
	)
	if err != nil {
		return nil, err
	}

	return &Order{
		delegator: args.Delegator,
		reference: args.Reference,
	}, nil
}

func (o *Order) Delegator() string {
	return o.delegator
}

func (o *Order) Reference() string {
	return o.reference
}
