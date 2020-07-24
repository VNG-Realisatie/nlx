// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { Selector, t } from 'testcafe'
import { getBaseUrl } from "../../../utils";

const baseUrl = getBaseUrl()

class Page {
  constructor() {
    this.url = `${baseUrl}/services`
    this.servicesList = Selector('[data-testid="services-list"]')
    this.alert = Selector('div[role="alert"]')
    this.alertContent = this.alert.find('[data-testid="content"]')
  }

  async getRowElementForService(serviceName) {
    return this.servicesList.find('tr').withText(serviceName)
  }
}

export default new Page()
