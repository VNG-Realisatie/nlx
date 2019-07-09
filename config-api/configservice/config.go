// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package configservice

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/config-api/configproto"
)

// compile-time interface implementation verification
var _ configproto.ConfigServer = &ConfigService{}

// ConfigService handles all requests for a config api
type ConfigService struct {
	*getInwayConfigHandler
	*setInwayConfigHandler
}

// New sets up a new ConfigService and returns an error when something failed during setup.
func New(logger *zap.Logger, db *sqlx.DB) (*ConfigService, error) {
	s := &ConfigService{}

	var err error

	s.getInwayConfigHandler, err = newGetInwayConfigHandler(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup GetInwayConfigHandler handler")
	}
	s.setInwayConfigHandler, err = newSetInwayConfigHandler(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup SetConfigHandler handler")
	}

	return s, nil
}
