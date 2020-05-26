// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe';

class Page {
  constructor() {
    this.serviceName = Selector('[data-testid="service-name"]')
    this.published = Selector('[data-testid="service-published"]')
    this.editButton = Selector('[data-testid="edit-button"]')
    this.closeButton = Selector('[data-testid="close-button"]')
    this.removeButton = Selector('[data-testid="remove-service"]')
    this.inways = Selector('[data-testid="service-inways"]')
    this.alert = Selector('[role="alert"]')

  }
}

export default new Page()
