const { getBaseUrl, isDebugging } = require('../environment')
const baseUrl = getBaseUrl(isDebugging())

describe('Home', () => {
    let page

    beforeAll(async () => {
        page = await global.__BROWSER__.newPage();

        await page.setBypassCSP(true)

        // page should be loaded
        await page.goto(`${baseUrl}/`, { waitUntil: 'load' })

        // services table should be rendered
        await page.waitForSelector('[data-test="services-table"]')
    })

    describe('clicking a service with a public_support_contact email address', () => {
        beforeAll(async () => {
            await page.goto(`${baseUrl}/?q=rdw`, { waitUntil: 'load' })

            // services table should be rendered
            await page.waitForSelector('[data-test="services-table"]')

            await page.click('[data-test="service-table-row"]')
        })

        describe('the detail pane', () => {
            let serviceDetailPane

            beforeAll(async () => {
                serviceDetailPane = await page.$eval('[data-test="service-detail-pane"]', e => e.innerHTML)
            })

            it('should display the support email address', async () => {
                const emailAddress = await page.$eval('[data-test="email-address-link"]', e => e.innerHTML)
                await page.screenshot({ path: 'screenshots/home.detail-pane-support-email-address.png' })
                expect(emailAddress).toBe('public-support@nlx.io')
            })
        })
    })
})
