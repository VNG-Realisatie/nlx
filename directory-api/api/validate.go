// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directoryapi

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Validate the SetOrganizationContactDetailsRequest, check if all fields are valid
func (request *SetOrganizationContactDetailsRequest) Validate() error {
	return validation.ValidateStruct(
		request,
		validation.Field(&request.EmailAddress, is.Email),
	)
}
