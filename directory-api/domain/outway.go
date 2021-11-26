// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Outway struct {
	name         string
	organization *Organization
	nlxVersion   string
	createdAt    time.Time
	updatedAt    time.Time
}

type NewOutwayArgs struct {
	Name         string
	Organization *Organization
	NlxVersion   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewOutway(args *NewOutwayArgs) (*Outway, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Name, validation.When(len(args.Name) > 0, validation.Match(nameRegex))),
		validation.Field(&args.Organization, validation.NotNil),
		validation.Field(&args.NlxVersion, validation.When(args.NlxVersion != NlxVersionUnknown, is.Semver)),
		validation.Field(&args.CreatedAt, validation.Max(time.Now()).Error("must not be in the future")),
		validation.Field(&args.UpdatedAt, validation.Max(time.Now()).Error("must not be in the future")),
	)

	if err != nil {
		return nil, err
	}

	return &Outway{
		name:         args.Name,
		organization: args.Organization,
		nlxVersion:   args.NlxVersion,
		createdAt:    args.CreatedAt,
		updatedAt:    args.UpdatedAt,
	}, nil
}

func (i *Outway) Name() string {
	return i.name
}

func (i *Outway) Organization() *Organization {
	return i.organization
}

func (i *Outway) NlxVersion() string {
	return i.nlxVersion
}

func (i *Outway) CreatedAt() time.Time {
	return i.createdAt
}

func (i *Outway) UpdatedAt() time.Time {
	return i.updatedAt
}

func (i *Outway) ToString() string {
	return fmt.Sprintf(
		"name: %s, organization serial number: %s, organization name: %s, nlx version: %s, created at: %s, updated at: %s",
		i.Name(), i.organization.serialNumber, i.organization.Name(), i.NlxVersion(), i.CreatedAt(), i.UpdatedAt(),
	)
}
