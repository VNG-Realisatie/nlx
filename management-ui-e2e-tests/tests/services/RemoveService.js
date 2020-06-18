// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { getBaseUrl } from '../../utils'
import { adminUser } from '../roles'
import { createService } from './actions'
import page from './page-models/service-detail'

const baseUrl = getBaseUrl()

fixture`ServiceDetails remove`.beforeEach(async (t) => {
  await t.useRole(adminUser)
  const serviceName = await createService()
  await t.navigateTo(`${baseUrl}/services/${serviceName}`)
})

test('Removing a service', async (t) => {
  await t
    .setNativeDialogHandler((type, text, url) => {
      if (text !== 'Wil je de service verwijderen?') {
        throw `Unexpected dialog text: ${text}`
      }
      return true
    })
  await t.click(page.removeButton).takeScreenshot()
  
  await t.expect(page.alert.visible).ok()
  await t.expect(page.alert.innerText).contains('De service is verwijderd.')
})
