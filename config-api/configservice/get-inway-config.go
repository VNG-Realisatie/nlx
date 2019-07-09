// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package configservice

import (
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/config-api/configproto"
)

type getInwayConfigHandler struct {
	logger *zap.Logger

	regexpName *regexp.Regexp
}

func newGetInwayConfigHandler(db *sqlx.DB, logger *zap.Logger) (*getInwayConfigHandler, error) {
	h := &getInwayConfigHandler{
		logger: logger.With(zap.String("handler", "get-config")),
	}

	var err error

	h.regexpName, err = regexp.Compile(`^[a-zA-Z0-9-]{1,100}$`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile regexpName")
	}

	return h, nil
}

func (h *getInwayConfigHandler) GetInwayConfig(req *configproto.GetInwayConfigRequest, stream configproto.Config_GetInwayConfigServer) error {
	h.logger.Info("rpc request GetInwayConfig", zap.String("inway name", req.Name))
	resp := &configproto.GetInwayConfigResponse{}
	err := stream.Send(resp)
	if err != nil {
		return errors.Wrap(err, "failed to send initial config")
	}

	return nil
}
