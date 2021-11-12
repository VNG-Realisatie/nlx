// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

// eslint-disable-next-line @typescript-eslint/no-var-requires
const { createProxyMiddleware } = require('http-proxy-middleware')

const getProxyUrl = (proxy) =>
  proxy || 'http://directory-api.shared.nlx.local:7905'

module.exports = function (app) {
  app.use(
    '/api',
    createProxyMiddleware({ target: getProxyUrl(process.env.PROXY) }),
  )
}
