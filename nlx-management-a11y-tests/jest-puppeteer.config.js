const puppeteer = require('puppeteer')

const getExecutablePath = () =>
    process.env.PUPPETEER_SKIP_CHROMIUM_DOWNLOAD ?
        'google-chrome-unstable' :
        puppeteer.executablePath()

// Prevent downloading Chrome for every E2E CI job.
// We can use the pre-installed 'google-chrome-unstable' executable instead.
// Via https://github.com/GoogleChrome/puppeteer/blob/de82b87cfa1c637787a09b90266318905ae16f42/docs/troubleshooting.md#running-puppeteer-in-docker
module.exports = {
    launch: {
        executablePath: getExecutablePath()
    }
}
