// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import {
  MemoryRouter,
  Routes,
  Route,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { configure } from 'mobx'
import { screen, waitFor } from '@testing-library/react'
import { act, renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { UserContextProvider } from '../../../user-context'
import { DirectoryApi, ManagementApi } from '../../../api'
import DirectoryPage from './index'

jest.mock('../../../components/PageTemplate')
jest.mock('../../../components/OrganizationName', () => () => null)

/* eslint-disable react/prop-types */
jest.mock('./components/DirectoryPageView', () => ({ services }) => {
  return (
    <div data-testid="mock-directory-services">
      {services.map((o, i) => (
        <span key={i} data-testid={`mock-directory-service-${i}`}>
          {o.serviceName}
        </span>
      ))}
    </div>
  )
})
/* eslint-enable react/prop-types */

jest.mock('../../../domain/environment-repository', () => ({
  getCurrent: async () => ({
    organizationSerialNumber: '12345678901234567890',
  }),
}))

const renderDirectoryPage = (store) =>
  renderWithProviders(
    <StoreProvider rootStore={store}>
      <UserContextProvider user={{}}>
        <MemoryRouter>
          <Routes>
            <Route path="*" element={<DirectoryPage />} />
          </Routes>
        </MemoryRouter>
      </UserContextProvider>
    </StoreProvider>,
  )

test('listing all services', async () => {
  configure({ safeDescriptors: false })
  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryListServices = jest.fn().mockResolvedValue({
    services: [
      {
        serviceName: 'Test Service',
      },
    ],
  })

  const rootStore = new RootStore({
    directoryApiClient,
  })
  const fetchAllSpy = jest.spyOn(rootStore.directoryServicesStore, 'fetchAll')

  renderDirectoryPage(rootStore)

  expect(fetchAllSpy).toHaveBeenCalledTimes(1)
  expect(screen.getByRole('progressbar')).toBeInTheDocument()
  expect(
    screen.queryByTestId('mock-directory-services'),
  ).not.toBeInTheDocument()

  expect(
    await screen.findByTestId('mock-directory-services'),
  ).toBeInTheDocument()
  expect(screen.queryByRole('progressbar')).not.toBeInTheDocument()
  expect(rootStore.directoryServicesStore.isInitiallyFetched).toEqual(true)
  expect(screen.getByTestId('mock-directory-service-0')).toHaveTextContent(
    'Test Service',
  )
})

test('no services', async () => {
  const directoryApiClient = new DirectoryApi()
  directoryApiClient.directoryListServices = jest.fn().mockResolvedValue({
    services: [],
  })

  const rootStore = new RootStore({
    directoryApiClient,
  })

  renderDirectoryPage(rootStore)

  expect(
    await screen.findByTestId('mock-directory-services'),
  ).toBeInTheDocument()
  expect(() => screen.getByTestId('mock-directory-service-0')).toThrow()
})

test('failed to load services', async () => {
  global.console.error = jest.fn()

  const rootStore = new RootStore({
    directoryRepository: {
      getAll: jest.fn().mockRejectedValue('There is an error'),
    },
  })

  renderDirectoryPage(rootStore)

  expect(await screen.findByTestId('error-message')).toHaveTextContent(
    /^Failed to load the directory$/,
  )
  expect(() => screen.getByTestId('mock-directory-services')).toThrow()

  global.console.error.mockRestore()
})

test('navigating to the detail page should re-fetch the directory model', async () => {
  // NOTE: we open the overview page before navigating to
  // the detail page this allows us to first put a spy on
  // the fetch-method of the ServiceDirectory model

  configure({ safeDescriptors: false })

  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryListServices = jest.fn().mockResolvedValue({
    services: [
      {
        organization: {
          serialNumber: '00000000000000000001',
          name: 'foo',
        },
        serviceName: 'bar',
        state: 'up',
      },
    ],
  })

  directoryApiClient.directoryGetOrganizationService = jest
    .fn()
    .mockResolvedValue({})

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [],
  })

  const rootStore = new RootStore({
    directoryApiClient,
    managementApiClient,
  })

  const history = createMemoryHistory()

  renderWithProviders(
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

  jest.spyOn(rootStore.directoryServicesStore, 'fetch')

  act(() => {
    history.push('/00000000000000000001/bar')
  })

  await waitFor(() => {
    expect(rootStore.directoryServicesStore.fetch).toHaveBeenNthCalledWith(
      1,
      '00000000000000000001',
      'bar',
    )
  })
})
