// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import OutwayModel from './models/OutwayModel'
import OutwayStore from './OutwayStore'

test('initializing the store', () => {
  const store = new OutwayStore({
    managementApiClient: new ManagementApi(),
  })

  expect(store.outways).toEqual([])
})

test('fetching, getting and updating from server', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'my-outway',
        version: 'v0.0.42',
      },
    ],
  })

  const store = new OutwayStore({
    managementApiClient,
  })

  await store.fetchAll()
  expect(store.outways).toHaveLength(1)
  const initialAuditLog = store.outways[0]
  expect(initialAuditLog).toBeInstanceOf(OutwayModel)
})
