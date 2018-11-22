module.exports = (env) => {
    let config

    if (env === 'prod') {
        config = require('./webpack/prod')
    } else {
        let dev = require('./webpack/dev')
        config = dev(env)
    }

    return config
}
