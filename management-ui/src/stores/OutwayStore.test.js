// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { ManagementServiceApi } from '../api'
import OutwayModel from './models/OutwayModel'
import OutwayStore from './OutwayStore'

test('initializing the store', () => {
  const store = new OutwayStore({
    managementApiClient: new ManagementServiceApi(),
  })

  expect(store.outways).toEqual([])
})

test('fetching all outways', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'my-outway',
          ipAddress: '127.0.0.1',
          publicKeyPem: 'public-key-pem',
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
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceGetOutway = jest.fn().mockResolvedValue({
    name: 'My Outway',
    ipAddress: '127.0.0.1',
    publicKeyPem: 'public-key-pem',
    publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
    version: 'v0.0.42',
  })

  const outwayStore = new OutwayStore({
    rootStore: {},
    managementApiClient,
  })

  expect(outwayStore.getByName('non-existing-outway-name')).toBeUndefined()

  const outway = await outwayStore.fetch({ name: 'My Outway' })

  expect(managementApiClient.managementServiceGetOutway).toHaveBeenCalledWith({
    name: 'My Outway',
  })
  expect(outway).toBeInstanceOf(OutwayModel)
  expect(outway.name).toEqual('My Outway')

  expect(outwayStore.getByName(outway.name).name).toEqual('My Outway')
})

test('retrieving the public key fingerprints', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'my-outway-1',
          ipAddress: '127.0.0.1',
          publicKeyPem: 'public-key-pem-a',
          publicKeyFingerprint: 'public-key-fingerprint-a',
          version: 'v0.0.42',
        },
        {
          name: 'my-outway-2',
          ipAddress: '127.0.0.1',
          publicKeyPem: 'public-key-pem-b',
          publicKeyFingerprint: 'public-key-fingerprint-b',
          version: 'v0.0.42',
        },
        {
          name: 'my-outway-3',
          ipAddress: '127.0.0.1',
          publicKeyPem: 'public-key-pem-c',
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
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'my-outway-1',
          ipAddress: '127.0.0.1',
          publicKeyPem: 'public-key-pem',
          publicKeyFingerprint: 'public-key-fingerprint-a',
          version: 'v0.0.42',
        },
        {
          name: 'my-outway-2',
          ipAddress: '127.0.0.1',
          publicKeyPem: 'public-key-pem',
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

test('removing an outway', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValueOnce({
      outways: [
        { name: 'Outway A' },
        { name: 'Outway B' },
        { name: 'Outway C' },
      ],
    })
    .mockResolvedValue({
      outways: [{ name: 'Outway A' }, { name: 'Outway C' }],
    })

  managementApiClient.managementServiceDeleteOutway = jest
    .fn()
    .mockResolvedValue({})

  const outwayStore = new OutwayStore({
    rootStore: {},
    managementApiClient,
  })

  await outwayStore.fetchAll()
  jest.spyOn(outwayStore, 'fetchAll')

  expect(outwayStore.getByName('Outway A')).toBeDefined()
  expect(outwayStore.getByName('Outway B')).toBeDefined()
  expect(outwayStore.getByName('Outway C')).toBeDefined()

  await outwayStore.removeOutway('Outway B')

  expect(
    managementApiClient.managementServiceDeleteOutway,
  ).toHaveBeenCalledWith({
    name: 'Outway B',
  })

  expect(outwayStore.fetchAll).toHaveBeenCalledTimes(1)
  expect(outwayStore.getByName('Outway A')).toBeDefined()
  expect(outwayStore.getByName('Outway B')).toBeUndefined()
  expect(outwayStore.getByName('Outway C')).toBeDefined()
})
