// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { Selector } from 'testcafe'

class Page {
  constructor() {
    // Wait until visible. Drawer animation can take some time and click interaction requires visibility
    this.closeButton = Selector('[data-testid="close-button"]', {
      visibilityCheck: true,
    })
    this.inwayName = Selector('[data-testid="gateway-name"]')
    this.gatewayType = Selector('[data-testid="gateway-type"]')
    this.inwaySpecs = Selector('[data-testid="inway-specs"]')
    this.services = Selector('[data-testid="inway-services"]')
    this.servicesList = Selector('[data-testid="inway-services-list"]')
  }
}

export default new Page()
