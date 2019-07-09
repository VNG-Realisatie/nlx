// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package configservice

import (
	"context"
	"regexp"

	google_protobuf "github.com/gogo/protobuf/types"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/config-api/configproto"
)

type setInwayConfigHandler struct {
	logger *zap.Logger

	regexpName *regexp.Regexp
}

func newSetInwayConfigHandler(db *sqlx.DB, logger *zap.Logger) (*setInwayConfigHandler, error) {
	h := &setInwayConfigHandler{
		logger: logger.With(zap.String("handler", "set-config")),
	}

	var err error

	h.regexpName, err = regexp.Compile(`^[a-zA-Z0-9-]{1,100}$`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile regexpName")
	}

	return h, nil
}

func (h *setInwayConfigHandler) SetInwayConfig(ctx context.Context, req *configproto.SetInwayConfigRequest) (*google_protobuf.Empty, error) {
	h.logger.Info("rpc request SetInwayConfig", zap.String("inway name", req.Config.Name))

	return &google_protobuf.Empty{}, nil
}
