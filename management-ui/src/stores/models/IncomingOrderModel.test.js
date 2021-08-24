// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import OrderStore from '../OrderStore'
import { ManagementApi } from '../../api'
import IncomingOrderModel from './IncomingOrderModel'

test('creating instance', () => {
  const order = new IncomingOrderModel({
    orderData: {
      reference: 'my-reference',
      delegator: 'Vergunningsoftware BV',
      description: 'description',
      services: [],
      revokedAt: '2020-10-03T12:00:00Z',
      validFrom: '2020-10-01T12:00:00Z',
      validUntil: '2020-10-02T12:00:00Z',
    },
  })

  expect(order.reference).toEqual('my-reference')
  expect(order.delegator).toBe('Vergunningsoftware BV')
  expect(order.description).toBe('description')
  expect(order.services).toEqual([])
  expect(order.revokedAt).toEqual(new Date('2020-10-03T12:00:00Z'))
  expect(order.validFrom).toEqual(new Date('2020-10-01T12:00:00Z'))
  expect(order.validUntil).toEqual(new Date('2020-10-02T12:00:00Z'))
})

test('revoke order', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementApi()

  const store = new OrderStore({
    rootStore: {},
    managementApiClient,
  })

  const order = new IncomingOrderModel({
    orderStore: store,
    orderData: {},
  })

  jest.spyOn(store, 'revokeIncoming').mockResolvedValue()

  await order.revoke()

  expect(store.revokeIncoming).toBeCalledWith(order)
})