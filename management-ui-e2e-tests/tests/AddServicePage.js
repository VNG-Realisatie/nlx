// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'
import { adminUser } from './roles'
import page, { AUTHORIZATION_TYPE_NONE } from './page-objects/add-service'
import { INWAY_NAME } from './environment'
import { generateServiceName, removeService } from './services'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

fixture `Add Service page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/services/add-service`)
    await waitForReact()
  })
  .afterEach(async (t) => {
    if (t.ctx.serviceName) {
      await removeService()
    }
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Page title is visible', async t => {
  await t
    .expect(page.title.visible).ok()
    .expect(page.title.innerText).eql('Nieuwe service toevoegen')
})

test('Adding a new service', async t => {
  await page.fillAndSubmitForm({
    endpointUrl: 'my-service.test:8000',
    documentationUrl: 'my-service.test:8000/docs',
    apiSpecificationUrl: 'my-service.test:8000/openapi.json',
    publishToCentralDirectory: true,
    techSupportContact: 'tech@organization.test',
    publicSupportContact: 'public@organization.test',
    authorizationType: AUTHORIZATION_TYPE_NONE,
    inways: [INWAY_NAME],
  })
  await t.expect(page.nameFieldError.innerText).contains('Dit veld is verplicht.')

  t.ctx.serviceName = generateServiceName()
  await page.fillAndSubmitForm({ name: t.ctx.serviceName })
  await t.expect(page.alert.innerText).contains('De service is toegevoegd.')
})
