// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

const { createProxyMiddleware } = require('http-proxy-middleware')

module.exports = function(app) {
    app.use(
        '/api',
        createProxyMiddleware({
            pathRewrite: {
                '^/api': '', // rewrite path
            },
            target: 'http://localhost:6010',
        }),
    )
}
