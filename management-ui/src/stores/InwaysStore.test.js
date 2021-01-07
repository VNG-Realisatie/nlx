// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import InwayModel from '../models/InwayModel'
import InwaysStore from './InwaysStore'
import { RootStore } from './index'

test('initializing the store', () => {
  const managementApiClient = new ManagementApi()

  const store = new InwaysStore({
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

  const inwaysStore = new InwaysStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  await inwaysStore.fetchInways()

  expect(inwaysStore.isInitiallyFetched).toBe(true)
  expect(inwaysStore.inways).toHaveLength(2)
})

test('fetching a single inway', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetInway = jest
    .fn()
    .mockResolvedValue({ name: 'Inway A' })

  const inwaysStore = new InwaysStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  const inway = await inwaysStore.fetch({ name: 'Inway A' })

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

  const inwaysStore = new InwaysStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  await inwaysStore.fetchInways()

  expect(inwaysStore.error).toEqual('arbitrary error')
  expect(inwaysStore.inways).toHaveLength(0)
  expect(inwaysStore.isInitiallyFetched).toBe(true)
})

test('getting an inway', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [{ name: 'Inway A' }],
  })

  const inwaysStore = new InwaysStore({
    rootStore: new RootStore(),
    managementApiClient,
  })

  await inwaysStore.fetchInways()

  let selectedInway = inwaysStore.getInway({ name: 'non-existing-inway-name' })
  expect(selectedInway).toBeUndefined()

  selectedInway = inwaysStore.getInway({ name: 'Inway A' })
  expect(selectedInway.name).toEqual('Inway A')
})
