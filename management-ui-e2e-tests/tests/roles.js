// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { Role } from 'testcafe'

import { getBaseUrl } from '../utils'
import { LOGIN_USERNAME, LOGIN_PASSWORD } from '../environment'
import loginPage from './authentication/page-models/login'
import dexPage from './authentication/page-models/dex'

const baseUrl = getBaseUrl()

export const adminUser = Role(
  `${baseUrl}/login`,
  async (t) => {
    await t
      .setTestSpeed(0.5)
      .click(loginPage.loginButton)
      .typeText(dexPage.loginText, LOGIN_USERNAME)
      .typeText(dexPage.passwordText, LOGIN_PASSWORD)
      .click(dexPage.submitLoginButton)
      .click(dexPage.grantAccessButton)
      .setTestSpeed(1)
  },
  { preserveUrl: true },
)
