// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import getLocation from '../getLocation'
import { INWAY_NAME, INWAY_SELF_ADDRESS, INWAY_VERSION } from './environment'
import { adminUser } from './roles'
import page from './page-objects/inway-detail'

const baseUrl = require('../getBaseUrl')()

fixture `InwayDetails page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/inways/${INWAY_NAME}`)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Inway details are visible', async t => {
  await t
    .expect((await page.inwayName.innerText).trim()).eql(INWAY_NAME)
    .expect(page.inwaySpecs.innerText).contains(INWAY_SELF_ADDRESS)
    .expect(page.inwaySpecs.innerText).contains(INWAY_VERSION)
})

test('Lists connected services', async t => {
  const services = page.services
  const toggleText = await services.innerText

  await t
    .expect(toggleText.substring(toggleText.length - 1)).eql('1')
})

test('Links to service detail', async t => {
  await t
    .click(page.services)
    .click(Selector('td').withText('kentekenregister'))
    .expect(getLocation()).contains(`${baseUrl}/services/kentekenregister`);
})

test('Close navigates to the InwaysPage', async t => {
  await t
    .expect(page.closeButton.exists).ok()
    .click(page.closeButton)
    .expect(getLocation()).contains(`${baseUrl}/inways`);
})
