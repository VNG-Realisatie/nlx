// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { RequestLogger } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'

import { INWAY_NAME } from '../../environment'
import {
  getBaseUrl,
  getLocation,
  doAccessibilityTest,
  dismissAlertWithText,
  saveBrowserConsoleAndRequests,
} from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService } from './actions'

import servicesPage from './page-models/services'
import serviceDetailPage from './page-models/service-detail'
import addEditPage from './page-models/add-edit-service'

const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders: false,
  logResponseBody: true,
  stringifyResponseBody: true,
})

fixture`Services`
  .beforeEach(async (t) => {
    await t.useRole(adminUser)
    await waitForReact()
  })
  .afterEach(async (t) => saveBrowserConsoleAndRequests(t, logger.requests))
  .requestHooks(logger)

test('Services view accessibility test', async (t) => {
  await doAccessibilityTest(t)
})

test('Add and remove service', async (t) => {
  await createService()
  await dismissAlertWithText(t, 'service is toegevoegd')

  const {
    servicesList,
    getRowElementForService,
    alert,
    alertContent,
  } = servicesPage
  const serviceRow = await getRowElementForService(t.ctx.serviceName)

  await t.expect(getLocation()).eql(`${baseUrl}/services/${t.ctx.serviceName}`)
  await t.expect(servicesList.visible).ok()
  await t.expect(serviceRow.exists).ok()

  // Detail page
  const { editButton, closeButton, removeButton } = serviceDetailPage

  await t.click(serviceRow)
  await t.expect(editButton.visible).ok()
  await t.expect(closeButton.visible).ok()
  await t.expect(removeButton.visible).ok()
  await doAccessibilityTest(t)

  // Remove
  await removeService()
  await t.expect(getLocation()).eql(`${baseUrl}/services`)
  await t.expect(serviceRow.exists).notOk('', { timeout: 100 })
  await t.expect(alert.exists).ok
  await t.expect(alertContent.innerText).contains('service is verwijderd')
})

test('Edit service', async (t) => {
  await createService({ publishToCentralDirectory: true })
  await dismissAlertWithText(t, 'service is toegevoegd')

  const { alert } = servicesPage
  const { editButton } = serviceDetailPage

  // Edit from detail page
  await t.expect(alert.withText('Service nog niet benaderbaar').exists).ok()
  await t.click(editButton)
  await t
    .expect(getLocation())
    .eql(`${baseUrl}/services/${t.ctx.serviceName}/edit-service`)

  await doAccessibilityTest(t)
  await addEditPage.fillAndSubmitForm({
    inways: [INWAY_NAME],
  })
  await t.expect(alert.withText('service is bijgewerkt').exists).ok
  await dismissAlertWithText(t, 'service is bijgewerkt')

  // Approach with deeplink and cancel
  await t.navigateTo(`${baseUrl}/services/${t.ctx.serviceName}/edit-service`)
  await t.click(addEditPage.backButton)
  await t.expect(getLocation()).eql(`${baseUrl}/services/${t.ctx.serviceName}`)

  await removeService()
})
