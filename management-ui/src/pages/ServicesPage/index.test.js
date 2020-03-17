import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import ServicesPage from './index'

test('renders the page content', () => {
  const { getByText } = renderWithProviders(
    <Router>
      <ServicesPage />
    </Router>,
  )
  expect(getByText(/^An overview of the services here\.$/)).toBeInTheDocument()
})
