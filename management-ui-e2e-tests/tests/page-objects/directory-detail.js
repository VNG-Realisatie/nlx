// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe';

class Page {
  constructor() {
    this.closeButton = Selector('[data-testid="close-button"]', { visibilityCheck: true })
    this.detailHeader = Selector('[data-testid="directory-detail-header"]')
    this.requestAccess = Selector('[data-testid="request-access-section"]')
  }
}

export default new Page()
