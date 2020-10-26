// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { act } from '@testing-library/react'
import deferredPromise from '../test-utils/deferred-promise'
import OutgoingAccessRequestModel from '../models/OutgoingAccessRequestModel'
import DirectoryServiceModel from '../models/DirectoryServiceModel'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'
import AccessProofStore from './AccessProofStore'
import DirectoryServicesStore from './DirectoryServicesStore'

test('fetching all directory services', async () => {
  const mockGetAll = deferredPromise()
  const directoryRepository = {
    getAll: jest.fn().mockReturnValue(mockGetAll),
  }

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore: {
      outgoingAccessRequestsStore: new OutgoingAccessRequestStore({
        accessRequestRepository: {},
      }),
      accessProofStore: new AccessProofStore(),
    },
    directoryRepository,
  })

  expect(directoryServicesStore.services).toEqual([])
  expect(directoryServicesStore.isInitiallyFetched).toBe(false)

  directoryServicesStore.fetchAll()

  expect(directoryServicesStore.isInitiallyFetched).toBe(false)
  expect(directoryRepository.getAll).toHaveBeenCalled()

  await act(async () => {
    await mockGetAll.resolve([
      {
        organizationName: 'Org A',
        serviceName: 'Service A',
        latestAccessRequest: { id: 'abc' },
      },
      { organizationName: 'Org B', serviceName: 'Service B' },
    ])
  })

  await expect(directoryServicesStore.isInitiallyFetched).toBe(true)
  expect(directoryServicesStore.services).toHaveLength(2)
})

test('selecting and updating a single service', async () => {
  const directoryRepository = {
    getAll: jest
      .fn()
      .mockReturnValue([
        { organizationName: 'Org A', serviceName: 'Service A' },
      ]),

    getByName: jest.fn().mockReturnValue({
      organizationName: 'Org A',
      serviceName: 'Service A',
      latestAccessRequest: { id: 'abc', state: 'CREATED' },
    }),
  }

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore: {
      outgoingAccessRequestsStore: new OutgoingAccessRequestStore({
        accessRequestRepository: {},
      }),
      accessProofStore: new AccessProofStore(),
    },
    directoryRepository,
  })

  await directoryServicesStore.fetchAll()

  expect(directoryServicesStore.services).toHaveLength(1)

  const service = directoryServicesStore.getService({
    organizationName: 'Org A',
    serviceName: 'Service A',
  })

  expect(service.latestAccessRequest).toBeNull()

  await directoryServicesStore.fetch(service)

  expect(service.latestAccessRequest).not.toBeNull()
})

test('handle error while fetching directory services', async () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})

  const mockGetAll = deferredPromise()
  const directoryRepository = {
    getAll: jest.fn(() => mockGetAll),
  }

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore: {},
    directoryRepository,
  })

  expect(directoryServicesStore.services).toEqual([])

  directoryServicesStore.fetchAll()

  expect(directoryServicesStore.isInitiallyFetched).toBe(false)
  expect(directoryRepository.getAll).toHaveBeenCalled()

  await mockGetAll.reject('some error')

  expect(directoryServicesStore.error).toEqual('some error')
  expect(directoryServicesStore.services).toEqual([])
  expect(directoryServicesStore.isInitiallyFetched).toBe(true)

  errorSpy.mockRestore()
})

test('requesting access to a service in the directory', async () => {
  const accessRequestRepository = {
    createAccessRequest: jest.fn().mockResolvedValue({
      id: '42',
    }),
  }

  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
    accessRequestRepository,
  })

  const outgoingAccessRequestStoreCreateSpy = jest.spyOn(
    outgoingAccessRequestStore,
    'create',
  )

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore: {
      outgoingAccessRequestsStore: outgoingAccessRequestStore,
    },
    directoryRepository: {},
  })

  const directoryService = new DirectoryServiceModel({
    service: {
      organizationName: 'organization',
      serviceName: 'service',
    },
  })
  const outgoingAccessRequest = await directoryServicesStore.requestAccess(
    directoryService,
  )

  expect(outgoingAccessRequestStoreCreateSpy).toHaveBeenCalledWith({
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
