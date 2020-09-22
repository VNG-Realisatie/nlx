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
let inwayRepository

beforeEach(() => {
  rootStore = {}
  inwayRepository = {}
})

test('createInwaysStore returns an instance', () => {
  const store = createInwaysStore({ rootStore, inwayRepository })
  expect(store).toBeInstanceOf(InwaysStore)
})

test('fetching inways', async () => {
  const request = deferredPromise()
  inwayRepository = {
    getAll: jest.fn(() => request),
  }

  const inwaysStore = new InwaysStore({
    rootStore,
    inwayRepository,
  })

  expect(inwaysStore.isInitiallyFetched).toBe(false)
  expect(inwaysStore.inways).toEqual([])

  inwaysStore.fetchInways()

  expect(inwaysStore.isInitiallyFetched).toBe(false)
  expect(inwayRepository.getAll).toHaveBeenCalled()

  const inwaysList = [{ name: 'Inway A' }, { name: 'Inway B' }]
  await request.resolve(inwaysList)

  expect(inwaysStore.isInitiallyFetched).toBe(true)
  expect(inwaysStore.inways).toHaveLength(2)
  expect(inwaysStore.inways).not.toBe([])
})

test('handle error while fetching inways', async () => {
  const request = deferredPromise()
  inwayRepository = {
    getAll: jest.fn(() => request),
  }

  const inwaysStore = new InwaysStore({
    rootStore,
    inwayRepository,
  })

  expect(inwaysStore.inways).toEqual([])

  inwaysStore.fetchInways()

  expect(inwaysStore.isInitiallyFetched).toBe(false)
  expect(inwayRepository.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(inwaysStore.error).toEqual('some error')
  expect(inwaysStore.inways).toEqual([])
  expect(inwaysStore.isInitiallyFetched).toBe(true)
})

test('selecting an inway', async () => {
  const mockInwayModelA = mockInwayModel({ name: 'Inway A' })
  const mockInwayModelB = mockInwayModel({ name: 'Inway B' })
  const mockInwayModelC = mockInwayModel({ name: 'Inway C' })

  inwayRepository = {
    getAll: jest
      .fn()
      .mockResolvedValue([mockInwayModelA, mockInwayModelB, mockInwayModelC]),
  }

  const inwaysStore = new InwaysStore({ rootStore, inwayRepository })
  await inwaysStore.fetchInways()

  const selectedInway = inwaysStore.selectInway('Inway B')

  expect(selectedInway.name).toEqual('Inway B')
  expect(mockInwayModelB.fetch).toHaveBeenCalled()

  expect(mockInwayModelA.fetch).not.toHaveBeenCalled()
  expect(mockInwayModelC.fetch).not.toHaveBeenCalled()
})

test('selecting an inway that is not present in the store', async () => {
  const mockInwayModelA = mockInwayModel({ name: 'Inway A' })
  const mockInwayModelB = mockInwayModel({ name: 'Inway B' })
  const mockInwayModelC = mockInwayModel({ name: 'Inway C' })

  inwayRepository = {
    getAll: jest
      .fn()
      .mockResolvedValue([mockInwayModelA, mockInwayModelB, mockInwayModelC]),
  }

  const inwaysStore = new InwaysStore({ rootStore, inwayRepository })
  await inwaysStore.fetchInways()

  const selectedInway = inwaysStore.selectInway('arbitrary inway name')
  expect(selectedInway).toBeUndefined()

  expect(mockInwayModelA.fetch).not.toHaveBeenCalled()
  expect(mockInwayModelB.fetch).not.toHaveBeenCalled()
  expect(mockInwayModelC.fetch).not.toHaveBeenCalled()
})
