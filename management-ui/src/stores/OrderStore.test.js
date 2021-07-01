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
  expect(auditLogStore.outgoingOrders).toEqual([])
})

test('fetch outgoing orders', async () => {
  const managementApiClient = new ManagementApi()
  //
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

  await expect(store.fetchOutgoing()).rejects.toThrowError('arbitrary error')
  expect(store.isLoading).toBe(false)

  store.fetchOutgoing()
  expect(store.isLoading).toBe(true)

  await waitFor(() => expect(store.isLoading).toBe(false))
  expect(store.outgoingOrders).toEqual([{ reference: 'reference' }])
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

test('fetch incoming orders', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListIncomingOrders = jest
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

  await expect(store.fetchIncoming()).rejects.toThrowError('arbitrary error')
  expect(store.isLoading).toBe(false)

  store.fetchIncoming()
  expect(store.isLoading).toBe(true)

  await waitFor(() => expect(store.isLoading).toBe(false))
  expect(store.incomingOrders).toEqual([{ reference: 'reference' }])
})

test('update incoming orders', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementSynchronizeOrders = jest
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

  await expect(store.updateIncoming()).rejects.toThrowError('arbitrary error')
  expect(store.isLoading).toBe(false)

  store.updateIncoming()
  expect(store.isLoading).toBe(true)

  await waitFor(() => expect(store.isLoading).toBe(false))
  expect(store.incomingOrders).toEqual([{ reference: 'reference' }])
})
