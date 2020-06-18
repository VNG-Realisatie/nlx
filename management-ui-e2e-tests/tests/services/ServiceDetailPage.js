// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { getBaseUrl, getLocation } from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService} from './actions'
import page from './page-models/service-detail'

const baseUrl = getBaseUrl()

fixture`ServiceDetails page`
  .beforeEach(async t => {
    await t.useRole(adminUser)
    const serviceName = await createService()
    
    await t.navigateTo(`${baseUrl}/services/${serviceName}`)
    await waitForReact()
  })
  .afterEach(async () => {
    await removeService()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Service details are visible', async t => {
  await t.expect((await page.serviceName.innerText).trim()).eql(t.ctx.serviceName)
  await t.expect(page.published.visible).ok()
})

test('Close navigates to the ServicesPage', async t => {
  await t.expect(page.closeButton.exists).ok()
  await t.click(page.closeButton)
  await t.expect(getLocation()).contains(`${baseUrl}/services`);
})
