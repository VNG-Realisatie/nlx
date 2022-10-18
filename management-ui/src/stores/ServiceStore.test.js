// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { ManagementServiceApi } from '../api'
import ServiceModel from './models/ServiceModel'
import { RootStore } from './index'

test('initializing the store', async () => {
  const managementApiClient = new ManagementServiceApi()
  const rootStore = new RootStore({
    managementApiClient,
  })
  const serviceStore = rootStore.servicesStore

  expect(serviceStore.services).toEqual([])
  expect(serviceStore.isInitiallyFetched).toBe(false)
  expect(serviceStore.isFetching).toBe(false)
})

test('fetch and getting a single service', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceGetService = jest
    .fn()
    .mockResolvedValue({
      name: 'Service A',
    })

  const rootStore = new RootStore({
    managementApiClient,
  })
  const serviceStore = rootStore.servicesStore

  jest
    .spyOn(rootStore.incomingAccessRequestsStore, 'fetchForService')
    .mockResolvedValue([])

  jest
    .spyOn(rootStore.accessGrantStore, 'fetchForService')
    .mockResolvedValue([])

  await serviceStore.fetch({ name: 'Service A' })

  expect(
    rootStore.incomingAccessRequestsStore.fetchForService,
  ).toHaveBeenCalled()
  expect(rootStore.accessGrantStore.fetchForService).toHaveBeenCalled()

  expect(serviceStore.getByName('Service A').name).toEqual('Service A')
})

test('fetching services', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [{ name: 'Service A' }, { name: 'Service B' }],
    })

  const rootStore = new RootStore({
    managementApiClient,
  })
  const serviceStore = rootStore.servicesStore

  await serviceStore.fetchAll()

  expect(serviceStore.isInitiallyFetched).toBe(true)
  expect(serviceStore.services).toHaveLength(2)
})

test('handle error while fetching services', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockRejectedValue('arbitrary error')

  const rootStore = new RootStore({
    managementApiClient,
  })
  const serviceStore = rootStore.servicesStore

  await serviceStore.fetchAll()

  expect(serviceStore.error).toEqual('arbitrary error')
  expect(serviceStore.services).toEqual([])
  expect(serviceStore.isInitiallyFetched).toBe(true)
})

test('creating a service', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceCreateService = jest
    .fn()
    .mockResolvedValue({
      name: 'New service',
      endpointUrl: 'api.io',
    })

  const rootStore = new RootStore({
    managementApiClient,
  })
  const serviceStore = rootStore.servicesStore

  const service = await serviceStore.create({
    name: 'New service',
    endpointUrl: 'api.io',
  })

  expect(service).toBeInstanceOf(ServiceModel)
  expect(serviceStore.services[0]).toBe(service)
})

test('removing a service', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [{ name: 'Service A' }],
    })

  managementApiClient.managementServiceDeleteService = jest
    .fn()
    .mockResolvedValue()

  const rootStore = new RootStore({
    managementApiClient,
  })
  const serviceStore = rootStore.servicesStore

  await serviceStore.fetchAll()
  expect(serviceStore.getByName('Service A')).toBeInstanceOf(ServiceModel)

  await serviceStore.removeService('Service A')
  expect(
    managementApiClient.managementServiceDeleteService,
  ).toHaveBeenCalledWith({
    name: 'Service A',
  })
})

test('fetching statistics', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [{ name: 'Service', incomingAccessRequestCount: 0 }],
    })

  managementApiClient.managementServiceGetStatisticsOfServices = jest
    .fn()
    .mockResolvedValue({
      services: [{ name: 'Service', incomingAccessRequestCount: 1 }],
    })

  const rootStore = new RootStore({
    managementApiClient,
  })
  const serviceStore = rootStore.servicesStore

  await serviceStore.fetchAll()

  expect(serviceStore.getByName('Service').incomingAccessRequestCount).toBe(0)

  await serviceStore.fetchStats()

  expect(
    managementApiClient.managementServiceGetStatisticsOfServices,
  ).toHaveBeenCalledTimes(1)
  expect(serviceStore.getByName('Service').incomingAccessRequestCount).toBe(1)
})
