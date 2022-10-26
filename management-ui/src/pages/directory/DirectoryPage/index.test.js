// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter, Routes, Route } from 'react-router-dom'
import { configure } from 'mobx'
import { act, fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { UserContextProvider } from '../../../user-context'
import { DirectoryServiceApi, ManagementServiceApi } from '../../../api'
import { INTERVAL } from '../../../hooks/use-polling'
import deferredPromise from '../../../utils/deferred-promise'
import { SERVICE_STATE_UP } from '../../../components/StateIndicator'
import DirectoryPage from './index'

jest.mock('../DirectoryDetailPage', () => () => null)
jest.mock('../../../components/PageTemplate')
jest.mock('../../../components/OrganizationName', () => () => null)
jest.mock('../../../domain/environment-repository', () => ({
  getCurrent: async () => ({
    organizationSerialNumber: '12345678901234567890',
  }),
}))

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  jest.useRealTimers()
})

test('Directory overview page', async () => {
  configure({ safeDescriptors: false })

  const listServicesResponse = deferredPromise()

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceListServices = jest
    .fn()
    .mockResolvedValue(listServicesResponse)

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceSynchronizeAllOutgoingAccessRequests =
    jest.fn().mockResolvedValue({})

  const rootStore = new RootStore({
    directoryApiClient,
    managementApiClient,
  })

  const synchronizeSpy = jest.spyOn(
    rootStore.directoryServicesStore,
    'syncAllOutgoingAccessRequests',
  )
  const fetchAllSpy = jest.spyOn(rootStore.directoryServicesStore, 'fetchAll')

  const history = createMemoryHistory({ initialEntries: ['/'] })

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <UserContextProvider user={{}}>
        <HistoryRouter history={history}>
          <Routes>
            <Route path="*" element={<DirectoryPage />} />
          </Routes>
        </HistoryRouter>
      </UserContextProvider>
    </StoreProvider>,
  )

  expect(screen.getByRole('progressbar')).toBeInTheDocument()

  listServicesResponse.resolve({
    services: [
      {
        organization: {
          name: 'Test Organization',
          serialNumber: '01234567890123456789',
        },
        serviceName: 'Test Service',
        state: SERVICE_STATE_UP,
      },
    ],
  })

  expect(await screen.findByRole('progressbar')).not.toBeInTheDocument()

  expect(fetchAllSpy).toHaveBeenCalledTimes(1)
  expect(synchronizeSpy).toHaveBeenCalledTimes(1)

  expect(await screen.findByText('Test Service')).toBeInTheDocument()
  expect(await screen.findByText('Test Organization')).toBeInTheDocument()

  await act(async () => {
    jest.advanceTimersByTime(INTERVAL)
  })

  expect(fetchAllSpy).toHaveBeenCalledTimes(1)
  expect(synchronizeSpy).toHaveBeenCalledTimes(2)

  // polling should be paused when opening the directory detail pane
  await act(async () => {
    fireEvent.click(await screen.findByText('Test Service'))
  })

  expect(history.location.pathname).toEqual(
    '/directory/01234567890123456789/Test Service',
  )

  await act(async () => {
    jest.advanceTimersByTime(INTERVAL)
  })

  expect(fetchAllSpy).toHaveBeenCalledTimes(1)
  expect(synchronizeSpy).toHaveBeenCalledTimes(2)
})

test('no services', async () => {
  configure({ safeDescriptors: false })

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [],
    })

  const managementApiClient = new ManagementServiceApi()

  const rootStore = new RootStore({
    directoryApiClient,
    managementApiClient,
  })

  const fetchAllSpy = jest.spyOn(rootStore.directoryServicesStore, 'fetchAll')

  const synchronizeSpy = jest.spyOn(
    rootStore.directoryServicesStore,
    'syncAllOutgoingAccessRequests',
  )

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <UserContextProvider user={{}}>
        <MemoryRouter>
          <Routes>
            <Route path="*" element={<DirectoryPage />} />
          </Routes>
        </MemoryRouter>
      </UserContextProvider>
    </StoreProvider>,
  )

  expect(
    await screen.findByText('There are no services yet'),
  ).toBeInTheDocument()

  expect(fetchAllSpy).toHaveBeenCalledTimes(1)
  expect(synchronizeSpy).toHaveBeenCalledTimes(0)
})

test('failed to load services', async () => {
  global.console.error = jest.fn()

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceListServices = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceSynchronizeAllOutgoingAccessRequests =
    jest.fn().mockResolvedValue({})

  const rootStore = new RootStore({
    directoryApiClient,
    managementApiClient,
  })

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <UserContextProvider user={{}}>
        <MemoryRouter>
          <Routes>
            <Route path="*" element={<DirectoryPage />} />
          </Routes>
        </MemoryRouter>
      </UserContextProvider>
    </StoreProvider>,
  )

  expect(await screen.findByTestId('error-message')).toHaveTextContent(
    /^Failed to load the directory$/,
  )

  global.console.error.mockRestore()
})
