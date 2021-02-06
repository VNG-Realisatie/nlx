// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter, Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { configure } from 'mobx'
import { act, renderWithProviders } from '../../../test-utils'
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

/* eslint-disable react/prop-types */
jest.mock('../DirectoryDetailPage', () => ({ service }) => (
  <div data-testid="mock-directory-service">{service.serviceName}</div>
))
/* eslint-enable react/prop-types */

const renderDirectory = (store) =>
  renderWithProviders(
    <StoreProvider rootStore={store}>
      <UserContextProvider user={{}}>
        <MemoryRouter>
          <DirectoryPage />
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

  const {
    getByRole,
    getByTestId,
    findByTestId,
    queryByTestId,
    queryByRole,
  } = renderDirectory(rootStore)

  expect(fetchAllSpy).toHaveBeenCalled()
  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(queryByTestId('mock-directory-services')).not.toBeInTheDocument()

  expect(await findByTestId('mock-directory-services')).toBeInTheDocument()
  expect(queryByRole('progressbar')).not.toBeInTheDocument()
  expect(rootStore.directoryServicesStore.isInitiallyFetched).toEqual(true)
  expect(getByTestId('mock-directory-service-0')).toHaveTextContent(
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

  const { findByTestId, getByTestId } = renderDirectory(rootStore)

  expect(await findByTestId('mock-directory-services')).toBeInTheDocument()
  expect(() => getByTestId('mock-directory-service-0')).toThrow()
})

test('failed to load services', async () => {
  global.console.error = jest.fn()

  const rootStore = new RootStore({
    directoryRepository: {
      getAll: jest.fn().mockRejectedValue('There is an error'),
    },
  })

  const { findByTestId, getByTestId } = renderDirectory(rootStore)

  expect(await findByTestId('error-message')).toHaveTextContent(
    /^Failed to load the directory$/,
  )
  expect(() => getByTestId('mock-directory-services')).toThrow()

  global.console.error.mockRestore()
})

test('navigating to the detail page should re-fetch the directory model', async () => {
  // NOTE: we open the overview page before navigating to
  // the detail page this allows us to first put a spy on
  // the fetch-method of the ServiceDirectory model

  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryListServices = jest.fn().mockResolvedValue({
    services: [
      {
        organizationName: 'foo',
        serviceName: 'bar',
        state: 'up',
      },
    ],
  })

  const rootStore = new RootStore({
    directoryApiClient,
  })

  const history = createMemoryHistory({ initialEntries: ['/directory'] })

  await act(async () => {
    renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <UserContextProvider user={{}}>
          <Router history={history}>
            <DirectoryPage />
          </Router>
        </UserContextProvider>
      </StoreProvider>,
    )
  })

  const serviceModel = rootStore.directoryServicesStore.getService({
    organizationName: 'foo',
    serviceName: 'bar',
  })
  jest.spyOn(serviceModel, 'fetch').mockResolvedValue({})

  history.push('/directory/foo/bar')

  expect(serviceModel.fetch).toHaveBeenCalledTimes(1)
})
