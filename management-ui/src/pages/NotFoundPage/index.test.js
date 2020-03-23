import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../test-utils'
import NotFoundPage from './index'

test('renders a welcome message', () => {
  const { getByText } = renderWithProviders(
    <Router>
      <NotFoundPage />
    </Router>,
  )

  expect(getByText(/^Page not found$/)).toBeInTheDocument()
})
