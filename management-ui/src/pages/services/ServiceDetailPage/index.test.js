// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Router, StaticRouter } from 'react-router-dom'

import { act } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../test-utils'
import { StoreProvider } from '../../../stores'
import { mockServicesStore } from '../../../stores/ServicesStore.mock'
import ServiceDetailPage from './index'

// eslint-disable-next-line react/prop-types
jest.mock('./ServiceDetailView', () => ({ removeHandler }) => (
  <div data-testid="service-details">
    <button type="button" onClick={removeHandler}>
      Remove service
    </button>
  </div>
))

let fetchIncomingAccessRequests
let fetchAccessGrants

beforeEach(() => {
  fetchIncomingAccessRequests = jest.fn()
  fetchAccessGrants = jest.fn()
})

test('display service details', () => {
  const store = mockServicesStore({})
  const { getByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/services/forty-two">
      <Route path="/services/:name">
        <StoreProvider store={store}>
          <ServiceDetailPage
            service={{
              name: 'forty-two',
              fetchIncomingAccessRequests,
              fetchAccessGrants,
            }}
          />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )

  expect(getByTestId('service-details')).toBeInTheDocument()
  expect(getByText('forty-two')).toBeInTheDocument()
})

test('fetching a non-existing component', async () => {
  const getService = jest.fn()
  const store = mockServicesStore({ getService })

  const { findByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/services/forty-two">
      <Route path="/services/:name">
        <StoreProvider store={store}>
          <ServiceDetailPage />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )
  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the service.')

  expect(getByText('forty-two')).toBeInTheDocument()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})

test('fetching service details fails for an unknown reason', async () => {
  const store = mockServicesStore({ error: 'arbitrary reason' })

  const { findByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/services/42">
      <Route path="/services/:name">
        <StoreProvider store={store}>
          <ServiceDetailPage />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the service.')

  expect(getByText('42')).toBeInTheDocument()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})

test('removing the service', async () => {
  const history = createMemoryHistory({
    initialEntries: ['/services/dummy-service'],
  })
  const removeService = jest.fn()
  const store = mockServicesStore({ removeService })

  const { findByText } = renderWithProviders(
    <Router history={history}>
      <Route path="/services/:name">
        <StoreProvider store={store}>
          <ServiceDetailPage
            service={{
              name: 'dummy-service',
              fetchIncomingAccessRequests,
              fetchAccessGrants,
            }}
          />
        </StoreProvider>
      </Route>
    </Router>,
  )

  const removeButton = await findByText('Remove service')
  act(() => {
    removeButton.click()
  })

  expect(removeService).toHaveBeenCalledTimes(1)
  await act(async () => {})
  expect(history.location.pathname).toEqual('/services/dummy-service')
  expect(history.location.search).toEqual('?lastAction=removed')
})
