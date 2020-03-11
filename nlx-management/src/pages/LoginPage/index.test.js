import React from 'react'
import { render } from '@testing-library/react'
import LoginPage from './index'

test('renders a welcome message', () => {
  const { getByText } = render(<LoginPage />)
  expect(getByText('Welkom')).toBeInTheDocument()
})
