// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { Selector } from 'testcafe'

class Page {
  constructor() {
    this.title = Selector('h1')

    this.organizationName = Selector('[data-testid="organizationName"]')
    this.loginButton = Selector('[data-testid="login"]')
    this.logoutButton = Selector('[data-testid="logout"]')
  }
}

export default new Page()
