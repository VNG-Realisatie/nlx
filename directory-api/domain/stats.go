// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type VersionStatistics struct {
	gatewayType VersionStatisticsType
	version     string
	amount      uint32
}

type VersionStatisticsType string

const (
	TypeInway  VersionStatisticsType = "inway"
	TypeOutway VersionStatisticsType = "outway"
)

type NewVersionStatisticsArgs struct {
	GatewayType VersionStatisticsType
	Version     string
	Amount      uint32
}

func NewVersionStatistics(args *NewVersionStatisticsArgs) (*VersionStatistics, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.GatewayType, validation.Required),
		validation.Field(&args.Version, validation.Required),
		validation.Field(&args.Amount, validation.Required),
	)
	if err != nil {
		return nil, err
	}

	return &VersionStatistics{
		gatewayType: args.GatewayType,
		version:     args.Version,
		amount:      args.Amount,
	}, nil
}

func (v *VersionStatistics) GatewayType() VersionStatisticsType {
	return v.gatewayType
}

func (v *VersionStatistics) Version() string {
	return v.version
}

func (v *VersionStatistics) Amount() uint32 {
	return v.amount
}
