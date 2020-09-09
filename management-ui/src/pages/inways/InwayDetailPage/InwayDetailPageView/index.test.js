// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../../../test-utils'
import InwayDetails from './index'

const inway = {
  name: 'name',
  hostname: 'host.name',
  selfAddress: 'self.address',
  services: [],
}

describe('InwayDetails', () => {
  beforeEach(() => {
    jest.useFakeTimers()
  })

  it('should display', () => {
    const { getByTestId, getByText } = renderWithProviders(
      <Router>
        <InwayDetails inway={inway} />
      </Router>,
    )

    expect(getByTestId('gateway-type')).toHaveTextContent('inway')
    expect(getByText('host.name')).toBeInTheDocument()
    expect(getByText('self.address')).toBeInTheDocument()
  })
})
