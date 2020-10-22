// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../test-utils/deferred-promise'
import ServiceModel from '../models/ServiceModel'
import ServicesStore from './ServicesStore'
import { mockServiceModel } from './ServicesStore.mock'

let rootStore
let serviceRepository
let accessRequestRepository
let accessGrantRepository

beforeEach(() => {
  rootStore = {}
  serviceRepository = {}
  accessRequestRepository = {}
  accessGrantRepository = {}
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
    accessRequestRepository,
    accessGrantRepository,
  })

  expect(servicesStore.isInitiallyFetched).toBe(false)
  expect(servicesStore.services).toEqual([])

  servicesStore.fetchAll()

  expect(servicesStore.isInitiallyFetched).toBe(false)
  expect(serviceRepository.getAll).toHaveBeenCalled()

  await request.resolve(serviceList)

  await expect(servicesStore.isInitiallyFetched).toBe(true)
  expect(servicesStore.services).toHaveLength(2)
  expect(servicesStore.services).not.toBe([])
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
    accessRequestRepository,
    accessGrantRepository,
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

test('selecting a service', () => {
  const mockServiceModelA = mockServiceModel({ name: 'Service A' })
  const mockServiceModelB = mockServiceModel({ name: 'Service B' })
  const serviceList = [mockServiceModelA, mockServiceModelB]

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
    accessRequestRepository,
    accessGrantRepository,
  })
  servicesStore.services = serviceList

  const selectedService = servicesStore.selectService('Service A')

  expect(selectedService.name).toEqual(serviceList[0].name)
  expect(mockServiceModelA.fetch).toHaveBeenCalled()
  expect(mockServiceModelB.fetch).not.toHaveBeenCalled()
})

test('removing a service', async () => {
  const serviceList = [
    mockServiceModel({ name: 'Service A' }),
    mockServiceModel({ name: 'Service B' }),
  ]
  serviceRepository = { remove: jest.fn() }

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
    accessRequestRepository,
    accessGrantRepository,
  })
  servicesStore.services = serviceList

  const selectedService = servicesStore.selectService('Service A')

  await servicesStore.removeService(selectedService)

  expect(serviceRepository.remove).toHaveBeenCalled()
  expect(servicesStore.services).not.toContain(selectedService)
})

test('creating a service', async () => {
  serviceRepository = {
    create: jest.fn((service) => ({ ...service })),
  }

  const servicesStore = new ServicesStore({
    rootStore,
    serviceRepository,
    accessRequestRepository,
    accessGrantRepository,
  })

  expect(servicesStore.services).toHaveLength(0)

  const newService = { name: 'Service A' }
  await servicesStore.create(newService)

  expect(serviceRepository.create).toHaveBeenCalled()
  expect(servicesStore.services[0]).toEqual(
    new ServiceModel({
      store: servicesStore,
      service: {
        name: 'Service A',
      },
    }),
  )
})
