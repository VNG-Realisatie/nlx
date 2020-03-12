import React from 'react'
import { render } from '../../../test-utils' // TODO: root alias, so we don't need relative paths
import LoginPage from './index'

test('renders a welcome message', () => {
  const { getByText } = render(<LoginPage />)
  expect(getByText(/^Welkom$/)).toBeInTheDocument()
})
