// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import ServiceModel from './models/ServiceModel'
import ServiceStore from './ServiceStore'
import { RootStore } from './index'

test('initializing the store', async () => {
  const managementApiClient = new ManagementApi()

  const servicesStore = new ServiceStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  expect(servicesStore.services).toEqual([])
  expect(servicesStore.isInitiallyFetched).toBe(false)
  expect(servicesStore.isFetching).toBe(false)
})

test('fetch and getting a single service', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetService = jest.fn().mockResolvedValue({
    name: 'Service A',
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const servicesStore = rootStore.servicesStore

  jest
    .spyOn(rootStore.incomingAccessRequestsStore, 'fetchForService')
    .mockResolvedValue([])

  jest
    .spyOn(rootStore.accessGrantStore, 'fetchForService')
    .mockResolvedValue([])

  await servicesStore.fetch({ name: 'Service A' })

  expect(
    rootStore.incomingAccessRequestsStore.fetchForService,
  ).toHaveBeenCalled()
  expect(rootStore.accessGrantStore.fetchForService).toHaveBeenCalled()

  expect(servicesStore.getService('Service A').name).toEqual('Service A')
})

test('fetching services', async () => {
  const managementApi = new ManagementApi()

  managementApi.managementListServices = jest.fn().mockResolvedValue({
    services: [{ name: 'Service A' }, { name: 'Service B' }],
  })

  const servicesStore = new ServiceStore({
    rootStore: new RootStore(),
    managementApiClient: managementApi,
  })

  await servicesStore.fetchAll()

  expect(servicesStore.isInitiallyFetched).toBe(true)
  expect(servicesStore.services).toHaveLength(2)
})

test('handle error while fetching services', async () => {
  const managementApi = new ManagementApi()

  managementApi.managementListServices = jest
    .fn()
    .mockRejectedValue('arbitrary error')

  const servicesStore = new ServiceStore({
    rootStore: new RootStore(),
    managementApiClient: managementApi,
  })

  await servicesStore.fetchAll()

  expect(servicesStore.error).toEqual('arbitrary error')
  expect(servicesStore.services).toEqual([])
  expect(servicesStore.isInitiallyFetched).toBe(true)
})

test('creating a service', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementCreateService = jest.fn().mockResolvedValue({
    name: 'New service',
    endpointURL: 'api.io',
  })

  const servicesStore = new ServiceStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  const service = await servicesStore.create({
    name: 'New service',
    endpointURL: 'api.io',
  })

  expect(service).toBeInstanceOf(ServiceModel)
  expect(servicesStore.services[0]).toBe(service)
})

test('removing a service', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [{ name: 'Service A' }],
  })

  managementApiClient.managementDeleteService = jest.fn().mockResolvedValue()

  const servicesStore = new ServiceStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  await servicesStore.fetchAll()
  expect(servicesStore.getService('Service A')).toBeInstanceOf(ServiceModel)

  await servicesStore.removeService('Service A')
  expect(managementApiClient.managementDeleteService).toHaveBeenCalledWith({
    name: 'Service A',
  })
})
