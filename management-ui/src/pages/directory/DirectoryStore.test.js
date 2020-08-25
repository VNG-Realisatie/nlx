// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../../test-utils/deferred-promise'
import DirectoryStore, { createDirectoryStore } from './DirectoryStore'

jest.mock('../../models/DirectoryServiceModel', () => ({
  createDirectoryService: ({ service }) => ({ ...service }),
}))

let rootStore
let domain

beforeEach(() => {
  rootStore = {}
  domain = {}
})

test('createDirectoryStore returns an instance', () => {
  const directoryStore = createDirectoryStore({ rootStore, domain })
  expect(directoryStore).toBeInstanceOf(DirectoryStore)
})

test('fetching services', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn(() => request),
  }

  const serviceList = [
    { organizationName: 'Org A', serviceName: 'Service A' },
    { organizationName: 'Org B', serviceName: 'Service B' },
  ]

  const directoryStore = new DirectoryStore({ rootStore, domain })

  expect(directoryStore.services).toEqual([])

  directoryStore.fetchServices()

  expect(directoryStore.isReady).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.resolve(serviceList)

  expect(directoryStore.services).toEqual(serviceList)
  expect(directoryStore.isReady).toBe(true)
})

test('handle error while fetching services', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn(() => request),
  }

  const directoryStore = new DirectoryStore({ rootStore, domain })

  expect(directoryStore.services).toEqual([])

  directoryStore.fetchServices()

  expect(directoryStore.isReady).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(directoryStore.error).toEqual('some error')
  expect(directoryStore.services).toEqual([])
  expect(directoryStore.isReady).toBe(true)
})

test('selecting a service', () => {
  const serviceList = [
    { organizationName: 'Org A', serviceName: 'Service A', state: 'up' },
    { organizationName: 'Org B', serviceName: 'Service B', state: 'down' },
  ]

  const directoryStore = new DirectoryStore({ rootStore, domain })
  directoryStore.services = serviceList

  const selectedService = directoryStore.selectService({
    organizationName: 'Org A',
    serviceName: 'Service A',
  })

  expect(selectedService).toEqual(serviceList[0])
})
