// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { adminUser } from "./roles"
import getLocation from '../getLocation'
import page from './page-objects/service-detail'

const baseUrl = require('../getBaseUrl')()

fixture `ServiceDetails page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/services/kentekenregister`)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Service details are visible', async t => {
  await t
    .expect((await page.serviceName.innerText).trim()).eql('kentekenregister')
    .expect(page.published.visible).ok()
})

test('Close navigates to the ServicesPage', async t => {
  await t
    .expect(page.closeButton.exists).ok()
    .click(page.closeButton)
    .expect(getLocation()).contains(`${baseUrl}/services`);
})
