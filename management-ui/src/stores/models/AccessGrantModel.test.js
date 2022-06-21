// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessGrantModel from './AccessGrantModel'

test('should properly construct object', () => {
  const model = new AccessGrantModel({
    accessGrantStore: {},
    accessGrantData: {
      id: '42',
      organization: {
        name: 'Organization One',
        serialNumber: '00000000000000000001',
      },
      serviceName: 'Service',
      publicKeyFingerprint: 'f1ng3r',
      createdAt: '2020-10-01',
      revokedAt: null,
    },
  })

  expect(model.id).toBe('42')
  expect(model.organization.name).toBe('Organization One')
  expect(model.organization.serialNumber).toBe('00000000000000000001')
  expect(model.serviceName).toBe('Service')
  expect(model.publicKeyFingerprint).toBe('f1ng3r')
  expect(model.createdAt).toEqual(new Date('2020-10-01'))
  expect(model.revokedAt).toEqual(null)
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

test('rejecting request handles as expected', async () => {
  const revokeAccessGrant = jest.fn().mockResolvedValue(null)

  const accessGrant = new AccessGrantModel({
    accessGrantStore: {
      revokeAccessGrant,
      fetchForService: jest.fn(),
    },
    accessGrantData: {},
  })

  accessGrant.revoke()

  expect(revokeAccessGrant).toHaveBeenCalled()
})
