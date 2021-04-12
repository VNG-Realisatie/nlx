// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import OrderStore from './OrderStore'

test('create an order', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementCreateOrder = jest.fn().mockResolvedValue({
    id: 'orderid',
  })

  const store = new OrderStore({
    rootStore: {},
    managementApiClient,
  })

  expect(await store.create()).toEqual('orderid')
})
