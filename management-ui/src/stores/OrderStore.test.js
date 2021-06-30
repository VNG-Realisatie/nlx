// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { waitFor } from '@testing-library/react'
import { ManagementApi } from '../api'
import OrderStore from './OrderStore'

test('initializing the store', () => {
  const auditLogStore = new OrderStore({
    managementApiClient: new ManagementApi(),
  })

  expect(auditLogStore.isLoading).toEqual(false)
  expect(auditLogStore.orders).toEqual([])
})

test('fetch all orders', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockRejectedValueOnce(new Error('arbitrary error'))
    .mockResolvedValue({
      orders: [
        {
          reference: 'reference',
        },
      ],
    })

  const store = new OrderStore({
    rootStore: {},
    managementApiClient,
  })

  await expect(store.fetchAll()).rejects.toThrowError('arbitrary error')
  expect(store.isLoading).toBe(false)

  store.fetchAll()
  expect(store.isLoading).toBe(true)

  await waitFor(() => expect(store.isLoading).toBe(false))
  expect(store.orders).toEqual([{ reference: 'reference' }])
})

test('create an order', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementCreateOutgoingOrder = jest
    .fn()
    .mockResolvedValue({
      id: 'orderid',
    })

  const store = new OrderStore({
    rootStore: {},
    managementApiClient,
  })

  expect(await store.create()).toEqual('orderid')
})
