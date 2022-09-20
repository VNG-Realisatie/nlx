// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter, Routes, Route } from 'react-router-dom'
import { configure } from 'mobx'
import { screen } from '@testing-library/react'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { UserContextProvider } from '../../../user-context'
import { DirectoryApi } from '../../../api'
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
  renderWithAllProviders(
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
        organization: {},
        serviceName: 'Test Service',
      },
    ],
  })

  const rootStore = new RootStore({
    directoryApiClient,
  })
  const fetchAllSpy = jest.spyOn(rootStore.directoryServicesStore, 'fetchAll')
  const syncAllOutgoingAccessRequestsSpy = jest.spyOn(
    rootStore.directoryServicesStore,
    'syncAllOutgoingAccessRequests',
  )

  renderDirectoryPage(rootStore)

  expect(screen.getByRole('progressbar')).toBeInTheDocument()
  expect(
    screen.queryByTestId('mock-directory-services'),
  ).not.toBeInTheDocument()
  expect(
    await screen.findByTestId('mock-directory-services'),
  ).toBeInTheDocument()

  expect(screen.queryByRole('progressbar')).not.toBeInTheDocument()
  expect(fetchAllSpy).toHaveBeenCalledTimes(1)
  expect(syncAllOutgoingAccessRequestsSpy).toHaveBeenCalledTimes(1)
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
