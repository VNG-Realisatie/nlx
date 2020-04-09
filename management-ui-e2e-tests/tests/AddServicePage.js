// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'
import { adminUser } from './roles'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

fixture `Add Service page`
  .page(`${baseUrl}/services/add-service`)
  .beforeEach(async (t) => {
    await t.useRole(adminUser)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Page title is visible', async t => {
  const pageTitle = Selector('h1')

  await t
    .expect(pageTitle.visible).ok()
    .expect(pageTitle.innerText).eql('Nieuwe service toevoegen')
})

test('Adding a new service', async t => {
  const submitButton = Selector('button[type="submit"]')
  const alert = Selector('[role="alert"]')
  const nameError = Selector('[data-testid="name-error"]')

  await t
    .typeText('#endpointURL', 'my-service.test:8000')
    .typeText('#documentationURL', 'my-service.test:8000/docs')
    .typeText('#apiSpecificationURL', 'my-service.test:8000/openapi.json')
    .click('#internal')
    .typeText('#techSupportContact', 'tech@organization.test')
    .typeText('#publicSupportContact', 'public@organization.test')
    .click('#authorizationModeNone')
    .click(submitButton)

    .expect(nameError.innerText).contains('Dit veld is verplicht.')
    .typeText('#name', 'my-service')
    .click(submitButton)

    .expect(alert.innerText).contains('De service is toegevoegd.')
})
