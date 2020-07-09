// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter, Router } from 'react-router-dom'

import { act } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../test-utils'
import ServiceDetailPage from './index'

jest.mock('../../components/ServiceDetails', () => ({ removeHandler }) => (
  <div data-testid="service-details">
    <button type="button" onClick={removeHandler}>
      Remove service
    </button>
  </div>
))

test('display service details', async () => {
  const getServiceByName = jest.fn().mockResolvedValue({ name: 'forty-two' })

  jest.useFakeTimers()

  const { findByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/services/forty-two">
      <Route path="/services/:name">
        <ServiceDetailPage getServiceByName={getServiceByName} />
      </Route>
    </StaticRouter>,
  )
  expect(await findByTestId('service-details')).toBeInTheDocument()
  expect(getByText('forty-two')).toBeInTheDocument()
  expect(getServiceByName).toHaveBeenCalledWith('forty-two')
})

test('fetching a non-existing component', async () => {
  const getServiceByName = jest.fn().mockRejectedValue(new Error('not found'))

  const { findByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/services/forty-two">
      <Route path="/services/:name">
        <ServiceDetailPage getServiceByName={getServiceByName} />
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
  const getServiceByName = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary reason'))

  const { findByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/services/42">
      <Route path="/services/:name">
        <ServiceDetailPage getServiceByName={getServiceByName} />
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
  const history = createMemoryHistory()
  const getServiceByName = jest
    .fn()
    .mockResolvedValue({ name: 'dummy-service' })
  const refreshHandler = jest.fn()
  const removeService = jest.fn().mockResolvedValue()

  const { findByText } = renderWithProviders(
    <Router history={history}>
      <ServiceDetailPage
        getServiceByName={getServiceByName}
        refreshHandler={refreshHandler}
        removeService={removeService}
      />
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
  expect(refreshHandler).toHaveBeenCalledTimes(1)
})
