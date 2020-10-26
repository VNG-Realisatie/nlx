// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessProofModel from '../models/AccessProofModel'
import AccessProofStore from './AccessProofStore'

test('updating from server', async () => {
  const accessProofStore = new AccessProofStore()

  expect(accessProofStore.accessProofs.size).toEqual(0)
  expect(accessProofStore.updateFromServer()).toBeNull()

  let accessProof = await accessProofStore.updateFromServer({
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: null,
  })

  // new model should be created
  expect(accessProofStore.accessProofs.size).toEqual(1)
  expect(accessProof).toBeInstanceOf(AccessProofModel)

  accessProof = await accessProofStore.updateFromServer({
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: '2020-10-02',
  })

  // existing model should be updated
  expect(accessProofStore.accessProofs.size).toEqual(1)
  expect(accessProof).toBeInstanceOf(AccessProofModel)
})
