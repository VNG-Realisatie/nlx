// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import InwayModel from '../stores/models/InwayModel'
import InwayStore from './InwayStore'
import { RootStore } from './index'

test('initializing the store', () => {
  const managementApiClient = new ManagementApi()

  const store = new InwayStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  expect(store.isInitiallyFetched).toBe(false)
  expect(store.inways).toEqual([])
})

test('fetching inways', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [{ name: 'Inway A' }, { name: 'Inway B' }],
  })

  const inwayStore = new InwayStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  await inwayStore.fetchInways()

  expect(inwayStore.isInitiallyFetched).toBe(true)
  expect(inwayStore.inways).toHaveLength(2)
})

test('fetching a single inway', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetInway = jest
    .fn()
    .mockResolvedValue({ name: 'Inway A' })

  const inwayStore = new InwayStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  const inway = await inwayStore.fetch({ name: 'Inway A' })

  expect(managementApiClient.managementGetInway).toHaveBeenCalledWith({
    name: 'Inway A',
  })
  expect(inway).toBeInstanceOf(InwayModel)
  expect(inway.name).toEqual('Inway A')
})

test('handle error while fetching inways', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListInways = jest
    .fn()
    .mockRejectedValue('arbitrary error')

  const inwayStore = new InwayStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  await inwayStore.fetchInways()

  expect(inwayStore.error).toEqual('arbitrary error')
  expect(inwayStore.inways).toHaveLength(0)
  expect(inwayStore.isInitiallyFetched).toBe(true)
})

test('getting an inway', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [{ name: 'Inway A' }],
  })

  const inwayStore = new InwayStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  await inwayStore.fetchInways()

  let selectedInway = inwayStore.getInway({ name: 'non-existing-inway-name' })
  expect(selectedInway).toBeUndefined()

  selectedInway = inwayStore.getInway({ name: 'Inway A' })
  expect(selectedInway.name).toEqual('Inway A')
})
