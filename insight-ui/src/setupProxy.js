// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

const proxy = require('http-proxy-middleware')

const getProxyUrl = (proxy) =>
    proxy ? proxy : 'http://directory.dev.nlx.minikube:30080/'

module.exports = function(app) {
    const proxyUrl = getProxyUrl(process.env.PROXY)
    app.use(
        proxy('/api', {
            target: proxyUrl,
            secure: false,
            changeOrigin: true,
        }),
    )
}
