// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

const { createProxyMiddleware } = require('http-proxy-middleware')

const getProxyUrl = (proxy) =>
  proxy || 'http://directory.nlx-dev-directory.minikube/'

module.exports = function(app) {
  app.use(
    '/api',
    createProxyMiddleware({
      pathRewrite: {
        '^/api': '', // rewrite path
      },
      target: getProxyUrl(process.env.PROXY),
    }),
  )
}
