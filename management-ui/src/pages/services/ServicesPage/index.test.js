// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter, Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders, waitFor } from '../../../test-utils'
import { INTERVAL } from '../../../hooks/use-polling'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import ServicesPage from './index'

jest.mock('../../../components/PageTemplate/OrganizationInwayCheck', () => () =>
  null,
)
jest.mock('./ServicesPageView', () => () => (
  <p data-testid="services-list">mock-services</p>
))
jest.mock('../../../components/OrganizationName', () => () => <span>test</span>)

test('fetching all services', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [{ name: 'my-first-service' }],
  })

  const history = createMemoryHistory({ initialEntries: ['/services'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { getByRole, getByTestId, getByLabelText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <ServicesPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('services-list')).toThrow()

  await waitFor(() =>
    expect(getByTestId('services-list')).toHaveTextContent('mock-services'),
  )
  expect(getByTestId('service-count')).toHaveTextContent('1Services')

  const linkAddService = getByLabelText(/Add service/)
  expect(linkAddService.getAttribute('href')).toBe('/services/add-service')
})

test('failed to load services', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest
    .fn()
    .mockRejectedValue('arbitrary error')

  const store = new RootStore({
    managementApiClient,
  })

  const { findByText, getByTestId } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <ServicesPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(() => getByTestId('services-list')).toThrow()
  expect(await findByText(/^Failed to load the services$/)).toBeInTheDocument()
  expect(getByTestId('service-count')).toHaveTextContent('0Services')
})

test('service statistics should be polled', () => {
  jest.useFakeTimers()
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [{ name: 'my-first-service' }],
  })

  managementApiClient.managementGetStatisticsOfServices = jest
    .fn()
    .mockResolvedValue({
      services: [],
    })

  const history = createMemoryHistory({ initialEntries: ['/services'] })
  const rootStore = new RootStore({
    managementApiClient,
  })

  renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <ServicesPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  jest.advanceTimersByTime(INTERVAL * 2)

  expect(
    managementApiClient.managementGetStatisticsOfServices,
  ).toHaveBeenCalledTimes(2)

  jest.useRealTimers()
})
