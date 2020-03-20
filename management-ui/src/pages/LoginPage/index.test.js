import React from 'react'
import { renderWithProviders } from '../../test-utils'
import LoginPage from './index'

test('renders a welcome message', () => {
  const { getByText } = renderWithProviders(<LoginPage />)
  expect(getByText(/^Welcome$/)).toBeInTheDocument()
})
