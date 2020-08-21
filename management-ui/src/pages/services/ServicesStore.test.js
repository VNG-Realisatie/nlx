// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../../test-utils/deferred-promise'
import ServicesStore, { createServicesStore } from './ServicesStore'

jest.mock('../../models/ServiceModel', () => ({
  createService: ({ service }) => ({ ...service }),
}))

let rootStore
let domain

beforeEach(() => {
  rootStore = {}
  domain = {}
})

test('createServicesStore returns an instance', () => {
  const directoryStore = createServicesStore({ rootStore, domain })
  expect(directoryStore).toBeInstanceOf(ServicesStore)
})

test('fetching services', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn(() => request),
  }

  const serviceList = [{ name: 'Service A' }, { name: 'Service B' }]

  const servicesStore = new ServicesStore({ rootStore, domain })

  expect(servicesStore.services).toEqual([])

  servicesStore.fetchServices()

  expect(servicesStore.isReady).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.resolve(serviceList)

  expect(servicesStore.services).toEqual(serviceList)
  expect(servicesStore.isReady).toBe(true)
})

test('handle error while fetching services', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn(() => request),
  }

  const servicesStore = new ServicesStore({ rootStore, domain })

  expect(servicesStore.services).toEqual([])

  servicesStore.fetchServices()

  expect(servicesStore.isReady).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(servicesStore.error).toEqual('some error')
  expect(servicesStore.services).toEqual([])
  expect(servicesStore.isReady).toBe(true)
})

test('selecting a service', () => {
  const serviceList = [{ name: 'Service A' }, { name: 'Service B' }]

  const servicesStore = new ServicesStore({ rootStore, domain })
  servicesStore.services = serviceList

  const selectedService = servicesStore.selectService('Service A')

  expect(selectedService).toEqual(serviceList[0])
})

test('removing a service', async () => {
  const serviceList = [{ name: 'Service A' }, { name: 'Service B' }]
  domain = { remove: jest.fn() }

  const servicesStore = new ServicesStore({ rootStore, domain })
  servicesStore.services = serviceList

  const selectedService = servicesStore.selectService('Service A')

  await servicesStore.removeService(selectedService)

  expect(domain.remove).toHaveBeenCalled()
  expect(servicesStore.services).not.toContain(selectedService)
})

test('adding a service', async () => {
  const serviceList = [{ name: 'Service A' }, { name: 'Service B' }]
  domain = {
    create: jest.fn((service) => ({ ...service })),
  }

  const servicesStore = new ServicesStore({ rootStore, domain })
  servicesStore.services = serviceList

  const newService = { name: 'Service C' }
  await servicesStore.addService(newService)

  expect(domain.create).toHaveBeenCalled()
  expect(servicesStore.services).toContainEqual(newService)
})
