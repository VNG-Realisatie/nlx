// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector, Role } from 'testcafe'
import loginPage from './page-objects/login'

const baseUrl = require('../getBaseUrl')()

export const adminUser = Role(`${baseUrl}/login`, async t => {
  const managementLoginButton = loginPage.loginButton;
  const dexLoginText = Selector('#login');
  const dexPasswordText = Selector('#password');
  const dexSubmitLoginButton = Selector('#submit-login');
  const dexGrantAccessButton = Selector('button[type="submit"]');

  await t
    .setTestSpeed(0.5)
    .click(managementLoginButton)
    .typeText(dexLoginText, 'admin@example.com')
    .typeText(dexPasswordText, 'password')
    .click(dexSubmitLoginButton)
    .click(dexGrantAccessButton)
}, { preserveUrl: true });
