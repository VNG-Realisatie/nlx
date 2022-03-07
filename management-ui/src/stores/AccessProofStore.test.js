// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessProofModel from '../stores/models/AccessProofModel'
import AccessProofStore from './AccessProofStore'

test('updating from server', async () => {
  const accessProofStore = new AccessProofStore()

  expect(accessProofStore.accessProofs).toHaveLength(0)
  expect(accessProofStore.updateFromServer()).toBeNull()

  let accessProof = await accessProofStore.updateFromServer({
    id: 'abcd',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'organization-name',
    },
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: null,
  })

  // new model should be created
  expect(accessProofStore.accessProofs).toHaveLength(1)
  expect(accessProof).toBeInstanceOf(AccessProofModel)

  accessProof = await accessProofStore.updateFromServer({
    id: 'abcd',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'organization-name',
    },
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: '2020-10-02',
  })

  // existing model should be updated
  expect(accessProofStore.accessProofs).toHaveLength(1)
  expect(accessProof).toBeInstanceOf(AccessProofModel)
})

test('get model by id', () => {
  const accessProofStore = new AccessProofStore()

  expect(accessProofStore.getById('42')).toBeUndefined()

  accessProofStore.updateFromServer({
    id: '42',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'organization-name',
    },
    serviceName: 'service-name',
    createdAt: '2020-10-01',
    revokedAt: null,
  })

  expect(accessProofStore.getById('42')).toBeInstanceOf(AccessProofModel)
})
