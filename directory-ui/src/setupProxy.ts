// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

// eslint-disable-next-line @typescript-eslint/no-var-requires
const { createProxyMiddleware } = require('http-proxy-middleware')

const getProxyUrl = (proxy?: string): string =>
  proxy || 'http://directory-inspection-api.shared.nlx.local:7902'

module.exports = function (app) {
  app.use(
    '/api',
    createProxyMiddleware({ target: getProxyUrl(process.env.PROXY) }),
  )
}
