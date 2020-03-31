// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector, Role } from 'testcafe'
import { axeCheck, createReport } from 'axe-testcafe'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl();

fixture `Login page`
  .page `${baseUrl}`
  .beforeEach(async (t) => {
    await t.useRole(Role.anonymous())
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
  const managementLoginButton = Selector('#login');
  const managementLogoutButton = Selector('#logout');
  const dexLoginText = Selector('#login');
  const dexPasswordText = Selector('#password');
  const dexSubmitLoginButton = Selector('#submit-login');
  const dexGrantAccessButton = Selector('button[type="submit"]');
  await t
      .expect(managementLoginButton.visible).ok()
      .click(managementLoginButton)

      .expect(dexSubmitLoginButton.visible).ok()
      .typeText(dexLoginText, 'admin@example.com')
      .typeText(dexPasswordText, 'password')
      .click(dexSubmitLoginButton)

      .expect(dexGrantAccessButton.visible).ok()
      .click(dexGrantAccessButton)

      .expect(managementLogoutButton.visible).ok()
})
