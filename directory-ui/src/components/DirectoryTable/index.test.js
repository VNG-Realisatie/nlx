// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import { SERVICE_STATE_UP } from '../../components/StateIndicator'
import DirectoryServices from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <Router>
        <DirectoryServices services={[]} />
      </Router>,
    ),
  ).not.toThrow()
})

test('show a empty services message', () => {
  const { getByTestId } = renderWithProviders(
    <Router>
      <DirectoryServices services={[]} />
    </Router>,
  )
  expect(getByTestId('directory-no-services')).toHaveTextContent(
    'Geen services gevonden',
  )
})

test('show a table with rows for every service', () => {
  const { getByTestId, getByRole } = renderWithProviders(
    <Router>
      <DirectoryServices
        services={[
          {
            organization: 'Test Organization',
            name: 'Test Service',
            status: SERVICE_STATE_UP,
          },
        ]}
      />
    </Router>,
  )

  expect(getByRole('grid')).toBeTruthy()
  expect(getByTestId('directory-service-row')).toBeInTheDocument()
})
