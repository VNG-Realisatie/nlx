// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { INWAY_NAME } from '../../environment'
import { getBaseUrl } from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService } from './actions'
import addPage from './page-models/add-service'
import detailPage from './page-models/service-detail'

const baseUrl = getBaseUrl()

fixture`Edit Service page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/services`)
    await createService()
    await waitForReact()
  })
  .afterEach(async (t) => {
    await removeService()
  })

test('Automated accessibility testing', async t => {
  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}/edit-service`)
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Page title is visible', async t => {
  const pageTitle = addPage.title

  await t
    .navigateTo(`${baseUrl}/services/${t.ctx.serviceName}/edit-service`)
    .expect(pageTitle.visible).ok()
    .expect(pageTitle.innerText).eql('Service bewerken')
})

test('Updating the service', async t => {
  await t
    .navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)
    .click(detailPage.editButton)
    await addPage.fillAndSubmitForm({publishToCentralDirectory: false, inways: [INWAY_NAME]})
    await t
    .expect(addPage.alert.innerText).contains('De service is bijgewerkt.')
    .navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)
    .expect(detailPage.published.innerText).contains( 'Niet zichtbaar in centrale directory')
    .expect(detailPage.inways.innerText).eql( 'Inways1')
})
