// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package is

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"go.nlx.io/nlx/common/tls"
)

var (
	ErrSerialNumber = validation.NewError("validation_is_serial_number", "must be a valid serial number")
)

var (
	SerialNumber = validation.NewStringRuleWithError(isSerialNumber, ErrSerialNumber)
)

func isSerialNumber(value string) bool {
	err := tls.ValidateSerialNumber(value)

	return err == nil
}
