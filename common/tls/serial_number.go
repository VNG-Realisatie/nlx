// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls

import (
	"errors"
	"fmt"
)

const maxSerialNumberByteLength = 20

var (
	ErrSerialNumberEmpty   = errors.New("serial number is empty")
	ErrSerialNumberTooLong = fmt.Errorf("serial number is too long, max %d bytes", maxSerialNumberByteLength)
)

func ValidateSerialNumber(serialNumber string) error {
	if serialNumber == "" {
		return ErrSerialNumberEmpty
	}

	bytes := []byte(serialNumber)
	if len(bytes) > maxSerialNumberByteLength {
		return ErrSerialNumberTooLong
	}

	return nil
}
