// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import Header from './index'

test('Header with page elements', () => {
  const { getByText, getByTestId } = renderWithProviders(
    <Router>
      <Header title="Page title" description="Page description" />
    </Router>,
  )

  expect(getByText(/^Page title$/)).toBeInTheDocument()
  expect(getByText(/^Page description$/)).toBeInTheDocument()
  expect(getByTestId('user-navigation')).toBeInTheDocument()
})
