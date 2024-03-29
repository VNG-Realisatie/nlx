// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { waitFor } from '@testing-library/react'
import { ManagementServiceApi } from '../api'
import OrderStore from './OrderStore'
import OutgoingOrderModel from './models/OutgoingOrderModel'
import IncomingOrderModel from './models/IncomingOrderModel'
import { RootStore } from './index'

test('initializing the store', () => {
  const auditLogStore = new OrderStore({
    managementApiClient: new ManagementServiceApi(),
  })

  expect(auditLogStore.isLoading).toEqual(false)
  expect(auditLogStore.outgoingOrders).toEqual([])
})

test('fetch outgoing orders', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockRejectedValueOnce(new Error('arbitrary error'))
    .mockResolvedValue({
      orders: [
        {
          delegatee: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
          reference: 'reference',
        },
        {
          delegatee: {
            serialNumber: '00000000000000000002',
            name: 'Organization Two',
          },
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

  const firstOrder = store.getOutgoing('00000000000000000001', 'reference')
  expect(firstOrder).toBeInstanceOf(OutgoingOrderModel)
})

test('revoke outgoing order', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegatee: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
          reference: 'reference',
        },
      ],
    })

  managementApiClient.managementServiceRevokeOutgoingOrder = jest
    .fn()
    .mockResolvedValue()

  const store = new OrderStore({
    rootStore: {},
    managementApiClient,
  })

  await store.fetchOutgoing()

  const firstOrder = store.getOutgoing('00000000000000000001', 'reference')

  jest.spyOn(store, 'revokeOutgoing')

  await store.revokeOutgoing(firstOrder)

  expect(
    managementApiClient.managementServiceRevokeOutgoingOrder,
  ).toBeCalledWith({
    delegatee: '00000000000000000001',
    reference: 'reference',
  })

  expect(firstOrder.revokedAt).not.toBeNull()
})

test('creating an outgoing order', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceCreateOutgoingOrder = jest
    .fn()
    .mockResolvedValue()

  const rootStore = new RootStore({
    managementApiClient,
  })
  const orderStore = rootStore.orderStore

  expect(
    async () =>
      await orderStore.create({
        delegatee: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        reference: 'my-reference',
      }),
  ).not.toThrow()
})

test('fetch incoming orders', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListIncomingOrders = jest
    .fn()
    .mockRejectedValueOnce(new Error('arbitrary error'))
    .mockResolvedValue({
      orders: [
        {
          delegator: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
          reference: 'reference',
        },
        {
          delegator: {
            serialNumber: '00000000000000000002',
            name: 'Organization Two',
          },
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

  const firstOrder = store.getIncoming('00000000000000000001', 'reference')
  expect(firstOrder).toBeInstanceOf(IncomingOrderModel)
})
