// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { RequestLogger, Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { getBaseUrl, saveBrowserConsoleAndRequests } from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService } from './actions'

const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders: false,
  logResponseBody: true,
  stringifyResponseBody: true,
})

fixture`Services page`
  .beforeEach(async (t) => {
    await t.useRole(adminUser).navigateTo(`${baseUrl}/services`)
    await waitForReact()
  })
  .afterEach(async (t) => saveBrowserConsoleAndRequests(t, logger.requests))
  .requestHooks(logger)

test('Automated accessibility testing', async (t) => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
})

test('Page title is visible', async (t) => {
  const pageTitle = Selector('h1')

  await t.expect(pageTitle.visible).ok()
  await t.expect(pageTitle.innerText).eql('Services')
})

test
  .before(async (t) => {
    await t.useRole(adminUser)
    await createService({ publishToCentralDirectory: true })
    await t.navigateTo(`${baseUrl}/services`)
    await waitForReact()
  })('Service details are displayed', async (t) => {
    const { serviceName } = t.ctx

    const servicesList = Selector('[data-testid="services-list"]')
    const service = Selector('tr').withText(serviceName)
    const serviceColumns = service.find('td')

    const nameCell = serviceColumns.nth(0)
    const accessCell = serviceColumns.nth(1)
    const alertColumn = serviceColumns.nth(2)

    await t.expect(servicesList.visible).ok()
    // in a concurrent test we can't assert the exact amount of services, but at least the service in this test should be there
    await t.expect(servicesList.find('tbody tr').count).gte(1)

    await t.expect(nameCell.textContent).eql(serviceName)
    await t.expect(accessCell.textContent).eql('Open')
    await t.expect(alertColumn.textContent).eql('Service nog niet benaderbaar')
  })
  .after(async (t) => {
    await removeService()
  })
