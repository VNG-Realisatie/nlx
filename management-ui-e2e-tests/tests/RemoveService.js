// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { adminUser } from "./roles"
import { createService } from './services'
import page from './page-objects/service-detail'

const baseUrl = require('../getBaseUrl')()

fixture`ServiceDetails remove`
  .beforeEach(async t => {
    await t
      .useRole(adminUser)
    const serviceName = await createService()
    await t.navigateTo(`${baseUrl}/services/${serviceName}`)
  })

test('Removing a service', async t => {
  await t
    .setNativeDialogHandler((type, text, url) => {
      if (text !== 'Wil je de service verwijderen?') {
        throw `Unexpected dialog text: ${text}`
      }
      return true
    })
    .click(page.removeButton)
    .takeScreenshot()
    .expect(page.alert.visible).ok()
    .expect(page.alert.innerText).contains('De service is verwijderd.')
})
