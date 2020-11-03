// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessGrantModel from './AccessGrantModel'

let accessGrantData

beforeEach(() => {
  accessGrantData = {
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    publicKeyFingerprint: 'f1ng3r',
    createdAt: '2020-10-01',
    revokedAt: null,
  }
})

test('should properly construct object', () => {
  const accessProof = new AccessGrantModel({ accessGrantData })

  expect(accessProof.id).toBe(accessGrantData.id)
  expect(accessProof.organizationName).toBe(accessGrantData.organizationName)
  expect(accessProof.serviceName).toBe(accessGrantData.serviceName)
  expect(accessProof.publicKeyFingerprint).toBe(
    accessGrantData.publicKeyFingerprint,
  )
  expect(accessProof.createdAt).toEqual(new Date(accessGrantData.createdAt))
  expect(accessProof.revokedAt).toEqual(accessGrantData.revokedAt)
})
