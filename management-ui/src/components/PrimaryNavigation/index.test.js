import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import PrimaryNavigation from './index'

test('PrimaryNavigation', () => {
  const { getByTestId } = renderWithProviders(
    <Router>
      <PrimaryNavigation />
    </Router>,
  )

  const linkHome = getByTestId('link-home')
  expect(linkHome.getAttribute('href')).toBe('/')

  const linkServices = getByTestId('link-services')
  expect(linkServices.getAttribute('href')).toBe('/services')
})
