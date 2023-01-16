// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package domain

import "time"

type Participant struct {
	Name          string
	Since         time.Time
	ServicesCount uint
	InwaysCount   uint
	OutwaysCount  uint
}
