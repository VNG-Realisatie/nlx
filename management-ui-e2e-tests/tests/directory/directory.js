// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { RequestLogger } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'

import {
  DIRECTORY_ORGANIZATION_NAME,
  DIRECTORY_SERVICE_NAME,
} from '../../environment'
import {
  getBaseUrl,
  doAccessibilityTest,
  saveBrowserConsoleAndRequests,
} from '../../utils'
import { adminUser } from '../roles'

const baseUrl = getBaseUrl()

const logger = RequestLogger(/api/, {
  logResponseHeaders: false,
  logResponseBody: true,
  stringifyResponseBody: true,
})

fixture`Directory`
  .beforeEach(async (t) => {
    await t.useRole(adminUser).navigateTo(`${baseUrl}/directory`)
    await waitForReact()
  })
  .afterEach(async (t) => saveBrowserConsoleAndRequests(t, logger.requests))
  .requestHooks(logger)

test('Directory list and detail page pass accessibility test', async (t) => {
  await doAccessibilityTest(t)

  await t.navigateTo(
    `${baseUrl}/directory/${DIRECTORY_ORGANIZATION_NAME}/${DIRECTORY_SERVICE_NAME}`,
  )

  // disable 'tabindex' because the 'focus-lock' dependency
  // creates an element with tabindex="1"
  // https://github.com/theKashey/react-focus-lock/blob/2b6ae70f0b15046ee3ac3227c53bb7c21f551ff4/src/Lock.js#L127
  const axeOptions = { rules: { tabindex: { enabled: false } } }
  await doAccessibilityTest(t, undefined, axeOptions)
})
