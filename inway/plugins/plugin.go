// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

import (
	"net/http"

	"go.uber.org/zap"
)

type Destination struct {
	Organization string
	Service      *Service
	Path         string
}

type Context struct {
	Logger      *zap.Logger
	Destination *Destination
	Response    http.ResponseWriter
	Request     *http.Request
	LogData     map[string]string
	AuthInfo    *AuthInfo
}

type AuthInfo struct {
	OrganizationName     string
	PublicKeyFingerprint string
}

type ServeFunc func(context *Context) error

type Plugin interface {
	Serve(next ServeFunc) ServeFunc
}

func BuildChain(serve ServeFunc, pluginList ...Plugin) ServeFunc {
	if len(pluginList) == 0 {
		return serve
	}

	return pluginList[0].Serve(BuildChain(serve, pluginList[1:]...))
}
