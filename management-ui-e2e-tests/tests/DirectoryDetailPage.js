// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import getLocation from '../getLocation'
import { adminUser } from './roles'
import page from './page-objects/directory-detail'

const baseUrl = require('../getBaseUrl')()

fixture `DirectoryDetails page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/directory/BRP/basisregistratie`)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Directory service details are visible', async t => {
  const detailHeaderContent = page.detailHeader.textContent
  
  await t
    .expect(detailHeaderContent).contains('basisregistratie')
    .expect(detailHeaderContent).contains('BRP')
    .expect(detailHeaderContent).contains('OpenAPI3')
    .expect(detailHeaderContent).contains('Beschikbaar')
})

test('Request access button exists', async t => {
  const button = page.requestAccess.find('button')

  await t
    .expect(button.exists).ok()
})

test('Request access can be cancelled in dialog', async t => {
  const button = page.requestAccess.find('button')

  await t
    .setNativeDialogHandler((type) => {
      if (type === 'confirm') return false
    })
    .click(button)
    .expect(Object.keys(button.attributes)).notContains('disabled')
})

test('Requesting access disables button', async t => {
  const button = page.requestAccess.find('button')

  await t
    .setNativeDialogHandler((type) => {
      if (type === 'confirm') return true
    })
    .click(button)
    .expect(button.withAttribute('disabled').exists).ok
})

// In IE11 the transition doesn't always complete when directly navigating to detail
// So X may not be visible/clickable
test.before( async t => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/directory`)
  })
  ('Opens and closes detail view', async t => {
    const row = Selector('tr').withText('basisregistratie')

    await t
      .click(row)
      .expect(getLocation()).contains(`${baseUrl}/directory/BRP/basisregistratie`)
      .expect(page.closeButton.exists).ok()
      .click(page.closeButton)
      .expect(getLocation()).contains(`${baseUrl}/directory`);
  })
