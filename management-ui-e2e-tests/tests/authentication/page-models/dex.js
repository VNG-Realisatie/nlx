// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { Selector } from 'testcafe'

class Page {
  constructor() {
    this.loginText = Selector('#login')
    this.passwordText = Selector('#password')
    this.submitLoginButton = Selector('#submit-login')
    this.grantAccessButton = Selector('button[type="submit"]')
    this.error = Selector('#login-error')
  }
}

export default new Page()
