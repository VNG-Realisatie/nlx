// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"errors"
	"fmt"
	"net/http"
)

type DelegationPlugin struct {
}

func NewDelegationPlugin() *DelegationPlugin {
	return &DelegationPlugin{}
}

func isDelegatedRequest(r *http.Request) bool {
	return r.Header.Get("X-NLX-Request-Delegator") != "" ||
		r.Header.Get("X-NLX-Request-OrderReference") != ""
}

func parseRequestMetadata(r *http.Request) (name, orderReference string, err error) {
	name = r.Header.Get("X-NLX-Request-Delegator")
	orderReference = r.Header.Get("X-NLX-Request-OrderReference")

	if name == "" {
		return "", "", errors.New("missing organization-name in delegation headers")
	}

	if orderReference == "" {
		return "", "", errors.New("missing order-reference in delegation headers")
	}

	return
}

func (plugin *DelegationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context Context) error {
		if !isDelegatedRequest(context.Request) {
			return next(context)
		}

		name, orderRef, err := parseRequestMetadata(context.Request)
		if err != nil {
			msg := fmt.Sprintf("failed to parse delegation metadata: %s", err.Error())

			context.Logger.Error(msg)
			http.Error(context.Response, msg, http.StatusInternalServerError)

			return nil
		}

		context.LogData["delegation"] = name
		context.LogData["orderReference"] = orderRef

		return next(context)
	}
}
