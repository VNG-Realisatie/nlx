// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { getBaseUrl, getLocation } from '../../utils'
import { adminUser } from '../roles'
import { createService } from './actions'
import page from './page-models/service-detail'
import servicesPage from './page-models/services'

const baseUrl = getBaseUrl()

fixture`ServiceDetails remove`.beforeEach(async (t) => {
  await t.useRole(adminUser)
  const serviceName = await createService()
  t.ctx.serviceName = serviceName
  await t.navigateTo(`${baseUrl}/services/${serviceName}`)
})

test('Removing a service', async (t) => {
  await t
    .setNativeDialogHandler((type, text, url) => {
      if (text !== 'Wil je de service verwijderen?') {
        throw `Unexpected dialog text: ${text}`
      }
      return true
    })

  await page.removeService()

  await t.expect(getLocation()).eql(servicesPage.url)

  const serviceName = t.ctx.serviceName
  const serviceRow = await servicesPage.getRowElementForService(serviceName)
  await t.expect(serviceRow.exists).notOk()

  const serviceRemovedAlert = servicesPage.alert
  await t.expect(serviceRemovedAlert.visible).ok()
  await t.expect(servicesPage.alertContent.withExactText('De service is verwijderd.').exists).ok()
})
