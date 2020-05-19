// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import addPage, { AUTHORIZATION_TYPE_NONE } from './page-objects/add-service'
import { Selector, t } from "testcafe"

import getBaseUrl from '../getBaseUrl'

const baseUrl = getBaseUrl()

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
  const removeButton = Selector('[data-testid="remove-service"]')
  const alert = Selector('[role="alert"]')

  await t.navigateTo(`${baseUrl}/services/${(t.ctx.serviceName)}`)

  await t
    .setNativeDialogHandler((type, text, url) => {
      if (text !== 'Wil je de service verwijderen?') {
        throw `Unexpected dialog text: ${text}`
      }
      return true
    })
    .click(removeButton)

  await t
    .expect(alert.innerText).contains('De service is verwijderd.')
  t.ctx.serviceName = null
}