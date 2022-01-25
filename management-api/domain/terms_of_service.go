// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TermsOfServiceStatus struct {
	username  string
	createdAt time.Time
}

type NewTermsOfServiceStatusArgs struct {
	Username  string
	CreatedAt time.Time
}

func NewTermsOfServiceStatus(args *NewTermsOfServiceStatusArgs) (*TermsOfServiceStatus, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Username, validation.Required),
		validation.Field(&args.CreatedAt, validation.Max(time.Now()).Error("must not be in the future")),
	)

	if err != nil {
		return nil, err
	}

	return &TermsOfServiceStatus{
		username:  args.Username,
		createdAt: args.CreatedAt,
	}, nil
}

func (s *TermsOfServiceStatus) Username() string {
	return s.username
}

func (s *TermsOfServiceStatus) CreatedAt() time.Time {
	return s.createdAt
}
