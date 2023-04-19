// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package service

import (
	"context"

	"go.nlx.io/nlx/management-ui-fsc/app"
)

type NewApplicationArgs struct {
	Context context.Context
}

func NewApplication(_ *NewApplicationArgs) (*app.Application, error) {
	application := &app.Application{
		Queries: app.Queries{},
	}

	return application, nil
}
