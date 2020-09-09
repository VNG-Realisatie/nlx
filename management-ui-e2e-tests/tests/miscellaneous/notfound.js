// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { RequestLogger, Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'

import {
  getBaseUrl,
  doAccessibilityTest,
  saveBrowserConsoleAndRequests,
} from '../../utils'
const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders: false,
  logResponseBody: true,
  stringifyResponseBody: true,
})

fixture`Not Found (404) page`
  .beforeEach(async (t) => {
    await t.navigateTo(`${baseUrl}/page-that-does-not-exist'`)
    await waitForReact()
  })
  .afterEach(async (t) => saveBrowserConsoleAndRequests(t, logger.requests))
  .requestHooks(logger)

test('404 page exists and is accessible', async (t) => {
  await doAccessibilityTest(t)
  await t.expect(Selector('h1').innerText).eql('Pagina niet gevonden')
})
