// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { RequestLogger, Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'

import { INWAY_NAME } from '../../environment'
import {
  getBaseUrl,
  getLocation,
  doAccessibilityTest,
  saveBrowserConsoleAndRequests,
} from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService } from '../services/actions'

import inwayDetailPage from './page-models/inway-detail'

const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders: false,
  logResponseBody: true,
  stringifyResponseBody: true,
})

fixture`Inways`
  .beforeEach(async (t) => {
    await t.useRole(adminUser).navigateTo(`${baseUrl}/inways`)
    await waitForReact()
  })
  .afterEach(async (t) => saveBrowserConsoleAndRequests(t, logger.requests))
  .requestHooks(logger)

test('Inway overview accessibility test', async (t) => {
  await doAccessibilityTest(t)
})

test('Inway details are displayed and can be closed', async (t) => {
  const inwaysList = Selector('[data-testid="inways-list"]')
  const initialService = Selector('tr').withText(INWAY_NAME)

  await t.expect(inwaysList.visible).ok()
  await t.click(initialService)

  await t.expect(getLocation()).eql(`${baseUrl}/inways/${INWAY_NAME}`)
  await doAccessibilityTest(t)

  await t.click(inwayDetailPage.closeButton)
  await t.expect(getLocation()).eql(`${baseUrl}/inways`)
})

// First create a service, then check if it appears on inway and we can link to it
test
  .before(async (t) => {
    await t.useRole(adminUser)
    await createService({ inways: [INWAY_NAME] })

    await t.navigateTo(`${baseUrl}/inways/${INWAY_NAME}`)
    await waitForReact()
  })(
    'Deeplink to inway page and go to connected service detail page',
    async (t) => {
      const { serviceName } = t.ctx // set by `createService`

      await t.click(inwayDetailPage.services)
      await t.click(Selector('td').withText(serviceName))
      await t
        .expect(getLocation())
        .contains(`${baseUrl}/services/${serviceName}`)
    },
  )
  .after(async (t) => {
    await removeService()
  })
