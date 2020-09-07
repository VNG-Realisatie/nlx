// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { RequestLogger } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { INWAY_NAME } from '../../environment'
import {
  getBaseUrl,
  getLocation,
  saveBrowserConsoleAndRequests,
} from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService } from './actions'
import addPage from './page-models/add-service'
import detailPage from './page-models/service-detail'

const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders: false,
  logResponseBody: true,
  stringifyResponseBody: true,
})

fixture`Edit Service page`
  .beforeEach(async (t) => {
    await t.useRole(adminUser)
    await t.navigateTo(`${baseUrl}/services`)
    await createService()
    await waitForReact()
  })
  .afterEach(async (t) => {
    await removeService()
    await saveBrowserConsoleAndRequests(t, logger.requests)
  })
  .requestHooks(logger)

test('Automated accessibility testing', async (t) => {
  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}/edit-service`)
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
})

test('Page title is visible', async (t) => {
  const pageTitle = addPage.title

  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}/edit-service`)
  await t.expect(pageTitle.visible).ok()
  await t.expect(pageTitle.innerText).eql('Service bewerken')
})

test('Updating the service', async (t) => {
  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)
  await t.click(detailPage.editButton)

  await addPage.fillAndSubmitForm({
    publishToCentralDirectory: false,
    inways: [INWAY_NAME],
  })

  // TODO: Test... submitting form should navigate to detail page automatically
  // await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)
  await t.expect(getLocation()).eql(`${baseUrl}/services/${t.ctx.serviceName}`)
  await t
    .expect(detailPage.published.innerText)
    .contains('Niet zichtbaar in centrale directory')
  await t.expect(detailPage.inways.innerText).eql('Inways1')
})

test('Show the missing inways warning', async (t) => {
  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)
  await t.click(detailPage.editButton)

  await addPage.fillAndSubmitForm({
    publishToCentralDirectory: true,
  })

  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}`)
})
