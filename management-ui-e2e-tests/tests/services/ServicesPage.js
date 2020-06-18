// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { getBaseUrl } from '../../utils'
import { adminUser } from '../roles'
import { createService, removeService} from './actions'

const baseUrl = getBaseUrl()

fixture `Services page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/services`)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Page title is visible', async t => {
  const pageTitle = Selector('h1')

  await t
    .expect(pageTitle.visible).ok()
    .expect(pageTitle.innerText).eql('Services')
})

test
  .before(async t => {
    await t.useRole(adminUser)
    await createService()
    await t.navigateTo(`${baseUrl}/services`)
    await waitForReact()
  })
  ('Service details are displayed', async t => {
    const { serviceName } = t.ctx

    const servicesList = Selector('[data-testid="services-list"]');
    const kentekenService = Selector('tr').withText(serviceName)
    const kentekenServiceColumns = kentekenService.find('td')

    const nameCell = kentekenServiceColumns.nth(0)
    const accessCell = kentekenServiceColumns.nth(1)

    await t
      .expect(servicesList.visible).ok()
      .expect(servicesList.find('tbody tr').count).gte(1) // until we have the delete option, we can't assert the exact amount of services

      .expect(nameCell.textContent).eql(serviceName)
      .expect(accessCell.textContent).eql('Open')
  })
  .after(async t => {
    await removeService()
  })
