// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package record

import (
	"encoding/json"
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"go.nlx.io/nlx/common/tls"
)

type Record struct {
	sourceOrganization      string
	destinationOrganization string
	direction               OrderDirection
	serviceName             string
	orderReference          string
	delegator               string
	data                    json.RawMessage
	transactionID           string
	createdAt               time.Time
	dataSubjects            map[string]string
}

type OrderDirection string

const (
	IN  OrderDirection = "in"
	OUT OrderDirection = "out"
)

var orderDirections = []interface{}{IN, OUT}

type NewRecordArgs struct {
	SourceOrganization      string
	DestinationOrganization string
	Direction               OrderDirection
	ServiceName             string
	OrderReference          string
	Delegator               string
	Data                    json.RawMessage
	TransactionID           string
	CreatedAt               time.Time
	DataSubjects            map[string]string
}

func NewRecord(args *NewRecordArgs) (*Record, error) {
	if args.Delegator == "" && args.OrderReference != "" {
		return nil, errors.New("empty delegator, both the delegator and order reference should be provided")
	}

	if args.Delegator != "" && args.OrderReference == "" {
		return nil, errors.New("empty order reference, both the delegator and order reference should be provided")
	}

	err := validation.ValidateStruct(
		args,
		validation.Field(&args.SourceOrganization, validation.Required, validation.By(validateSerialNumber)),
		validation.Field(&args.DestinationOrganization, validation.Required, validation.By(validateSerialNumber)),
		validation.Field(&args.Delegator, validation.When(args.Delegator != "", validation.By(validateSerialNumber))),
		validation.Field(&args.Direction, validation.Required, validation.In(orderDirections...)),
		validation.Field(&args.ServiceName, validation.Required),
		validation.Field(&args.TransactionID, validation.Required),
		validation.Field(&args.CreatedAt, validation.Required),
	)
	if err != nil {
		return nil, err
	}

	return &Record{
		sourceOrganization:      args.SourceOrganization,
		destinationOrganization: args.DestinationOrganization,
		direction:               args.Direction,
		serviceName:             args.ServiceName,
		orderReference:          args.OrderReference,
		delegator:               args.Delegator,
		data:                    args.Data,
		transactionID:           args.TransactionID,
		createdAt:               args.CreatedAt,
		dataSubjects:            args.DataSubjects,
	}, nil
}

func validateSerialNumber(value interface{}) error {
	s, _ := value.(string)
	return tls.ValidateSerialNumber(s)
}

func (r *Record) SourceOrganization() string {
	return r.sourceOrganization
}

func (r *Record) DestinationOrganization() string {
	return r.destinationOrganization
}

func (r *Record) Direction() OrderDirection {
	return r.direction
}

func (r *Record) ServiceName() string {
	return r.serviceName
}

func (r *Record) OrderReference() string {
	return r.orderReference
}

func (r *Record) Delegator() string {
	return r.delegator
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

func (r *Record) DataSubjects() map[string]string {
	return r.dataSubjects
}
