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
