// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../test-utils'
import PageTemplate from './index'

test('PageTemplate with page elements', () => {
  const { getByText, getByTestId } = renderWithProviders(
    <Router>
      <PageTemplate title="Page title" description="Page description">
        <p>Page content</p>
      </PageTemplate>
    </Router>,
  )

  expect(getByText(/^Page title$/)).toBeInTheDocument()
  expect(getByText(/^Page description$/)).toBeInTheDocument()
  expect(getByText(/^Page content$/)).toBeInTheDocument()
  expect(getByTestId('user-navigation')).toBeInTheDocument()
})
