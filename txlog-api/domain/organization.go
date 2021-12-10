// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"fmt"

	"go.nlx.io/nlx/common/tls"
)

type Organization struct {
	serialNumber string
}

func NewOrganization(serialNumber string) (*Organization, error) {
	err := tls.ValidateSerialNumber(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("error validating organization serial number: %s", err)
	}

	return &Organization{
		serialNumber: serialNumber,
	}, nil
}

func (i *Organization) SerialNumber() string {
	return i.serialNumber
}
