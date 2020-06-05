// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../test-utils'
import ServiceDetailPage from './index'

jest.mock('../../components/ServiceDetails', () => ({ service }) => (
  <div data-testid="service-details" />
))

test('display service details', async () => {
  const getServiceByName = jest.fn().mockResolvedValue({ name: 'forty-two' })

  jest.useFakeTimers()

  const { findByTestId, getByText } = renderWithProviders(
    <Router location="/services/forty-two">
      <Route path="/services/:name">
        <ServiceDetailPage getServiceByName={getServiceByName} />
      </Route>
    </Router>,
  )
  expect(await findByTestId('service-details')).toBeInTheDocument()
  expect(getByText('forty-two')).toBeInTheDocument()
  expect(getServiceByName).toHaveBeenCalledWith('forty-two')
})

test('fetching a non-existing component', async () => {
  const getServiceByName = jest.fn().mockRejectedValue(new Error('not found'))

  const { findByTestId, getByText } = renderWithProviders(
    <Router location="/services/forty-two">
      <Route path="/services/:name">
        <ServiceDetailPage getServiceByName={getServiceByName} />
      </Route>
    </Router>,
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
    <Router location="/services/42">
      <Route path="/services/:name">
        <ServiceDetailPage getServiceByName={getServiceByName} />
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the service.')

  expect(getByText('42')).toBeInTheDocument()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
