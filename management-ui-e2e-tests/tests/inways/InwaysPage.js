// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { RequestLogger, Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import {
  INWAY_NAME,
  INWAY_SELF_ADDRESS,
  INWAY_VERSION,
} from '../../environment'
import { getBaseUrl, saveBrowserConsoleAndRequests } from '../../utils'
import { adminUser } from '../roles'

const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders:    false,
  logResponseBody:       true,
  stringifyResponseBody: true,
})

fixture`Inways page`.beforeEach(async (t) => {
  await t.useRole(adminUser).navigateTo(`${baseUrl}/inways`)
  await waitForReact()
})
  .afterEach(async (t) =>
    saveBrowserConsoleAndRequests(t, logger.requests)
  ).requestHooks(logger)

test('Automated accessibility testing', async (t) => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
})

test('Page title is visible', async (t) => {
  const pageTitle = Selector('h1')

  await t.expect(pageTitle.visible).ok()
  await t.expect(pageTitle.innerText).eql('Inways')
})

test('Inway details are displayed', async (t) => {
  const inwaysList = Selector('[data-testid="inways-list"]')
  const initialService = Selector('tr').withText(INWAY_NAME)
  const initialServiceColumns = initialService.find('td')

  const nameCell = initialServiceColumns.nth(1)
  const hostnameCell = initialServiceColumns.nth(2)
  const selfAddressCell = initialServiceColumns.nth(3)
  const serviceCountCell = initialServiceColumns.nth(4)
  const versionCell = initialServiceColumns.nth(5)

  await t.expect(inwaysList.visible).ok()
  await t.expect(inwaysList.find('tbody tr').count).gte(1) // until we have the delete option, we can't assert the exact amount of services
  
  await t.expect(nameCell.textContent).eql(INWAY_NAME)
  await t.expect(hostnameCell.textContent).notEql('') // we only check for not empty, because the hostname is nondeterministic
  await t.expect(selfAddressCell.textContent).eql(INWAY_SELF_ADDRESS)
  await t.expect(Number(serviceCountCell.textContent)).typeOf('number')
  await t.expect(versionCell.textContent).eql(INWAY_VERSION)
})
