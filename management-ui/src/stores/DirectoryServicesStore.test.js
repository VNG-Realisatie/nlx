// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import OutgoingAccessRequestModel from '../stores/models/OutgoingAccessRequestModel'
import DirectoryServiceModel from '../stores/models/DirectoryServiceModel'
import AccessProofModel from '../stores/models/AccessProofModel'
import { DirectoryApi, ManagementApi } from '../api'
import DirectoryServicesStore from './DirectoryServicesStore'
import { RootStore } from './index'

test('initializing the store', async () => {
  const directoryServicesStore = new DirectoryServicesStore({
    rootStore: {},
  })

  expect(directoryServicesStore.services).toEqual([])
  expect(directoryServicesStore.isInitiallyFetched).toBe(false)
})

test('fetching all directory services', async () => {
  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryListServices = jest.fn().mockResolvedValue({
    services: [
      {
        organizationName: 'Org A',
        serviceName: 'Service A',
      },
      {
        organizationName: 'Org B',
        serviceName: 'Service B',
      },
    ],
  })

  const rootStore = new RootStore({
    directoryApiClient,
  })

  const directoryServicesStore = rootStore.directoryServicesStore
  await directoryServicesStore.fetchAll()

  await expect(directoryServicesStore.isInitiallyFetched).toBe(true)
  expect(directoryServicesStore.services).toHaveLength(2)
})

test('handle error while fetching all directory services', async () => {
  global.console.error = jest.fn()

  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryListServices = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const rootStore = new RootStore({
    directoryApiClient,
  })

  const directoryServicesStore = rootStore.directoryServicesStore
  await directoryServicesStore.fetchAll()

  expect(directoryServicesStore.error).toEqual(new Error('arbitrary error'))
  expect(directoryServicesStore.services).toEqual([])
  expect(directoryServicesStore.isInitiallyFetched).toBe(true)
})

test('fetching a single service', async () => {
  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryGetOrganizationService = jest
    .fn()
    .mockResolvedValueOnce({
      organizationName: 'Org A',
      serviceName: 'Service A',
    })
    .mockReturnValue({
      organizationName: 'Org A',
      serviceName: 'Service A',
      latestAccessRequest: { id: 'abc', state: 'CREATED' },
      latestAccessProof: { id: 'abc' },
    })

  const rootStore = new RootStore({
    directoryApiClient,
  })

  const directoryServicesStore = rootStore.directoryServicesStore
  expect(directoryServicesStore.services).toHaveLength(0)

  let service = await directoryServicesStore.fetch({
    organizationName: 'Org A',
    serviceName: 'Service A',
  })

  expect(directoryServicesStore.services).toHaveLength(1)
  expect(service).toBeInstanceOf(DirectoryServiceModel)
  expect(service.latestAccessRequest).toBeNull()
  expect(service.latestAccessProof).toBeNull()

  service = await directoryServicesStore.fetch({
    organizationName: 'Org A',
    serviceName: 'Service A',
  })

  expect(service).toBeInstanceOf(DirectoryServiceModel)
  expect(service.latestAccessRequest).toBeInstanceOf(OutgoingAccessRequestModel)
  expect(service.latestAccessProof).toBeInstanceOf(AccessProofModel)
})

test('requesting access to a service in the directory', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({
    managementApiClient,
  })

  jest.spyOn(rootStore.outgoingAccessRequestStore, 'create').mockResolvedValue()

  const directoryService = new DirectoryServiceModel({
    serviceData: {
      organizationName: 'organization',
      serviceName: 'service',
    },
  })

  await rootStore.directoryServicesStore.requestAccess(directoryService)

  expect(rootStore.outgoingAccessRequestStore.create).toHaveBeenCalledWith({
    organizationName: 'organization',
    serviceName: 'service',
  })
})
