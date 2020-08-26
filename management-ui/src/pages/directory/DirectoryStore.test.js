// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../../test-utils/deferred-promise'
import DirectoryStore, { createDirectoryStore } from './DirectoryStore'
import { mockDirectoryServiceModel } from './DirectoryStore.mock'

jest.mock('../../models/DirectoryServiceModel', () => ({
  createDirectoryService: (...args) => mockDirectoryServiceModel(...args),
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

test('fetching directory services', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn().mockReturnValue(request),
  }

  const serviceList = [
    { organizationName: 'Org A', serviceName: 'Service A' },
    { organizationName: 'Org B', serviceName: 'Service B' },
  ]

  const directoryStore = new DirectoryStore({ rootStore, domain })

  expect(directoryStore.services).toEqual([])
  expect(directoryStore.isInitiallyFetched).toBe(false)

  directoryStore.fetchServices()

  expect(directoryStore.isInitiallyFetched).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.resolve(serviceList)

  await expect(directoryStore.isInitiallyFetched).toBe(true)
  expect(directoryStore.services).toHaveLength(2)
  expect(directoryStore.services).not.toBe([])
})

test('handle error while fetching directory services', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn(() => request),
  }

  const directoryStore = new DirectoryStore({ rootStore, domain })

  expect(directoryStore.services).toEqual([])

  directoryStore.fetchServices()

  expect(directoryStore.isInitiallyFetched).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(directoryStore.error).toEqual('some error')
  expect(directoryStore.services).toEqual([])
  expect(directoryStore.isInitiallyFetched).toBe(true)
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

  const directoryStore = new DirectoryStore({ rootStore, domain })
  directoryStore.services = serviceList

  const selectedService = directoryStore.selectService({
    organizationName: 'Org A',
    serviceName: 'Service A',
  })

  expect(mockDirectoryServiceModelA.fetch).toHaveBeenCalled()
  expect(mockDirectoryServiceModelB.fetch).not.toHaveBeenCalled()
  expect(selectedService.state).toBe('state-a')
})
