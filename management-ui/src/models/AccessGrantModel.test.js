// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessGrantModel from './AccessGrantModel'

let accessGrantStore
let accessGrantData

beforeEach(() => {
  accessGrantStore = {}
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
  const accessProof = new AccessGrantModel({
    accessGrantStore,
    accessGrantData,
  })

  expect(accessProof.id).toBe(accessGrantData.id)
  expect(accessProof.organizationName).toBe(accessGrantData.organizationName)
  expect(accessProof.serviceName).toBe(accessGrantData.serviceName)
  expect(accessProof.publicKeyFingerprint).toBe(
    accessGrantData.publicKeyFingerprint,
  )
  expect(accessProof.createdAt).toEqual(new Date(accessGrantData.createdAt))
  expect(accessProof.revokedAt).toEqual(accessGrantData.revokedAt)
})

test('rejecting request handles as expected', async () => {
  const revokeAccessGrant = jest.fn().mockResolvedValue(null)

  accessGrantStore = {
    revokeAccessGrant,
    fetchForService: jest.fn(),
  }

  const accessGrant = new AccessGrantModel({
    accessGrantStore,
    accessGrantData,
  })

  accessGrant.revoke()

  expect(revokeAccessGrant).toHaveBeenCalled()
})
