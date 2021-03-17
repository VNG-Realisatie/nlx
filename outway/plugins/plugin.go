// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"net/http"

	"go.uber.org/zap"
)

type Destination struct {
	Organization string
	Service      string
	Path         string
}

type Context struct {
	Logger      *zap.Logger
	Destination *Destination
	Response    http.ResponseWriter
	Request     *http.Request
	LogData     map[string]string
}

type ServeFunc func(context Context) error

type Plugin interface {
	Serve(next ServeFunc) ServeFunc
}
