// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

const { createProxyMiddleware } = require('http-proxy-middleware')

module.exports = function(app) {
    app.use(
        '/api',
        createProxyMiddleware({
            target: 'http://localhost:6017',
        }),
    )
}
