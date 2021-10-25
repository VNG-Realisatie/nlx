// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../../test-utils'
import OutwayDetails from './index'

const outway = {
  name: 'name',
  ipAddress: '127.0.0.1',
  publicKeyPEM: 'public-key-pem',
}

beforeEach(() => {
  jest.useFakeTimers()
})

test('should display outway details', () => {
  const { getByTestId, getByText } = renderWithProviders(
    <Router>
      <OutwayDetails outway={outway} />
    </Router>,
  )

  expect(getByTestId('gateway-type')).toHaveTextContent('outway')
  expect(getByText('127.0.0.1')).toBeInTheDocument()
  expect(getByText('public-key-pem')).toBeInTheDocument()
})
