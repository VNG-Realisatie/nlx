// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../test-utils/deferred-promise'
import ServiceModel from '../models/ServiceModel'
import ServicesStore from './ServicesStore'
import { RootStore } from './index'

let rootStore
let serviceRepository

beforeEach(() => {
  rootStore = {}
  serviceRepository = {}
})

test('initializing the store', async () => {
  const servicesStore = new ServicesStore({
    rootStore: new RootStore(),
  })

  expect(servicesStore.services).toEqual([])
  expect(servicesStore.isInitiallyFetched).toBe(false)
  expect(servicesStore.isFetching).toBe(false)
})

test('fetch a single service', async () => {
  rootStore = {
    incomingAccessRequestsStore: {
      fetchForService: jest.fn(),
    },
    accessGrantStore: {
      fetchForService: jest.fn(),
    },
  }

  serviceRepository = {
    getByName: jest.fn().mockResolvedValue({ name: 'Service A' }),
  }

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
  })

  servicesStore.services = [
    new ServiceModel({ servicesStore, serviceData: { name: 'Service A' } }),
    new ServiceModel({ servicesStore, serviceData: { name: 'Serivce B' } }),
  ]

  await servicesStore.fetch({ name: 'Service A' })

  expect(serviceRepository.getByName).toHaveBeenCalled()
  expect(
    rootStore.incomingAccessRequestsStore.fetchForService,
  ).toHaveBeenCalled()
  expect(rootStore.accessGrantStore.fetchForService).toHaveBeenCalled()
})

test('fetching services', async () => {
  const request = deferredPromise()
  serviceRepository = {
    getAll: jest.fn(() => request),
  }

  const serviceList = [{ name: 'Service A' }, { name: 'Service B' }]
  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
  })

  expect(servicesStore.isInitiallyFetched).toBe(false)
  expect(servicesStore.services).toEqual([])

  servicesStore.fetchAll()

  expect(servicesStore.isInitiallyFetched).toBe(false)
  expect(serviceRepository.getAll).toHaveBeenCalled()

  await request.resolve(serviceList)

  await expect(servicesStore.isInitiallyFetched).toBe(true)
  expect(servicesStore.services).toHaveLength(2)
})

test('handle error while fetching services', async () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})

  const request = deferredPromise()
  serviceRepository = {
    getAll: jest.fn(() => request),
  }

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
  })

  expect(servicesStore.services).toEqual([])

  servicesStore.fetchAll()

  expect(servicesStore.isInitiallyFetched).toBe(false)
  expect(serviceRepository.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(servicesStore.error).toEqual('some error')
  expect(servicesStore.services).toEqual([])
  expect(servicesStore.isInitiallyFetched).toBe(true)

  errorSpy.mockRestore()
})

test('getting a service', () => {
  const serviceList = [{ name: 'Service A' }, { name: 'Service B' }]

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
  })
  servicesStore.services = serviceList

  const service = servicesStore.getService('Service A')
  expect(service.name).toEqual(serviceList[0].name)
})

test('creating a service', async () => {
  const formData = {
    name: 'New service',
    endpointURL: 'api.io',
  }

  serviceRepository = { create: jest.fn().mockResolvedValue(formData) }

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
  })

  await servicesStore.create(formData)

  expect(serviceRepository.create).toHaveBeenCalled()
  expect(servicesStore.services).toHaveLength(1)
  expect(servicesStore.services[0]).toBeInstanceOf(ServiceModel)
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
