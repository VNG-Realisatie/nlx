// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { waitFor } from '@testing-library/react'
import { ManagementApi } from '../api'
import OrderStore from './OrderStore'
import OutgoingOrderModel from './models/OutgoingOrderModel'
import IncomingOrderModel from './models/IncomingOrderModel'
import { RootStore } from './index'

test('initializing the store', () => {
  const auditLogStore = new OrderStore({
    managementApiClient: new ManagementApi(),
  })

  expect(auditLogStore.isLoading).toEqual(false)
  expect(auditLogStore.outgoingOrders).toEqual([])
})

test('fetch outgoing orders', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockRejectedValueOnce(new Error('arbitrary error'))
    .mockResolvedValue({
      orders: [
        {
          delegatee: 'delegatee 1',
          reference: 'reference',
        },
        {
          delegatee: 'delegatee 2',
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
  expect(store.outgoingOrders).toHaveLength(2)

  const firstOrder = store.getOutgoing('delegatee 1', 'reference')
  expect(firstOrder).toBeInstanceOf(OutgoingOrderModel)
})

test('revoke outgoing order', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegatee: 'delegatee 1',
          reference: 'reference',
        },
      ],
    })

  managementApiClient.managementRevokeOutgoingOrder = jest
    .fn()
    .mockResolvedValue()

  const store = new OrderStore({
    rootStore: {},
    managementApiClient,
  })

  await store.fetchOutgoing()

  const firstOrder = store.getOutgoing('delegatee 1', 'reference')

  jest.spyOn(store, 'revokeOutgoing')

  await store.revokeOutgoing(firstOrder)

  expect(managementApiClient.managementRevokeOutgoingOrder).toBeCalledWith({
    delegatee: 'delegatee 1',
    reference: 'reference',
  })

  expect(firstOrder.revokedAt).not.toBeNull()
})

test('creating an outgoing order', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementCreateOutgoingOrder = jest
    .fn()
    .mockResolvedValue()

  const rootStore = new RootStore({
    managementApiClient,
  })
  const orderStore = rootStore.orderStore

  const order = await orderStore.create({
    delegatee: 'delegatee',
    reference: 'my-reference',
  })

  expect(order).toBeInstanceOf(OutgoingOrderModel)
  expect(orderStore.getOutgoing('delegatee', 'my-reference')).toBe(order)
})

test('fetch incoming orders', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListIncomingOrders = jest
    .fn()
    .mockRejectedValueOnce(new Error('arbitrary error'))
    .mockResolvedValue({
      orders: [
        {
          delegator: 'delegator 1',
          reference: 'reference',
        },
        {
          delegator: 'delegator 2',
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
  expect(store.incomingOrders).toHaveLength(2)

  const firstOrder = store.getIncoming('delegator 1', 'reference')
  expect(firstOrder).toBeInstanceOf(IncomingOrderModel)
})

test('revoke incoming order', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListIncomingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegator: 'delegator 1',
          reference: 'reference',
        },
      ],
    })

  managementApiClient.managementRevokeIncomingOrder = jest
    .fn()
    .mockResolvedValue()

  const store = new OrderStore({
    rootStore: {},
    managementApiClient,
  })

  await store.fetchIncoming()

  const firstOrder = store.getIncoming('delegator 1', 'reference')

  jest.spyOn(store, 'revokeIncoming')

  await store.revokeIncoming(firstOrder)

  expect(managementApiClient.managementRevokeIncomingOrder).toBeCalledWith({
    delegator: 'delegator 1',
    reference: 'reference',
  })

  expect(firstOrder.revokedAt).not.toBeNull()
})
