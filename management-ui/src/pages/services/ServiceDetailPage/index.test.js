// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter, Router } from 'react-router-dom'

import { act } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { observable } from 'mobx'
import { renderWithProviders } from '../../../test-utils'
import { StoreProvider } from '../../../stores'
import ServiceDetailPage from './index'

jest.mock('./ServiceDetailView', () => ({ removeHandler }) => (
  <div data-testid="service-details">
    <button type="button" onClick={removeHandler}>
      Remove service
    </button>
  </div>
))
const storeTemplate = ({
  fetchServices = jest.fn(),
  removeService = jest.fn(),
  selectService = jest.fn(),
}) =>
  observable({
    servicesStore: {
      services: [{ name: 'forty-two' }],
      isReady: true,
      error: '',
      fetchServices,
      selectService,
      removeService,
      addService: jest.fn(),
    },
  })

test('display service details', async () => {
  const selectService = jest
    .fn()
    .mockReturnValue({ name: 'forty-two', fetch: jest.fn() })
  const store = storeTemplate({ selectService })
  const { findByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/services/forty-two">
      <Route path="/services/:name">
        <StoreProvider store={store}>
          <ServiceDetailPage />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )

  jest.useFakeTimers()

  expect(await findByTestId('service-details')).toBeInTheDocument()
  expect(getByText('forty-two')).toBeInTheDocument()
  expect(selectService).toHaveBeenCalledWith('forty-two')
})

test('fetching a non-existing component', async () => {
  const selectService = jest.fn()
  const store = storeTemplate({ selectService })

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
  const fetchServices = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary reason'))

  const store = storeTemplate({ fetchServices })

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
  const selectService = jest
    .fn()
    .mockReturnValue({ name: 'dummy-service', fetch: jest.fn() })
  const removeService = jest.fn()
  const store = storeTemplate({ selectService, removeService })

  const { findByText } = renderWithProviders(
    <Router history={history}>
      <Route path="/services/:name">
        <StoreProvider store={store}>
          <ServiceDetailPage />
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
