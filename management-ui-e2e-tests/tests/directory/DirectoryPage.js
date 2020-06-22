// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import {
  DIRECTORY_ORGANIZATION_NAME,
  DIRECTORY_SERVICE_NAME,
  DIRECTORY_STATUS,
  DIRECTORY_API_SPECIFICATION_TYPE,
} from '../../environment'
import { getBaseUrl } from '../../utils'
import { adminUser } from '../roles'

const baseUrl = getBaseUrl()

fixture`Directory page`.beforeEach(async (t) => {
  await t.useRole(adminUser).navigateTo(`${baseUrl}/directory`)
  await waitForReact()
})

test('Automated accessibility testing', async (t) => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
})

test('Page title is visible', async (t) => {
  const pageTitle = Selector('h1')

  await t.expect(pageTitle.visible).ok()
  await t.expect(pageTitle.innerText).eql('Directory')
})

test('Directory details are displayed', async (t) => {
  const directoryList = Selector('[data-testid="directory-services"]')
  const initialService = Selector('tr').withText(DIRECTORY_SERVICE_NAME)
  const initialServiceColumns = initialService.find('td')

  const organizationNameCell = initialServiceColumns.nth(0)
  const serviceNameCell = initialServiceColumns.nth(1)
  const statusTitle = initialServiceColumns.nth(2).find('svg title')
  const apiSpecificationTypeCell = initialServiceColumns.nth(3)
  const accessCell = initialServiceColumns.nth(4)

  await t.expect(directoryList.visible).ok()
  await t.expect(directoryList.find('tbody tr').count).gte(2)

  await t.expect(organizationNameCell.textContent).eql(DIRECTORY_ORGANIZATION_NAME)
  await t.expect(serviceNameCell.textContent).eql(DIRECTORY_SERVICE_NAME)
  await t.expect(statusTitle.textContent).eql(DIRECTORY_STATUS)
  await t.expect(apiSpecificationTypeCell.textContent).eql(DIRECTORY_API_SPECIFICATION_TYPE)
  await t.expect(accessCell.find('svg[data-testid="request-access"]')).ok()
})
