// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { configure } from 'mobx'
import OrderStore from '../OrderStore'
import { ManagementApi } from '../../api'
import OutgoingOrderModel from './OutgoingOrderModel'
import AccessProofModel from './AccessProofModel'

test('creating instance', () => {
  const accessProof = new AccessProofModel({
    accessProofData: {
      id: '42',
    },
  })

  const order = new OutgoingOrderModel({
    orderData: {
      reference: 'my-reference',
      delegatee: {
        serialNumber: '00000000000000000001',
        name: 'Organization One',
      },
      description: 'description',
      revokedAt: '2020-10-03T12:00:00Z',
      validFrom: '2020-10-01T12:00:00Z',
      validUntil: '2020-10-02T12:00:00Z',
    },
    accessProofs: [accessProof],
  })

  expect(order.reference).toEqual('my-reference')
  expect(order.delegatee.serialNumber).toBe('00000000000000000001')
  expect(order.delegatee.name).toBe('Organization One')
  expect(order.description).toBe('description')
  expect(order.revokedAt).toEqual(new Date('2020-10-03T12:00:00Z'))
  expect(order.validFrom).toEqual(new Date('2020-10-01T12:00:00Z'))
  expect(order.validUntil).toEqual(new Date('2020-10-02T12:00:00Z'))
  expect(order.accessProofs).toEqual([accessProof])
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
