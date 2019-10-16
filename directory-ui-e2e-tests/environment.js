module.exports.isDebugging = () =>
    process.env.NODE_ENV === 'debug'

module.exports.getBaseUrl = () => process.env.URL
