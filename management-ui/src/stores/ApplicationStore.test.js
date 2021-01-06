// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import { ManagementApi } from '../api'
import ApplicationStore from './ApplicationStore'

test('initializing the store', () => {
  const accessGrantStore = new ApplicationStore({
    managementApiClient: new ManagementApi(),
  })

  expect(accessGrantStore.isOrganizationInwaySet).toBeNull()
})

test.concurrent.each([
  [true, true],
  [false, false],
  ['', false],
  ['inway-name', true],
])('updating isOrganizationInwaySet to %s', (a, expected) => {
  const accessGrantStore = new ApplicationStore({
    managementApiClient: new ManagementApi(),
  })

  accessGrantStore.update({
    isOrganizationInwaySet: a,
  })

  expect(accessGrantStore.isOrganizationInwaySet).toBe(expected)
})
