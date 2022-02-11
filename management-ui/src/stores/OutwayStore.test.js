// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import OutwayModel from './models/OutwayModel'
import OutwayStore from './OutwayStore'

test('initializing the store', () => {
  const store = new OutwayStore({
    managementApiClient: new ManagementApi(),
  })

  expect(store.outways).toEqual([])
})

test('fetching all outways', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'my-outway',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
        version: 'v0.0.42',
      },
    ],
  })

  const store = new OutwayStore({
    managementApiClient,
  })

  await store.fetchAll()
  expect(store.outways).toHaveLength(1)
  const initialOutway = store.outways[0]
  expect(initialOutway).toBeInstanceOf(OutwayModel)
})

test('fetching a single outway', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetOutway = jest.fn().mockResolvedValue({
    name: 'My Outway',
    ipAddress: '127.0.0.1',
    publicKeyPEM: 'public-key-pem',
    publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
    version: 'v0.0.42',
  })

  const outwayStore = new OutwayStore({
    rootStore: {},
    managementApiClient,
  })

  expect(outwayStore.getByName('non-existing-outway-name')).toBeUndefined()

  const outway = await outwayStore.fetch({ name: 'My Outway' })

  expect(managementApiClient.managementGetOutway).toHaveBeenCalledWith({
    name: 'My Outway',
  })
  expect(outway).toBeInstanceOf(OutwayModel)
  expect(outway.name).toEqual('My Outway')

  expect(outwayStore.getByName(outway.name).name).toEqual('My Outway')
})

test('retrieving the public key fingerprints', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'my-outway-1',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        publicKeyFingerprint: 'public-key-fingerprint-a',
        version: 'v0.0.42',
      },
      {
        name: 'my-outway-2',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        publicKeyFingerprint: 'public-key-fingerprint-b',
        version: 'v0.0.42',
      },
      {
        name: 'my-outway-3',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        publicKeyFingerprint: 'public-key-fingerprint-a',
        version: 'v0.0.42',
      },
    ],
  })

  const store = new OutwayStore({
    managementApiClient,
  })

  await store.fetchAll()
  expect(store.publicKeyFingerprints).toEqual([
    'public-key-fingerprint-a',
    'public-key-fingerprint-b',
  ])
})

test('get outways by public key fingerprint', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'my-outway-1',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        publicKeyFingerprint: 'public-key-fingerprint-a',
        version: 'v0.0.42',
      },
      {
        name: 'my-outway-2',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        publicKeyFingerprint: 'public-key-fingerprint-b',
        version: 'v0.0.42',
      },
    ],
  })

  const store = new OutwayStore({
    managementApiClient,
  })

  await store.fetchAll()
  expect(
    store.getByPublicKeyFingerprint('public-key-fingerprint-a')[0].name,
  ).toEqual('my-outway-1')
})
