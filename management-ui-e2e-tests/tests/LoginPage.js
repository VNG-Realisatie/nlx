// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'

const makeUrl = require('../utils/makeUrl')

fixture `Login page`
  .page(makeUrl())

test('Welcome message is present', async t => {
  await t
    .expect(Selector('h1').innerText).eql('Welcome')
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
