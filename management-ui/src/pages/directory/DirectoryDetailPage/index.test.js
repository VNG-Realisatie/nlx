// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import DirectoryDetailPage from './index'

jest.mock('./components/DirectoryDetailView', () => ({ service }) => (
  <div data-testid="directory-service-details" />
))

test('display directory service details', async () => {
  const getService = jest.fn().mockResolvedValue({
    organizationName: 'organization',
    serviceName: 'service',
    state: 'up',
  })

  jest.useFakeTimers()

  const { findByTestId, getByText } = renderWithProviders(
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage getService={getService} />
      </Route>
    </Router>,
  )

  expect(await findByTestId('directory-service-details')).toBeInTheDocument()
  expect(getByText('organization')).toBeInTheDocument()
  expect(getService).toHaveBeenCalledWith('organization', 'service')
})

test('fetching a non-existing component', async () => {
  const getService = jest
    .fn()
    .mockRejectedValue(new Error('invalid user input'))

  const { findByTestId, getByText, queryByText } = renderWithProviders(
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage getService={getService} />
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeInTheDocument()
  expect(message.textContent).toBe('Failed to load the service.')

  expect(getByText('service')).toBeInTheDocument()
  expect(queryByText('organization')).toBeNull()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeInTheDocument()
})

test('fetching service details fails for an unknown reason', async () => {
  const getService = jest.fn().mockRejectedValue(new Error('arbitrary reason'))

  const { findByTestId, getByText, queryByText } = renderWithProviders(
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage getService={getService} />
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeInTheDocument()
  expect(message.textContent).toBe('Failed to load the service.')

  expect(getByText('service')).toBeInTheDocument()
  expect(queryByText('organization')).toBeNull()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeInTheDocument()
})
