// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { ManagementServiceApi } from '../api'
import InwayModel from '../stores/models/InwayModel'
import InwayStore from './InwayStore'

test('initializing the store', () => {
  const managementApiClient = new ManagementServiceApi()

  const store = new InwayStore({
    rootStore: {},
    managementApiClient,
  })

  expect(store.isInitiallyFetched).toBe(false)
  expect(store.inways).toEqual([])
})

test('fetching inways', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListInways = jest
    .fn()
    .mockResolvedValue({
      inways: [{ name: 'Inway A' }, { name: 'Inway B' }],
    })

  const inwayStore = new InwayStore({
    rootStore: {},
    managementApiClient,
  })

  await inwayStore.fetchInways()

  expect(inwayStore.isInitiallyFetched).toBe(true)
  expect(inwayStore.inways).toHaveLength(2)
})

test('fetching a single inway', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceGetInway = jest
    .fn()
    .mockResolvedValue({ inway: { name: 'Inway A' } })

  const inwayStore = new InwayStore({
    rootStore: {},
    managementApiClient,
  })

  const inway = await inwayStore.fetch({ name: 'Inway A' })

  expect(managementApiClient.managementServiceGetInway).toHaveBeenCalledWith({
    name: 'Inway A',
  })
  expect(inway).toBeInstanceOf(InwayModel)
  expect(inway.name).toEqual('Inway A')
})

test('handle error while fetching inways', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListInways = jest
    .fn()
    .mockRejectedValue('arbitrary error')

  const inwayStore = new InwayStore({
    rootStore: {},
    managementApiClient,
  })

  await inwayStore.fetchInways()

  expect(inwayStore.error).toEqual('arbitrary error')
  expect(inwayStore.inways).toHaveLength(0)
  expect(inwayStore.isInitiallyFetched).toBe(true)
})

test('getting an inway', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListInways = jest
    .fn()
    .mockResolvedValue({
      inways: [{ name: 'Inway A' }],
    })

  const inwayStore = new InwayStore({
    rootStore: {},
    managementApiClient,
  })

  await inwayStore.fetchInways()

  let selectedInway = inwayStore.getByName('non-existing-inway-name')
  expect(selectedInway).toBeUndefined()

  selectedInway = inwayStore.getByName('Inway A')
  expect(selectedInway.name).toEqual('Inway A')
})

test('removing an inway', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListInways = jest
    .fn()
    .mockResolvedValueOnce({
      inways: [{ name: 'Inway A' }, { name: 'Inway B' }, { name: 'Inway C' }],
    })
    .mockResolvedValue({
      inways: [{ name: 'Inway A' }, { name: 'Inway C' }],
    })

  managementApiClient.managementServiceDeleteInway = jest
    .fn()
    .mockResolvedValue({})

  const inwayStore = new InwayStore({
    rootStore: {},
    managementApiClient,
  })

  await inwayStore.fetchInways()
  jest.spyOn(inwayStore, 'fetchInways')

  expect(inwayStore.getByName('Inway A')).toBeDefined()
  expect(inwayStore.getByName('Inway B')).toBeDefined()
  expect(inwayStore.getByName('Inway C')).toBeDefined()

  await inwayStore.removeInway('Inway B')

  expect(managementApiClient.managementServiceDeleteInway).toHaveBeenCalledWith(
    {
      name: 'Inway B',
    },
  )

  expect(inwayStore.fetchInways).toHaveBeenCalledTimes(1)
  expect(inwayStore.getByName('Inway A')).toBeDefined()
  expect(inwayStore.getByName('Inway B')).toBeUndefined()
  expect(inwayStore.getByName('Inway C')).toBeDefined()
})
