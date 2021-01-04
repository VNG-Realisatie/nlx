// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel from '../models/OutgoingAccessRequestModel'
import DirectoryServiceModel from '../models/DirectoryServiceModel'
import AccessProofModel from '../models/AccessProofModel'
import { DirectoryApi } from '../api'
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
  const rootStore = new RootStore({
    directoryRepository: {
      getAll: jest.fn().mockResolvedValue([
        {
          organizationName: 'Org A',
          serviceName: 'Service A',
        },
        {
          organizationName: 'Org B',
          serviceName: 'Service B',
        },
      ]),
    },
  })

  const directoryServicesStore = rootStore.directoryServicesStore
  await directoryServicesStore.fetchAll()

  expect(directoryServicesStore.directoryRepository.getAll).toHaveBeenCalled()

  await expect(directoryServicesStore.isInitiallyFetched).toBe(true)
  expect(directoryServicesStore.services).toHaveLength(2)
})

test('handle error while fetching all directory services', async () => {
  global.console.error = jest.fn()

  const rootStore = new RootStore({
    directoryRepository: {
      getAll: jest.fn().mockRejectedValue(new Error('arbitrary error')),
    },
  })

  const directoryServicesStore = rootStore.directoryServicesStore
  await directoryServicesStore.fetchAll()

  expect(directoryServicesStore.directoryRepository.getAll).toHaveBeenCalled()

  expect(directoryServicesStore.error).toEqual(new Error('arbitrary error'))
  expect(directoryServicesStore.services).toEqual([])
  expect(directoryServicesStore.isInitiallyFetched).toBe(true)
})

test('fetching a single service', async () => {
  const directoryApiService = new DirectoryApi()

  directoryApiService.directoryGetOrganizationService = jest
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
    directoryApiService,
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
  const { outgoingAccessRequestStore, directoryServicesStore } = new RootStore({
    accessRequestRepository: {
      createAccessRequest: jest.fn().mockResolvedValue({
        id: '42',
      }),
    },
  })

  jest.spyOn(outgoingAccessRequestStore, 'create')

  const directoryService = new DirectoryServiceModel({
    serviceData: {
      organizationName: 'organization',
      serviceName: 'service',
    },
  })

  const outgoingAccessRequest = await directoryServicesStore.requestAccess(
    directoryService,
  )

  expect(outgoingAccessRequestStore.create).toHaveBeenCalledWith({
    organizationName: 'organization',
    serviceName: 'service',
  })

  expect(outgoingAccessRequest).toEqual(
    new OutgoingAccessRequestModel({
      accessRequestData: {
        id: '42',
      },
    }),
  )
})
