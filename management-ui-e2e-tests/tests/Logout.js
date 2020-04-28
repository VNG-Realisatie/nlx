import { adminUser } from './roles'
import { waitForReact } from 'testcafe-react-selectors'
import { Selector, Role } from 'testcafe'
import { axeCheck, createReport } from 'axe-testcafe'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

fixture`Logout`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/`)
    await waitForReact()
  })

test('Logging out should navigate to the login page', async t => {
  const userMenuButton = Selector('[aria-label="Account menu"]')
  const userMenuLogoutButton = Selector('button').withText('Uitloggen')
  const loginButton = Selector('#login')

  await t
    .expect(userMenuButton.visible).ok()
    .click(userMenuButton)
    .expect(userMenuLogoutButton.visible).ok()
    .click(userMenuLogoutButton)
    .wait(500)
    .expect(loginButton.visible).ok()
})
