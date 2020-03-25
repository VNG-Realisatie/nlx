const { analyzeAccessibility } = require('../accessibility')
const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl();

describe('LoginPage', () => {
    beforeAll(async () => {
        await page.setBypassCSP(true)

        // page should be loaded
        await page.goto(`${baseUrl}/login`, { waitUntil: 'load' })

        // page content should be rendered
        await page.waitForSelector('p')
    })

    it('should not have accessibility issues', async () => {
        const accessibilityReport = await analyzeAccessibility(page, `login.accessibility.png`)
        expect(accessibilityReport).toHaveNoAccessibilityIssues();
    })
})
