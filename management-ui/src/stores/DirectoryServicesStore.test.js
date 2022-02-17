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

test('requesting access to a service in the directory', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  jest.spyOn(rootStore.outgoingAccessRequestStore, 'create').mockResolvedValue()

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
    'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
  )

  expect(rootStore.outgoingAccessRequestStore.create).toHaveBeenCalledWith(
    '00000000000000000001',
    'service',
    'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
  )
})
