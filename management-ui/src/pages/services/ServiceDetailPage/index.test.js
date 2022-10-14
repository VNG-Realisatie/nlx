// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import {
  Route,
  Routes,
  MemoryRouter,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
import { act, fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementServiceApi } from '../../../api'
import ServiceDetailPage from './index'

// eslint-disable-next-line react/prop-types
jest.mock('./ServiceDetailView', () => ({ removeHandler }) => (
  <div data-testid="service-details">
    <button type="button" onClick={removeHandler}>
      Remove service
    </button>
  </div>
))

test('display service details', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceGetService = jest
    .fn()
    .mockResolvedValue({
      name: 'my-service',
    })

  managementApiClient.managementServiceListIncomingAccessRequests = jest
    .fn()
    .mockResolvedValue({ accessRequests: [] })

  managementApiClient.managementServiceListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({ accessGrants: [] })

  const rootStore = new RootStore({ managementApiClient })

  renderWithAllProviders(
    <MemoryRouter initialEntries={['/my-service']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":name" element={<ServiceDetailPage />} />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(await screen.findByTestId('service-details')).toBeInTheDocument()
  expect(screen.getByText('my-service')).toBeInTheDocument()
})

test('fetching a non-existing service', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceGetService = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const rootStore = new RootStore({ managementApiClient })

  renderWithAllProviders(
    <MemoryRouter initialEntries={['/my-service']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":name" element={<ServiceDetailPage />} />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )

  const message = await screen.findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the service')

  expect(screen.getByText('my-service')).toBeInTheDocument()

  const closeButton = await screen.findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})

test('removing the service', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceGetService = jest
    .fn()
    .mockResolvedValue({
      name: 'my-service',
    })

  managementApiClient.managementServiceListIncomingAccessRequests = jest
    .fn()
    .mockResolvedValue({ accessRequests: [] })

  managementApiClient.managementServiceListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({ accessGrants: [] })

  managementApiClient.managementServiceDeleteService = jest
    .fn()
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  const rootStore = new RootStore({
    managementApiClient,
  })
  jest.spyOn(rootStore.servicesStore, 'removeService')

  const history = createMemoryHistory({
    initialEntries: ['/my-service'],
  })

  renderWithAllProviders(
    <HistoryRouter history={history}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":name" element={<ServiceDetailPage />} />
        </Routes>
      </StoreProvider>
    </HistoryRouter>,
  )

  fireEvent.click(await screen.findByText('Remove service'))

  expect(rootStore.servicesStore.removeService).toHaveBeenCalledTimes(1)
  expect(await screen.findByRole('alert')).toHaveTextContent(
    "Failed to remove the serviceYou don't have the required permission.",
  )

  await act(async () => {
    fireEvent.click(await screen.findByText('Remove service'))
  })

  expect(rootStore.servicesStore.removeService).toHaveBeenCalledTimes(2)
  expect(history.location.pathname).toEqual('/my-service')
  expect(history.location.search).toEqual('?lastAction=removed')
})
