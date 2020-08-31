// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { RequestLogger, Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'
import { axeCheck, createReport } from 'axe-testcafe'

import { getBaseUrl, saveBrowserConsoleAndRequests } from '../../utils'
const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders:    false,
  logResponseBody:       true,
  stringifyResponseBody: true,
});

fixture`Not Found (404) page`.beforeEach(async (t) => {
  await t.navigateTo(`${baseUrl}/page-that-does-not-exist'`)
  await waitForReact()
})
  .afterEach(async (t) =>
    saveBrowserConsoleAndRequests(t, logger.requests)
  ).requestHooks(logger);

test('Automated accessibility testing', async (t) => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations))
})

test('Page not found message is present', async (t) => {
  await t.expect(Selector('h1').innerText).eql('Pagina niet gevonden')
})
