// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector, Role } from 'testcafe'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

export const adminUser = Role(`${baseUrl}/login`, async t => {
  const managementLoginButton = Selector('#login');
  const dexLoginText = Selector('#login');
  const dexPasswordText = Selector('#password');
  const dexSubmitLoginButton = Selector('#submit-login');
  const dexGrantAccessButton = Selector('button[type="submit"]');

  await t
    .click(managementLoginButton)
    .typeText(dexLoginText, 'admin@example.com')
    .typeText(dexPasswordText, 'password')
    .click(dexSubmitLoginButton)
    .click(dexGrantAccessButton)
});
