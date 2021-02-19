// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

// Direction declares whether the record was created by a requesting (out) or providing (in) organization.
type Direction string

const (
	// DirectionIn is the direction for records written by inway
	DirectionIn Direction = "in"
	// DirectionOut is the direction for records written by outway
	DirectionOut Direction = "out"
)
