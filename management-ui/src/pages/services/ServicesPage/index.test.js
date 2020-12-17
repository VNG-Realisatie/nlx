// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter, Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { act } from '@testing-library/react'
import { renderWithProviders, waitFor } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { StoreProvider } from '../../../stores'
import { mockServicesStore } from '../../../stores/ServicesStore.mock'
import ServicesPage from './index'

jest.mock('../../../components/PageTemplate/OrganizationInwayCheck', () => () =>
  null,
)
jest.mock('./ServicesPageView', () => () => (
  <p data-testid="services-list">mock-services</p>
))

test('fetching all services', async () => {
  const history = createMemoryHistory({ initialEntries: ['/services'] })

  const store = mockServicesStore({ services: null, isInitiallyFetched: false })
  const { getByRole, getByTestId, getByLabelText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider store={store}>
          <ServicesPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('services-list')).toThrow()

  await act(async () => {
    store.servicesStore.services = [
      {
        name: 'my-first-service',
        inways: [],
        internal: false,
      },
    ]
    store.servicesStore.isReady = true
  })

  waitFor(() =>
    expect(getByTestId('services-list')).toHaveTextContent('mock-services'),
  )
  expect(getByTestId('service-count')).toHaveTextContent('1Services')

  const linkAddService = getByLabelText(/Add service/)
  expect(linkAddService.getAttribute('href')).toBe('/services/add-service')
})

test('failed to load services', async () => {
  const store = mockServicesStore({ services: null, error: 'arbitrary error' })

  const { findByText, getByTestId } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider store={store}>
          <ServicesPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(() => getByTestId('services-list')).toThrow()
  expect(await findByText(/^Failed to load the services$/)).toBeInTheDocument()
  expect(getByTestId('service-count')).toHaveTextContent('0Services')
})
