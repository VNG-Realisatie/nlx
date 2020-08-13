// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

const { createProxyMiddleware } = require('http-proxy-middleware')

const getProxyUrl = (proxy) =>
  proxy || 'http://directory-inspection-api.shared.nlx.local:7902'

module.exports = function (app) {
  app.use(
    '/api',
    createProxyMiddleware({ target: getProxyUrl(process.env.PROXY) }),
  )
}
