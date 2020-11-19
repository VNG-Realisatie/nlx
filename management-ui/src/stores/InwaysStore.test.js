// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../test-utils/deferred-promise'
import InwaysStore from './InwaysStore'

let rootStore
let inwayRepository

beforeEach(() => {
  rootStore = {}
  inwayRepository = {}
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
  const inwaysStore = new InwaysStore({
    rootStore,
    inwayRepository: {
      getAll: jest.fn().mockResolvedValue([
        {
          name: 'Inway A',
        },
        {
          name: 'Inway B',
        },
      ]),
    },
  })

  await inwaysStore.fetchInways()

  const inwayA = inwaysStore.inways.find((inway) => inway.name === 'Inway A')
  const inwayB = inwaysStore.inways.find((inway) => inway.name === 'Inway B')

  jest.spyOn(inwayA, 'fetch').mockResolvedValue(null)
  jest.spyOn(inwayB, 'fetch').mockResolvedValue(null)

  let selectedInway = inwaysStore.selectInway('non-existing-inway-name')
  expect(selectedInway).toBeUndefined()

  selectedInway = inwaysStore.selectInway('Inway B')
  expect(selectedInway.name).toEqual('Inway B')
  expect(inwayB.fetch).toHaveBeenCalled()
  expect(inwayA.fetch).not.toHaveBeenCalled()
})
