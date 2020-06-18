// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { LOGIN_ORGANIZATION_NAME } from '../../environment'
import { getBaseUrl, getLocation } from '../../utils'
import { adminUser } from '../roles'
import page from './page-models/login'

const baseUrl = getBaseUrl()

fixture`Login page`.beforeEach(async (t) => {
  await t.navigateTo(`${baseUrl}/login`)
  await waitForReact()
})

test('Automated accessibility testing', async (t) => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
})

test('Login page contains all required elements', async (t) => {
  await t.expect(page.title.innerText).eql('Welkom')
  await t.expect(page.organizationName.exists).ok()
  await t.expect(page.organizationName.innerText).eql(LOGIN_ORGANIZATION_NAME)
  await t.expect(page.loginButton.visible).ok()
})

test('Base page redirects to Inways page when logged in', async (t) => {
  await t.useRole(adminUser).navigateTo(`${baseUrl}/`)
  await t.expect(getLocation()).contains('/inways')
})

test('Login page shows logout button when logged in', async (t) => {
  await t.useRole(adminUser).navigateTo(`${baseUrl}/login`)
  await t.expect(page.logoutButton.visible).ok()
})
