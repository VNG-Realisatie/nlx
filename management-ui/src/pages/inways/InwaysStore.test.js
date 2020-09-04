// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../../test-utils/deferred-promise'
import InwaysStore, { createInwaysStore } from './InwaysStore'
import { mockInwayModel } from './InwaysStore.mock'

jest.mock('../../models/InwayModel', () => ({
  createInway: ({ inway }) => ({ ...inway }),
}))

let rootStore
let domain

beforeEach(() => {
  rootStore = {}
  domain = {}
})

test('createInwaysStore returns an instance', () => {
  const store = createInwaysStore({ rootStore, domain })
  expect(store).toBeInstanceOf(InwaysStore)
})

test('fetching inways', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn(() => request),
  }

  const inwaysList = [{ name: 'Inway A' }, { name: 'Inway B' }]

  const inwaysStore = new InwaysStore({ rootStore, domain })

  expect(inwaysStore.isInitiallyFetched).toBe(false)
  expect(inwaysStore.inways).toEqual([])

  inwaysStore.fetchInways()

  expect(inwaysStore.isInitiallyFetched).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.resolve(inwaysList)

  await expect(inwaysStore.isInitiallyFetched).toBe(true)
  expect(inwaysStore.inways).toHaveLength(2)
  expect(inwaysStore.inways).not.toBe([])
})

test('handle error while fetching inways', async () => {
  const request = deferredPromise()
  domain = {
    getAll: jest.fn(() => request),
  }

  const inwaysStore = new InwaysStore({ rootStore, domain })

  expect(inwaysStore.inways).toEqual([])

  inwaysStore.fetchInways()

  expect(inwaysStore.isInitiallyFetched).toBe(false)
  expect(domain.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(inwaysStore.error).toEqual('some error')
  expect(inwaysStore.inways).toEqual([])
  expect(inwaysStore.isInitiallyFetched).toBe(true)
})

test('selecting a service', () => {
  const mockInwayModelA = mockInwayModel({ name: 'Inway A' })
  const mockInwayModelB = mockInwayModel({ name: 'Inway B' })
  const inwayList = [mockInwayModelA, mockInwayModelB]

  const inwaysStore = new InwaysStore({ rootStore, domain })
  inwaysStore.inways = inwayList

  const selectedService = inwaysStore.selectInway('Inway A')

  expect(selectedService).toEqual(inwayList[0])
  expect(mockInwayModelA.fetch).toHaveBeenCalled()
  expect(mockInwayModelB.fetch).not.toHaveBeenCalled()
})
