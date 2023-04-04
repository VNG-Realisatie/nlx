// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package app

import (
	"go.nlx.io/nlx/directory-ui/app/query"
)

type Application struct {
	Queries Queries
}

type Queries struct {
	ListServices     *query.ListServicesHandler
	ListParticipants *query.ListParticipantsHandler
}
