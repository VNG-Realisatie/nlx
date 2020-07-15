// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import {
  DIRECTORY_ORGANIZATION_NAME,
  DIRECTORY_SERVICE_NAME,
  DIRECTORY_STATUS,
  DIRECTORY_API_SPECIFICATION_TYPE,
} from '../../environment'
import { getBaseUrl, getLocation } from '../../utils'
import { adminUser } from '../roles'
import page from './page-models/directory-detail'

const baseUrl = getBaseUrl()

fixture `DirectoryDetails page`
  .beforeEach(async (t) => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/directory/${DIRECTORY_ORGANIZATION_NAME}/${DIRECTORY_SERVICE_NAME}`)
    await waitForReact()
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Directory service details are visible', async t => {
  const detailHeaderContent = page.detailHeader.textContent
  
  await t.expect(detailHeaderContent).contains(DIRECTORY_SERVICE_NAME)
  await t.expect(detailHeaderContent).contains(DIRECTORY_ORGANIZATION_NAME)
  await t.expect(detailHeaderContent).contains(DIRECTORY_API_SPECIFICATION_TYPE)
  await t.expect(detailHeaderContent).contains(DIRECTORY_STATUS)
})

// test('Request access button exists', async t => {
//   const button = page.requestAccess.find('button')
//   await t.expect(button.exists).ok()
// })

// test('Request access can be cancelled in dialog', async t => {
//   const button = page.requestAccess.find('button')

//   await t.setNativeDialogHandler((type) => {
//     if (type === 'confirm') return false
//   })
    
//   await t.click(button)
//   await t.expect(Object.keys(button.attributes)).notContains('disabled')
// })

// test('After requesting access, I see no button and confirmation of my request', async t => {
//   const button = page.requestAccess.find('button')

//   await t.setNativeDialogHandler((type) => {
//     if (type === 'confirm') return true
//   })
  
//   await t.click(button)
//   await t.expect(button.exists).notOk()
//   await t.expect(Selector('[data-testid="access-message"]').textContent).eql('Toegang aangevraagd')
// })

// In IE11 the transition doesn't always complete when directly navigating to detail
// So X may not be visible/clickable
test.before( async t => {
    await t
      .useRole(adminUser)
      .navigateTo(`${baseUrl}/directory`)
  })
  ('Opens and closes detail view', async t => {
    const row = Selector('tr').withText(DIRECTORY_SERVICE_NAME)

    await t.click(row)
    await t.expect(getLocation()).contains(`${baseUrl}/directory/${DIRECTORY_ORGANIZATION_NAME}/${DIRECTORY_SERVICE_NAME}`)
    await t.expect(page.closeButton.exists).ok()
    
    await t.click(page.closeButton)
    await t.expect(getLocation()).contains(`${baseUrl}/directory`)
  })
