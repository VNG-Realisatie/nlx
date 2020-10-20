// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { act } from '@testing-library/react'
import deferredPromise from '../test-utils/deferred-promise'
import DirectoryServiceModel from '../models/DirectoryServiceModel'
import OutgoingAccessRequestModel from '../models/OutgoingAccessRequestModel'
import DirectoryServicesStore from './DirectoryServicesStore'
import { mockDirectoryServiceModel } from './DirectoryServicesStore.mock'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'

let rootStore
let directoryRepository

beforeEach(() => {
  rootStore = {}
  directoryRepository = {}
})

test('fetching directory services', async () => {
  const request = deferredPromise()
  directoryRepository = {
    getAll: jest.fn().mockReturnValue(request),
  }

  const serviceList = [
    { organizationName: 'Org A', serviceName: 'Service A' },
    { organizationName: 'Org B', serviceName: 'Service B' },
  ]

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore,
    directoryRepository,
  })

  expect(directoryServicesStore.services).toEqual([])
  expect(directoryServicesStore.isInitiallyFetched).toBe(false)

  directoryServicesStore.fetchServices()

  expect(directoryServicesStore.isInitiallyFetched).toBe(false)
  expect(directoryRepository.getAll).toHaveBeenCalled()

  await act(async () => {
    await request.resolve(serviceList)
  })

  await expect(directoryServicesStore.isInitiallyFetched).toBe(true)
  expect(directoryServicesStore.services).toHaveLength(2)
  expect(directoryServicesStore.services).not.toBe([])
})

test('handle error while fetching directory services', async () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})

  const request = deferredPromise()
  directoryRepository = {
    getAll: jest.fn(() => request),
  }

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore,
    directoryRepository,
  })

  expect(directoryServicesStore.services).toEqual([])

  directoryServicesStore.fetchServices()

  expect(directoryServicesStore.isInitiallyFetched).toBe(false)
  expect(directoryRepository.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(directoryServicesStore.error).toEqual('some error')
  expect(directoryServicesStore.services).toEqual([])
  expect(directoryServicesStore.isInitiallyFetched).toBe(true)

  errorSpy.mockRestore()
})

test('selecting a directory service', () => {
  const mockDirectoryServiceModelA = mockDirectoryServiceModel({
    organizationName: 'Org A',
    serviceName: 'Service A',
    state: 'state-a',
  })
  const mockDirectoryServiceModelB = mockDirectoryServiceModel({
    organizationName: 'Org B',
    serviceName: 'Service B',
    state: 'state-b',
  })
  const serviceList = [mockDirectoryServiceModelA, mockDirectoryServiceModelB]

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore,
    directoryRepository,
  })
  directoryServicesStore.services = serviceList

  const selectedService = directoryServicesStore.selectService({
    organizationName: 'Org A',
    serviceName: 'Service A',
  })

  expect(mockDirectoryServiceModelA.fetch).toHaveBeenCalled()
  expect(mockDirectoryServiceModelB.fetch).not.toHaveBeenCalled()
  expect(selectedService.state).toBe('state-a')
})

test('requesting access to a service in the directory', async () => {
  const accessRequestRepository = {
    createAccessRequest: jest.fn().mockResolvedValue({
      id: '42',
    }),
  }

  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
    accessRequestRepository: accessRequestRepository,
  })

  const outgoingAccessRequestStoreCreateSpy = jest.spyOn(
    outgoingAccessRequestStore,
    'create',
  )

  const directoryServicesStore = new DirectoryServicesStore({
    rootStore: {
      outgoingAccessRequestsStore: outgoingAccessRequestStore,
    },
    directoryRepository,
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
