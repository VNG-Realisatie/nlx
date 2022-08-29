// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package record

import (
	"encoding/json"
	"errors"
	"time"

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

	if args.SourceOrganization == "" {
		return nil, errors.New("SourceOrganization: cannot be blank")
	}

	sourceErr := tls.ValidateSerialNumber(args.SourceOrganization)
	if sourceErr != nil {
		return nil, errors.New("SourceOrganization: " + sourceErr.Error())
	}

	if args.DestinationOrganization == "" {
		return nil, errors.New("DestinationOrganization: cannot be blank")
	}

	destErr := tls.ValidateSerialNumber(args.DestinationOrganization)
	if destErr != nil {
		return nil, errors.New("DestinationOrganization: " + destErr.Error())
	}

	if args.Delegator != "" {
		delegatorErr := tls.ValidateSerialNumber(args.Delegator)
		if delegatorErr != nil {
			return nil, errors.New("Delegator: " + delegatorErr.Error())
		}
	}

	if args.Direction == "" {
		return nil, errors.New("Direction: cannot be blank")
	}

	if args.Direction != IN && args.Direction != OUT {
		return nil, errors.New("Direction: must be IN or OUT")
	}

	if args.ServiceName == "" {
		return nil, errors.New("ServiceName: cannot be blank")
	}

	if args.TransactionID == "" {
		return nil, errors.New("TransactionID: cannot be blank")
	}

	if args.CreatedAt.IsZero() {
		return nil, errors.New("CreatedAt: cannot be blank")
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
