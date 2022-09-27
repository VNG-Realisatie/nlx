// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { fireEvent, screen, waitFor } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import OutwayDetails from './index'

const outway = {
  name: 'name',
  ipAddress: '127.0.0.1',
  publicKeyPem: 'public-key-pem',
}

beforeEach(() => {
  jest.useFakeTimers()
})

test('should display outway details', async () => {
  const { getByTestId, getByText } = renderWithProviders(
    <Router>
      <OutwayDetails outway={outway} />
    </Router>,
  )

  expect(getByTestId('gateway-type')).toHaveTextContent('outway')
  expect(getByText('127.0.0.1')).toBeInTheDocument()

  fireEvent.click(screen.getByText('Certificate'))

  await waitFor(() => {
    expect(getByText('public-key-pem')).toBeInTheDocument()
  })
})
