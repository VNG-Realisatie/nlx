// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import "go.nlx.io/nlx/common/delegation"

type StripHeadersPlugin struct {
	organizationSerialNumber string
}

func NewStripHeadersPlugin(organizationSerialNumber string) *StripHeadersPlugin {
	return &StripHeadersPlugin{
		organizationSerialNumber: organizationSerialNumber,
	}
}

func (plugin *StripHeadersPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		if plugin.organizationSerialNumber != context.Destination.OrganizationSerialNumber {
			context.Request.Header.Del("X-NLX-Requester-User")
			context.Request.Header.Del("X-NLX-Requester-Claims")
			context.Request.Header.Del("X-NLX-Request-Subject-Identifier")
			context.Request.Header.Del("X-NLX-Request-Application-Id")
			context.Request.Header.Del("X-NLX-Request-User-Id")
			context.Request.Header.Del(delegation.HTTPHeaderDelegator)
			context.Request.Header.Del(delegation.HTTPHeaderOrderReference)
		}

		context.Request.Header.Del("Proxy-Authorization")

		return next(context)
	}
}
