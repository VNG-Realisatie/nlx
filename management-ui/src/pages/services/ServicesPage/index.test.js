// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter, Route, Routes } from 'react-router-dom'
import { renderWithAllProviders, waitFor } from '../../../test-utils'
import { INTERVAL } from '../../../hooks/use-polling'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementServiceApi } from '../../../api'
import ServicesPage from './index'

jest.mock(
  '../../../components/PageTemplate/OrganizationInwayCheck',
  () => () => null,
)
jest.mock(
  '../../../components/PageTemplate/OrganizationEmailAddressCheck',
  () => () => null,
)
jest.mock('./ServicesPageView', () => () => (
  <p data-testid="services-list">mock-services</p>
))
jest.mock('../../../components/OrganizationName', () => () => <span>test</span>)

test('fetching all services', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [{ name: 'my-first-service' }],
    })

  const store = new RootStore({
    managementApiClient,
  })

  const { getByRole, getByTestId, getByLabelText } = renderWithAllProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <Routes>
            <Route path="*" element={<ServicesPage />} />
          </Routes>
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
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
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockRejectedValue('arbitrary error')

  const store = new RootStore({
    managementApiClient,
  })

  const { findByText, getByTestId } = renderWithAllProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <Routes>
            <Route path="*" element={<ServicesPage />} />
          </Routes>
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
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [{ name: 'my-first-service' }],
    })

  managementApiClient.managementServiceGetStatisticsOfServices = jest
    .fn()
    .mockResolvedValue({
      services: [],
    })

  const rootStore = new RootStore({
    managementApiClient,
  })

  renderWithAllProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <Routes>
            <Route path="*" element={<ServicesPage />} />
          </Routes>
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  jest.advanceTimersByTime(INTERVAL * 2)

  expect(
    managementApiClient.managementServiceGetStatisticsOfServices,
  ).toHaveBeenCalledTimes(2)

  jest.useRealTimers()
})
