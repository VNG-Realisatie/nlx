// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessProofModel from './AccessProofModel'
import AccessGrantModel from './AccessGrantModel'

test('should properly construct object', () => {
  const model = new AccessProofModel({
    accessProofData: {
      id: '42',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Organization One',
      },
      serviceName: 'Service',
      createdAt: '2020-10-01',
      revokedAt: null,
      publicKeyFingerprint: 'public-key-fingerprint',
    },
  })

  expect(model.id).toBe('42')
  expect(model.organization.name).toBe('Organization One')
  expect(model.organization.serialNumber).toBe('00000000000000000001')
  expect(model.serviceName).toBe('Service')
  expect(model.publicKeyFingerprint).toBe('public-key-fingerprint')
  expect(model.createdAt).toEqual(new Date('2020-10-01'))
  expect(model.revokedAt).toEqual(null)
})

test('organization name is empty', () => {
  const model = new AccessProofModel({
    accessProofData: {
      organization: {
        name: '',
        serialNumber: '00000000000000000001',
      },
    },
  })

  expect(model.organization.name).toBe('00000000000000000001')
  expect(model.organization.serialNumber).toBe('00000000000000000001')
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
