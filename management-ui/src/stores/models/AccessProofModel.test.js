// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessProofModel from './AccessProofModel'

let accessProofData

beforeEach(() => {
  accessProofData = {
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: null,
  }
})

test('should properly construct object', () => {
  const accessProof = new AccessProofModel({ accessProofData })

  expect(accessProof.id).toBe(accessProofData.id)
  expect(accessProof.organizationName).toBe(accessProofData.organizationName)
  expect(accessProof.serviceName).toBe(accessProofData.serviceName)
  expect(accessProof.createdAt).toEqual(new Date(accessProofData.createdAt))
  expect(accessProof.revokedAt).toEqual(accessProofData.revokedAt)
})
