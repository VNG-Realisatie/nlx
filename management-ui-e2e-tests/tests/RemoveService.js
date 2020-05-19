// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'

import { adminUser } from "./roles"
import { createService } from './services'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

fixture`ServiceDetails remove`
  .beforeEach(async t => {
    await t
      .useRole(adminUser)
    const serviceName = await createService()
    await t.navigateTo(`${baseUrl}/services/${serviceName}`)
  })

test('Removing a service', async t => {
  const removeButton = Selector('[data-testid="remove-service"]')
  const alert = Selector('[role="alert"]')

  await t
    .setNativeDialogHandler((type, text, url) => {
      if (text !== 'Wil je de service verwijderen?') {
        throw `Unexpected dialog text: ${text}`
      }
      return true
    })
    .click(removeButton)
    .takeScreenshot()
    .expect(alert.visible).ok()
    .expect(alert.innerText).contains('De service is verwijderd.')
})
