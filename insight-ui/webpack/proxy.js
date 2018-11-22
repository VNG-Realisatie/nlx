/**
 * Defining different proxy settings
 * based on provided environment prop
 * @param {string} env expected options are: local, test, demo, acc
 */
module.exports = (env) => {
    switch (env.toLowerCase()) {
        case 'local':
            return {
                '/api': {
                    target: 'http://directory.dev.nlx.minikube:30080',
                    secure: false,
                    changeOrigin: true,
                },
            }
        case 'dev':
            return {
                '/api': {
                    target: 'https://directory.test.nlx.io/',
                    secure: false,
                    changeOrigin: true,
                },
            }
        case 'demo':
            return {
                '/api': {
                    target: 'https://directory.demo.nlx.io/',
                    secure: false,
                    changeOrigin: true,
                },
            }
        // break;
        case 'acc':
            return {
                '/api': {
                    target: 'https://directory.acc.nlx.io/',
                    secure: false,
                    changeOrigin: true,
                },
            }
        // break;
        default:
            return null
    }
}
