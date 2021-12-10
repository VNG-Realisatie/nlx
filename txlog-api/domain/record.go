// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"encoding/json"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Record struct {
	source        *Organization
	destination   *Organization
	direction     OrderDirection
	service       *Service
	order         *Order
	data          json.RawMessage
	transactionID string
	createdAt     time.Time
}

type OrderDirection string

const (
	IN  OrderDirection = "in"
	OUT OrderDirection = "out"
)

var orderDirections = []interface{}{IN, OUT}

type NewRecordArgs struct {
	Source        *Organization
	Destination   *Organization
	Direction     OrderDirection
	Service       *Service
	Order         *Order
	Data          json.RawMessage
	TransactionID string
	CreatedAt     time.Time
}

func NewRecord(args *NewRecordArgs) (*Record, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Source, validation.Required),
		validation.Field(&args.Destination, validation.Required),
		validation.Field(&args.Direction, validation.Required, validation.In(orderDirections...)),
		validation.Field(&args.Service, validation.Required),
		validation.Field(&args.TransactionID, validation.Required),
		validation.Field(&args.CreatedAt, validation.Required),
	)
	if err != nil {
		return nil, err
	}

	return &Record{
		source:        args.Source,
		destination:   args.Destination,
		direction:     args.Direction,
		service:       args.Service,
		order:         args.Order,
		data:          args.Data,
		transactionID: args.TransactionID,
		createdAt:     args.CreatedAt,
	}, nil
}

func (r *Record) Source() *Organization {
	return r.source
}

func (r *Record) Destination() *Organization {
	return r.destination
}

func (r *Record) Direction() OrderDirection {
	return r.direction
}

func (r *Record) Service() *Service {
	return r.service
}

func (r *Record) Order() *Order {
	return r.order
}

func (r *Record) Data() json.RawMessage {
	return r.data
}

func (r *Record) TransactionID() string {
	return r.transactionID
}

func (r *Record) CreatedAt() time.Time {
	return r.createdAt
}
