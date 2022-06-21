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

  const model = new OutgoingOrderModel({
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

  expect(model.reference).toEqual('my-reference')
  expect(model.delegatee.serialNumber).toBe('00000000000000000001')
  expect(model.delegatee.name).toBe('Organization One')
  expect(model.description).toBe('description')
  expect(model.revokedAt).toEqual(new Date('2020-10-03T12:00:00Z'))
  expect(model.validFrom).toEqual(new Date('2020-10-01T12:00:00Z'))
  expect(model.validUntil).toEqual(new Date('2020-10-02T12:00:00Z'))
  expect(model.accessProofs).toEqual([accessProof])
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

test('organization name is empty', () => {
  const model = new OutgoingOrderModel({
    orderStore: {},
    orderData: {
      delegatee: {
        name: '',
        serialNumber: '00000000000000000001',
      },
    },
  })

  expect(model.delegatee.name).toBe('00000000000000000001')
  expect(model.delegatee.serialNumber).toBe('00000000000000000001')
})
