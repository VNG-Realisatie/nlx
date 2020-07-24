// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { t } from 'testcafe'

import { getBaseUrl, getLocation } from '../../utils'
import addPage, { AUTHORIZATION_TYPE_NONE } from './page-models/add-service'
import detailPage from './page-models/service-detail'
import servicesPage from './page-models/services'

const baseUrl = getBaseUrl()

function randomNumberString() {
  return `${Math.floor(Math.random() * 10000)}`
}

const testRunId = randomNumberString()

function getBrowserId() {
  return t.browser.alias.replace(':', '_')
}

export function generateServiceName() {
  return `service-e2e-${getBrowserId()}-${testRunId}-${randomNumberString()}`
}

export async function createService(serviceProperties = {}) {
  t.ctx.serviceName = generateServiceName()

  await t.navigateTo(`${baseUrl}/services/add-service`)

  await addPage.fillAndSubmitForm({
    name: t.ctx.serviceName,
    endpointUrl: 'my-service.test:8000',
    documentationUrl: 'my-service.test:8000/docs',
    apiSpecificationUrl: 'my-service.test:8000/openapi.json',
    techSupportContact: 'tech@organization.test',
    publicSupportContact: 'public@organization.test',
    authorizationType: AUTHORIZATION_TYPE_NONE,
    publishToCentralDirectory: false,
    ...serviceProperties,
  })

  return t.ctx.serviceName
}

export async function removeService() {
  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)

  await t.setNativeDialogHandler(() => true)

  await detailPage.removeService()
  await t.expect(getLocation()).eql(servicesPage.url)

  t.ctx.serviceName = null
}
