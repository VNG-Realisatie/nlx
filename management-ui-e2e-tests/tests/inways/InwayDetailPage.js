// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import {
  INWAY_NAME,
  INWAY_SELF_ADDRESS,
  INWAY_VERSION,
} from '../../environment'
import { getBaseUrl, getLocation } from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService } from '../services/actions'
import page from './page-models/inway-detail'

const baseUrl = getBaseUrl()

fixture`InwayDetails page`.beforeEach(async (t) => {
  await t.useRole(adminUser).navigateTo(`${baseUrl}/inways/${INWAY_NAME}`)
  await waitForReact()
})

test('Automated accessibility testing', async (t) => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
})

test('Inway details are visible', async (t) => {
  await t.expect((await page.inwayName.innerText).trim()).eql(INWAY_NAME)
  await t.expect(page.inwaySpecs.innerText).contains(INWAY_SELF_ADDRESS)
  await t.expect(page.inwaySpecs.innerText).contains(INWAY_VERSION)
})

// First create a service, then check if it appears on inway and we can link to it
test
  .before(async (t) => {
    await t.useRole(adminUser)
    await createService({ inways: [INWAY_NAME] })

    await t.navigateTo(`${baseUrl}/inways/${INWAY_NAME}`)
    await waitForReact()
  })('Links to connected service detail page', async (t) => {
    const { serviceName } = t.ctx // set by `createService`

    await t.click(page.services)
    await t.click(Selector('td').withText(serviceName))
    await t.expect(getLocation()).contains(`${baseUrl}/services/${serviceName}`)
  })
  .after(async (t) => {
    await removeService()
  })

// In IE11 the transition doesn't always complete when directly navigating to detail
// So X may not be visible/clickable
test
  .before(async (t) => {
    await t.useRole(adminUser).navigateTo(`${baseUrl}/inways`)
  })
  ('Opens and closes details view', async (t) => {
    const inwayRow = Selector('tr').withText(INWAY_NAME)

    await t.click(inwayRow)
    await t.expect(page.closeButton.exists).ok()
    await t.click(page.closeButton)
    await t.expect(getLocation()).contains(`${baseUrl}/inways`)
  })
