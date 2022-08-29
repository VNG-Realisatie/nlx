// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package query

import (
	"encoding/json"
	"time"
)

type Record struct {
	SourceOrganization      string
	DestinationOrganization string
	Direction               string
	ServiceName             string
	OrderReference          string
	Delegator               string
	Data                    json.RawMessage
	TransactionID           string
	CreatedAt               time.Time
}
