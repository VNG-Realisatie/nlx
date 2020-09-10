// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { RequestLogger } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'

import { LOGIN_ORGANIZATION_NAME } from '../../environment'
import {
  getBaseUrl,
  getLocation,
  doAccessibilityTest,
  saveBrowserConsoleAndRequests,
} from '../../utils'
import { adminUser } from '../roles'
import loginPage from './page-models/login'
import dexPage from './page-models/dex'

const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders: false,
  logResponseBody: true,
  stringifyResponseBody: true,
})

fixture`Login`
  .beforeEach(async (t) => {
    await t.navigateTo(`${baseUrl}/login`)
    await waitForReact()
  })
  .afterEach(async (t) => {
    saveBrowserConsoleAndRequests(t, logger.requests)
  })
  .requestHooks(logger)

test('Login page is accessible', async (t) => {
  await doAccessibilityTest(t)
})

test('Login page contains shows organization name', async (t) => {
  await t.expect(loginPage.organizationName.exists).ok()
  await t
    .expect(loginPage.organizationName.innerText)
    .eql(LOGIN_ORGANIZATION_NAME)
})

test('Successful login redirects to Inways page', async (t) => {
  await t.useRole(adminUser)
  await t.expect(getLocation()).contains('/inways')
})

test('Login attempt with invalid credentials shows error', async (t) => {
  await t.setTestSpeed(0.5).click(loginPage.loginButton)

  await t
    .typeText(dexPage.loginText, 'abc')
    .typeText(dexPage.passwordText, 'abc')
    .click(dexPage.submitLoginButton)
  await t.expect(dexPage.error.exists).ok()
})
