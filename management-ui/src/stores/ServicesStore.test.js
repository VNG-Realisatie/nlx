// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import ServiceModel from '../models/ServiceModel'
import { ManagementApi } from '../api'
import ServicesStore from './ServicesStore'
import { RootStore } from './index'

let rootStore
let serviceRepository

beforeEach(() => {
  rootStore = {}
  serviceRepository = {}
})

test('initializing the store', async () => {
  const managementApiClient = new ManagementApi()

  const servicesStore = new ServicesStore({
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

  servicesStore.services = [
    new ServiceModel({ servicesStore, serviceData: { name: 'Service A' } }),
    new ServiceModel({ servicesStore, serviceData: { name: 'Service B' } }),
  ]

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

  const servicesStore = new ServicesStore({
    rootStore,
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

  const servicesStore = new ServicesStore({
    rootStore,
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

  const servicesStore = new ServicesStore({
    rootStore,
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
  const serviceList = [{ name: 'Service A' }, { name: 'Service B' }]
  serviceRepository = { remove: jest.fn() }

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
  })
  servicesStore.services = serviceList

  const service = servicesStore.getService('Service A')
  await servicesStore.removeService(service)

  expect(serviceRepository.remove).toHaveBeenCalled()
  expect(servicesStore.services).not.toContain(service)
})
