// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { act, renderWithProviders } from '../../../test-utils'
import { StoreProvider } from '../../../stores'
import { UserContextProvider } from '../../../user-context'
import { mockDirectoryServicesStore } from '../../../stores/DirectoryServicesStore.mock'
import DirectoryPage from './index'

// Ignore this deeply nested component which has a separate request flow
jest.mock('../../../components/OrganizationName', () => () => null)

// Simplify showing of the services. We'll only require the serviceName.
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

const renderDirectory = (store) =>
  renderWithProviders(
    <StoreProvider store={store}>
      <UserContextProvider user={{}}>
        <Router>
          <DirectoryPage />
        </Router>
      </UserContextProvider>
    </StoreProvider>,
  )

test('listing all services', async () => {
  const store = mockDirectoryServicesStore({
    isInitiallyFetched: false,
  })
  const fetchAllSpy = jest.spyOn(store.directoryServicesStore, 'fetchAll')

  const { getByRole, getByTestId, findByTestId } = renderDirectory(store)

  expect(fetchAllSpy).toHaveBeenCalled()
  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('mock-directory-services')).toThrow()

  act(() => {
    store.directoryServicesStore.services = [{ serviceName: 'Test Service' }]
    store.directoryServicesStore.isInitiallyFetched = true
  })

  expect(await findByTestId('mock-directory-services')).toBeInTheDocument()
  expect(() => getByRole('progressbar')).toThrow()
  expect(getByTestId('mock-directory-service-0')).toHaveTextContent(
    'Test Service',
  )
})

test('no services', async () => {
  const store = mockDirectoryServicesStore({})

  const { findByTestId, getByTestId } = renderDirectory(store)

  expect(await findByTestId('mock-directory-services')).toBeInTheDocument()
  expect(() => getByTestId('mock-directory-service-0')).toThrow()
})

test('failed to load services', async () => {
  const store = mockDirectoryServicesStore({
    error: 'There is an error',
  })

  const { findByTestId, getByTestId } = renderDirectory(store)

  expect(await findByTestId('error-message')).toHaveTextContent(
    /^Failed to load the directory\.$/,
  )
  expect(() => getByTestId('mock-directory-services')).toThrow()
})
