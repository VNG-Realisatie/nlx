// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'
import { adminUser } from './roles'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

async function createService(t) {
  const serviceName = `edit-service-e2e-${Math.round(Math.random() * 100000)}`
  const submitButton = Selector('button[type="submit"]')
  const alert = Selector('[role="alert"]')
  await t
    .navigateTo(`${baseUrl}/services/add-service`)
    .typeText('#endpointURL', 'my-service.test:8000')
    .typeText('#documentationURL', 'my-service.test:8000/docs')
    .typeText('#apiSpecificationURL', 'my-service.test:8000/openapi.json')
    .typeText('#techSupportContact', 'tech@organization.test')
    .typeText('#publicSupportContact', 'public@organization.test')
    .click('#authorizationModeNone')
    .typeText('#name', serviceName)
    .click(submitButton)
    .expect(alert.innerText).contains('De service is toegevoegd.')

  return serviceName
}

fixture`Edit Service page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/services`)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const serviceName = await createService(t)
  await t.navigateTo(`${baseUrl}/services/${serviceName}/edit-service`)

  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Page title is visible', async t => {
  const serviceName = await createService(t)
  const pageTitle = Selector('h1')

  await t
    .navigateTo(`${baseUrl}/services/${serviceName}/edit-service`)
    .expect(pageTitle.visible).ok()
    .expect(pageTitle.innerText).eql('Bestaande service bijwerken')
})

test('Updating the service', async t => {
  const serviceName = await createService(t)
  const submitButton = Selector('button[type="submit"]')
  const alert = Selector('[role="alert"]')
  const editButton = Selector('[data-testid="edit-button"]')
  await t
    .navigateTo(`${baseUrl}/services/${serviceName}`)
    .click(editButton)
    .click('#publishedInDirectory')
    .click(submitButton)
    .expect(alert.innerText).contains('De service is bijgewerkt.')
    .navigateTo(`${baseUrl}/services/${serviceName}`)
    .expect(Selector('[data-testid="service-published"]').innerText).contains( 'Niet zichtbaar in centrale directory')
})
