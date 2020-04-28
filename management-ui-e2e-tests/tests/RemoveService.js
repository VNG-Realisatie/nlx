// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'

import { adminUser } from "./roles"

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

const id = `${Math.round(Math.random() * 1000000)}`

fixture`ServiceDetails remove`
  .beforeEach(async t => {
    t.ctx.removableServiceName = `remove-service-${t.testRun.browserConnection.browserInfo.alias.replace(':', '_')}-${id}`
    const submitButton = Selector('button[type="submit"]')
    const alert = Selector('[role="alert"]')

    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/services/add-service`)
      .typeText('#name', t.ctx.removableServiceName)
      .typeText('#endpointURL', `${t.ctx.removableServiceName}.test:8000`)
      .typeText('#documentationURL', `${t.ctx.removableServiceName}.test:8000/docs`)
      .typeText('#apiSpecificationURL', `${t.ctx.removableServiceName}.test:8000/openapi.json`)
      .click(submitButton)
      .expect(alert.innerText).contains('De service is toegevoegd.')
      .navigateTo(`${baseUrl}/services/${t.ctx.removableServiceName}`)
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
