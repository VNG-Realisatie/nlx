// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Participant struct {
	organization *Organization
	statistics   *ParticipantStatistics
	createdAt    time.Time
}

type ParticipantStatistics struct {
	inways   uint
	outways  uint
	services uint
}

type NewParticipantArgs struct {
	Organization *Organization
	Statistics   *NewParticipantStatisticsArgs
	CreatedAt    time.Time
}

type NewParticipantStatisticsArgs struct {
	Inways   uint
	Outways  uint
	Services uint
}

func NewParticipant(args *NewParticipantArgs) (*Participant, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Organization, validation.NotNil),
		validation.Field(&args.Statistics, validation.NotNil),
		validation.Field(&args.CreatedAt, validation.Max(time.Now()).Error("must not be in the future")),
	)

	if err != nil {
		return nil, err
	}

	return &Participant{
		organization: args.Organization,
		statistics: &ParticipantStatistics{
			inways:   args.Statistics.Inways,
			outways:  args.Statistics.Outways,
			services: args.Statistics.Services,
		},
		createdAt: args.CreatedAt,
	}, nil
}

func (p *Participant) Organization() *Organization {
	return p.organization
}

func (p *Participant) Statistics() *ParticipantStatistics {
	return p.statistics
}

func (p *Participant) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Participant) ToString() string {
	return fmt.Sprintf(
		"organization serial number: %s, organization name: %s, statistics: %s, created at: %s",
		p.organization.SerialNumber(), p.organization.Name(), p.Statistics().ToString(), p.CreatedAt(),
	)
}

func (p *ParticipantStatistics) Inways() uint {
	return p.inways
}

func (p *ParticipantStatistics) Outways() uint {
	return p.outways
}

func (p *ParticipantStatistics) Services() uint {
	return p.services
}

func (p *ParticipantStatistics) ToString() string {
	return fmt.Sprintf(
		"inways: %d, outways: %d, services: %d",
		p.Inways(), p.Outways(), p.Services(),
	)
}
