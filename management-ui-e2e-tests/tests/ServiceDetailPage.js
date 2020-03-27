// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { axeCheck, createReport } from 'axe-testcafe'
import { adminUser } from "./roles";
import getLocation from '../getLocation'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

fixture `ServiceDetails page`
  .page(`${baseUrl}/services/kentekenregister`)
  .beforeEach(async (t) => {
    await t.useRole(adminUser)
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Service details are visible', async t => {
  const serviceName = Selector('[data-testid="service-name"]');
  const servicePublished = Selector('[data-testid="service-published"]');

  await t
    .expect((await serviceName.innerText).trim()).eql('kentekenregister')
    .expect(servicePublished.visible).ok()
})

test('Close navigates to the ServicesPage', async t => {
  const closeButton = Selector('[data-testid="close-button"]');

  await t
    .click(closeButton)
    .expect(getLocation()).contains(`${baseUrl}/services`);
})
