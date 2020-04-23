// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import HeaderWithBackNavigation from './index'

test('Header with page elements', () => {
  const { getByText, queryByTestId, getByLabelText } = renderWithProviders(
    <Router>
      <HeaderWithBackNavigation title="Page title" backButtonTo="/link" />
    </Router>,
  )

  expect(getByText(/^Page title$/)).toBeInTheDocument()
  expect(getByLabelText('Back')).toBeInTheDocument()
  expect(queryByTestId('user-navigation')).not.toBeInTheDocument()
})
