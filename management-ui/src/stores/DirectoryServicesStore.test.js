// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import DirectoryServiceModel from '../stores/models/DirectoryServiceModel'
import { DirectoryApi, ManagementApi } from '../api'
import { ACCESS_REQUEST_STATES } from './models/OutgoingAccessRequestModel'
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
        organization: {
          serialNumber: '00000000000000000001',
          name: 'Org A',
        },
        serviceName: 'Service A',
      },
      {
        organization: {
          serialNumber: '00000000000000000001',
          name: 'Org A',
        },
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
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Org A',
      },
      serviceName: 'Service A',
    })
    .mockReturnValue({
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Org A',
      },
      serviceName: 'Service A',
      accessStates: [
        {
          accessRequest: {
            id: 'abc',
            state: ACCESS_REQUEST_STATES.APPROVED,
            publicKeyFingerprint:
              'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
          },
          accessProof: { id: 'abc' },
        },
      ],
    })

  const rootStore = new RootStore({
    directoryApiClient,
  })

  const directoryServicesStore = rootStore.directoryServicesStore
  expect(directoryServicesStore.services).toHaveLength(0)

  const initialService = await directoryServicesStore.fetch(
    '00000000000000000001',
    'Service A',
  )

  expect(directoryServicesStore.services).toHaveLength(1)
  expect(initialService).toBeInstanceOf(DirectoryServiceModel)

  const updatedService = await directoryServicesStore.fetch(
    '00000000000000000001',
    'Service A',
  )

  expect(updatedService).toBeInstanceOf(DirectoryServiceModel)
  expect(
    updatedService.hasAccess('h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc='),
  ).toEqual(true)

  expect(initialService).toBe(updatedService)
})

test('fetching a single service which has been removed', async () => {
  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryGetOrganizationService = jest
    .fn()
    .mockResolvedValueOnce({
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Org A',
      },
      serviceName: 'Service A',
    })
    .mockRejectedValue({
      status: 404,
    })

  const rootStore = new RootStore({
    directoryApiClient,
  })

  const directoryServicesStore = rootStore.directoryServicesStore
  expect(directoryServicesStore.services).toHaveLength(0)

  const initialService = await directoryServicesStore.fetch(
    '00000000000000000001',
    'Service A',
  )

  expect(directoryServicesStore.services).toHaveLength(1)
  expect(initialService).toBeInstanceOf(DirectoryServiceModel)

  const updatedService = await directoryServicesStore.fetch(
    '00000000000000000001',
    'Service A',
  )

  expect(directoryServicesStore.services).toHaveLength(0)
  expect(updatedService).toBeUndefined()
})

test('requesting access to a service in the directory', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  managementApiClient.managementSendAccessRequest = jest
    .fn()
    .mockResolvedValue()

  const rootStore = new RootStore({
    managementApiClient,
  })

  const directoryService = new DirectoryServiceModel({
    serviceData: {
      organization: {
        serialNumber: '00000000000000000001',
        name: 'organization',
      },
      serviceName: 'service',
    },
  })

  await rootStore.directoryServicesStore.requestAccess(
    directoryService.organization.serialNumber,
    directoryService.serviceName,
    'public-key-pem',
  )

  expect(managementApiClient.managementSendAccessRequest).toHaveBeenCalledWith({
    organizationSerialNumber: '00000000000000000001',
    serviceName: 'service',
    publicKeyPem: 'public-key-pem',
  })
})

test('retrieving all services with access', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryListServices = jest.fn().mockResolvedValue({
    services: [
      {
        organization: {
          serialNumber: '00000000000000000001',
          name: 'Org A',
        },
        serviceName: 'Service A',
      },
      {
        organization: {
          serialNumber: '00000000000000000001',
          name: 'Org A',
        },
        serviceName: 'Service B',
      },
    ],
  })

  const rootStore = new RootStore({
    directoryApiClient,
    managementApiClient,
  })

  rootStore.accessProofStore.updateFromServer({
    id: '1',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'Org A',
    },
    serviceName: 'Service A',
  })

  await rootStore.directoryServicesStore.fetchAll()

  const serviceWithAccess = rootStore.directoryServicesStore.getService(
    '00000000000000000001',
    'Service A',
  )

  const servicesWithAccess = rootStore.directoryServicesStore.servicesWithAccess

  expect(servicesWithAccess).toHaveLength(1)
  expect(servicesWithAccess[0]).toBe(serviceWithAccess)
})

test('syncing outgoing access requests for a service', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  managementApiClient.managementSynchronizeOutgoingAccessRequests = jest
    .fn()
    .mockResolvedValue({
      foo: 'bar',
    })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const result =
    await rootStore.directoryServicesStore.syncOutgoingAccessRequests(
      '00000000000000000001',
      'Service A',
    )

  expect(
    managementApiClient.managementSynchronizeOutgoingAccessRequests,
  ).toHaveBeenCalledWith({
    organizationSerialNumber: '00000000000000000001',
    serviceName: 'Service A',
  })
  expect(result).toEqual({ foo: 'bar' })
})

test('syncing all outgoing access requests', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  managementApiClient.managementSynchronizeAllOutgoingAccessRequests = jest
    .fn()
    .mockResolvedValue({
      foo: 'bar',
    })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const result =
    await rootStore.directoryServicesStore.syncAllOutgoingAccessRequests()

  expect(
    managementApiClient.managementSynchronizeAllOutgoingAccessRequests,
  ).toHaveBeenCalled()
  expect(result).toEqual({ foo: 'bar' })
})
