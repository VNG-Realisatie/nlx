// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { fireEvent } from '@testing-library/react'

import { renderWithProviders } from '../../../../test-utils'
import InwayDetails from './index'

const inway = {
  name: 'name',
  ipAddress: '127.0.0.1',
  hostname: 'host.name',
  selfAddress: 'self.address',
  services: [],
}

beforeEach(() => {
  jest.useFakeTimers()
})

test('should display inway details', () => {
  const { getByTestId, getByText } = renderWithProviders(
    <Router>
      <InwayDetails inway={inway} />
    </Router>,
  )

  expect(getByTestId('gateway-type')).toHaveTextContent('inway')
  expect(getByText('127.0.0.1')).toBeInTheDocument()
  expect(getByText('host.name')).toBeInTheDocument()
  expect(getByText('self.address')).toBeInTheDocument()

  fireEvent.click(getByTestId('inway-services'))

  expect(getByText('No services have been connected')).toBeInTheDocument()
})

test('should render list of connected services', () => {
  const inwayWithServices = {
    ...inway,
    services: [{ name: 'service' }],
  }

  const { getByTestId, getByText, queryByText } = renderWithProviders(
    <Router>
      <InwayDetails inway={inwayWithServices} />
    </Router>,
  )

  fireEvent.click(getByTestId('inway-services'))

  expect(getByText('service')).toBeInTheDocument()
  expect(queryByText('No services have been connected')).not.toBeInTheDocument()
})
