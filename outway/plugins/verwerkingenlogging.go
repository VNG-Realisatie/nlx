package plugins

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/verwerkingenlogging"
)

type VerwerkingenLogginPlugin struct {
	logger                      *zap.Logger
	organizationSerialNumber    string
	verwerkingenLogginAPIClient verwerkingenlogging.VerwerkingenLoggingAPI
}

func NewVerwerkingenLoggingPlugin(logger *zap.Logger, organizationSerialNumber string, client verwerkingenlogging.VerwerkingenLoggingAPI) *VerwerkingenLogginPlugin {
	return &VerwerkingenLogginPlugin{
		organizationSerialNumber:    organizationSerialNumber,
		verwerkingenLogginAPIClient: client,
		logger:                      logger,
	}
}

func (plugin *VerwerkingenLogginPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		plugin.logger.Info("verwerkingenloging plugin")

		request, err := verwerkingenlogging.BuildLogRequestFromHeaders(context.Request.Header)
		if err != nil {
			if errors.Is(err, verwerkingenlogging.ErrHeadersDoNotContainVerwerkingenLoggingData) {
				plugin.logger.Info("request does not contain verwerkingenlogging headers")
				return next(context)
			}

			plugin.logger.Error("error building write verwerkingenlog request from headers", zap.Error(err))

			http.Error(context.Response, `nlx-outway: failed to write to verwerkingenglogging api`, http.StatusInternalServerError)

			return nil
		}

		request.Uitvoerder = plugin.organizationSerialNumber
		context.Request.Header.Add(verwerkingenlogging.HeaderUitvoerder, plugin.organizationSerialNumber)

		err = plugin.verwerkingenLogginAPIClient.WriteLog(request)
		if err != nil {
			plugin.logger.Error("error writing to verwerkinglog", zap.Error(err))
			http.Error(context.Response, `nlx-outway: failed to write to verwerkingenglogging api`, http.StatusInternalServerError)

			return nil
		}

		return next(context)
	}
}
