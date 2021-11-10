// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

/* eslint-disable @typescript-eslint/no-var-requires */

const { createProxyMiddleware } = require('http-proxy-middleware')

const getProxyUrl = (proxy) =>
  proxy || 'http://management-api.organization-a.nlx.local:7912'

module.exports = function (app) {
  app.use(
    '/api',
    createProxyMiddleware({ target: getProxyUrl(process.env.PROXY) }),
  )
  app.use(
    '/oidc',
    createProxyMiddleware({ target: getProxyUrl(process.env.PROXY) }),
  )
  app.use(
    '/basic-auth',
    createProxyMiddleware({ target: getProxyUrl(process.env.PROXY) }),
  )
}
