// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import addPage, { AUTHORIZATION_TYPE_NONE } from './page-objects/add-service'
import detailPage from './page-objects/service-detail'
import { t } from "testcafe"

const baseUrl = require('../getBaseUrl')()

function randomNumberString() {
  return `${Math.floor(Math.random() * 10000)}`
}

const testRunId = randomNumberString()

function getBrowserId() {
  return t.browser.alias.replace(':', '_')
}

export function generateServiceName() {
  return `service-e2e-${(getBrowserId())}-${testRunId}-${randomNumberString()}`
}

export async function createService() {
  t.ctx.serviceName = generateServiceName()
  await t
    .navigateTo(`${baseUrl}/services/add-service`)
  await addPage.fillAndSubmitForm({
    name: t.ctx.serviceName,
    endpointUrl: 'my-service.test:8000',
    documentationUrl: 'my-service.test:8000/docs',
    apiSpecificationUrl: 'my-service.test:8000/openapi.json',
    techSupportContact: 'tech@organization.test',
    publicSupportContact: 'public@organization.test',
    authorizationType: AUTHORIZATION_TYPE_NONE,
  })
  await t
    .expect(addPage.alert.innerText).contains('De service is toegevoegd.')

  return t.ctx.serviceName
}

export async function removeService() {
  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)

  await t
    .setNativeDialogHandler((type, text, url) => {
      if (text !== 'Wil je de service verwijderen?') {
        throw `Unexpected dialog text: ${text}`
      }
      return true
    })
    .click(detailPage.removeButton)

  await t
    .expect(detailPage.alert.innerText).contains('De service is verwijderd.')
  t.ctx.serviceName = null
}
