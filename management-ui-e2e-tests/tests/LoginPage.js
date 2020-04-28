// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector, ClientFunction } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'
import { adminUser } from './roles';

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

test('Welcome message is present', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Welkom')
});

test('Login button is present', async t => {
  await t
      .expect(Selector('#login').visible).ok()
})

test('Login', async t => {
  const getLocation = ClientFunction(() => document.location.href.toString());
  const managementLogoutButton = Selector('#logout');
  await t
    .useRole(adminUser)
    .expect(getLocation()).contains('/inways')

    .navigateTo(`${baseUrl}/login`)
    .expect(managementLogoutButton.visible).ok()
})
