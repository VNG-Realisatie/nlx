// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessProofModel from './AccessProofModel'

let accessProofData

beforeEach(() => {
  accessProofData = {
    id: 'abcd',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'Organization',
    },
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: null,
    publicKeyFingerprint: 'public-key-fingerprint',
  }
})

test('should properly construct object', () => {
  const accessProof = new AccessProofModel({ accessProofData })

  expect(accessProof.id).toBe(accessProofData.id)
  expect(accessProof.organization.serialNumber).toBe(
    accessProofData.organization.serialNumber,
  )
  expect(accessProof.organization.name).toBe(accessProofData.organization.name)
  expect(accessProof.serviceName).toBe(accessProofData.serviceName)
  expect(accessProof.createdAt).toEqual(new Date(accessProofData.createdAt))
  expect(accessProof.revokedAt).toEqual(accessProofData.revokedAt)
  expect(accessProof.publicKeyFingerprint).toEqual(
    accessProofData.publicKeyFingerprint,
  )
})
