// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'
import { adminUser } from './roles'
import { LOGIN_ORGANIZATION_NAME } from './environment'
import page from './page-objects/login'
import getLocation from '../getLocation'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl();

fixture `Login page`
  .beforeEach(async (t) => {
    await t.navigateTo(`${baseUrl}/login`)
    await waitForReact();
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Login page contains all required elements', async t => {
  await t
    .expect(page.title.innerText).eql('Welkom')

    .expect(page.organizationName.exists).ok()
    .expect(page.organizationName.innerText).eql(LOGIN_ORGANIZATION_NAME)

    .expect(page.loginButton.visible).ok()
})

test('Base page redirects to Inways page when logged in', async t => {
  await t
    .useRole(adminUser)

    .navigateTo(`${baseUrl}/`)
    .expect(getLocation()).contains('/inways')
})

test('Login page shows logout button when logged in', async t => {
  await t
    .useRole(adminUser)

    .navigateTo(`${baseUrl}/login`)
    .expect(page.logoutButton.visible).ok()
})
