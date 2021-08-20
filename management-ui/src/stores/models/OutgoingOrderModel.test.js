// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { configure } from 'mobx'
import OrderStore from '../OrderStore'
import { ManagementApi } from '../../api'
import OutgoingOrderModel from './OutgoingOrderModel'

test('creating instance', () => {
  const order = new OutgoingOrderModel({
    orderData: {
      reference: 'my-reference',
      delegatee: 'Vergunningsoftware BV',
      description: 'description',
      services: [],
      revokedAt: '2020-10-03T12:00:00Z',
      validFrom: '2020-10-01T12:00:00Z',
      validUntil: '2020-10-02T12:00:00Z',
    },
  })

  expect(order.reference).toEqual('my-reference')
  expect(order.delegatee).toBe('Vergunningsoftware BV')
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

  const order = new OutgoingOrderModel({
    orderStore: store,
    orderData: {},
  })

  jest.spyOn(store, 'revokeOutgoing').mockResolvedValue()

  await order.revoke()

  expect(store.revokeOutgoing).toBeCalledWith(order)
})
