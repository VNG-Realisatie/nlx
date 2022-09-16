// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package external

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	common_is "go.nlx.io/nlx/common/validation/is"
)

func (req *RequestClaimRequest) Validate() error {
	return validation.ValidateStruct(
		req,
		validation.Field(&req.OrderReference, validation.Required),
		validation.Field(&req.ServiceName, validation.Required),
		validation.Field(&req.ServiceOrganizationSerialNumber, validation.Required, common_is.SerialNumber),
	)
}
