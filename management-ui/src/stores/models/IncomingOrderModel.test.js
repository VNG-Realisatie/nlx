// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import IncomingOrderModel from './IncomingOrderModel'
import AccessGrantModel from './AccessGrantModel'

test('creating instance', () => {
  const order = new IncomingOrderModel({
    orderData: {
      reference: 'my-reference',
      delegator: {
        serialNumber: '00000000000000000001',
        name: 'Organization One',
      },
      description: 'description',
      services: [],
      revokedAt: '2020-10-03T12:00:00Z',
      validFrom: '2020-10-01T12:00:00Z',
      validUntil: '2020-10-02T12:00:00Z',
    },
  })

  expect(order.reference).toEqual('my-reference')
  expect(order.delegator.serialNumber).toBe('00000000000000000001')
  expect(order.delegator.name).toBe('Organization One')
  expect(order.description).toBe('description')
  expect(order.services).toEqual([])
  expect(order.revokedAt).toEqual(new Date('2020-10-03T12:00:00Z'))
  expect(order.validFrom).toEqual(new Date('2020-10-01T12:00:00Z'))
  expect(order.validUntil).toEqual(new Date('2020-10-02T12:00:00Z'))
})

test('organization name is empty', () => {
  const model = new AccessGrantModel({
    accessGrantStore: {},
    accessGrantData: {
      organization: {
        name: '',
        serialNumber: '00000000000000000001',
      },
    },
  })

  expect(model.organization.name).toBe('00000000000000000001')
  expect(model.organization.serialNumber).toBe('00000000000000000001')
})
