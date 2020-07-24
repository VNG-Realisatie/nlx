// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { getBaseUrl, getLocation } from '../../utils'
import { adminUser } from '../roles'
import { createService } from './actions'
import page from './page-models/service-detail'
import { Selector } from "testcafe";

const baseUrl = getBaseUrl()

fixture`ServiceDetails remove`.beforeEach(async (t) => {
  await t.useRole(adminUser)
  const serviceName = await createService()
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

  await t.click(page.removeButton).takeScreenshot()

  await t.expect(getLocation()).eql(`${baseUrl}/services`)

  const serviceRemovedAlert = Selector('div[role="alert"]')
  await t.expect(serviceRemovedAlert.visible).ok()

  await t.expect(serviceRemovedAlert.withExactText('De service is verwijderd.')).ok()
})
