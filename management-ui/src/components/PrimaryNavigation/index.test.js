// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import PrimaryNavigation from './index'

test('PrimaryNavigation', () => {
  const { getByLabelText } = renderWithProviders(
    <Router>
      <PrimaryNavigation />
    </Router>,
  )

  const linkHome = getByLabelText('Homepage')
  expect(linkHome.getAttribute('href')).toBe('/')

  const linkInways = getByLabelText('Inways page')
  expect(linkInways.getAttribute('href')).toBe('/inways')

  const linkServices = getByLabelText('Services page')
  expect(linkServices.getAttribute('href')).toBe('/services')

  const linkSettings = getByLabelText('Settings page')
  expect(linkSettings.getAttribute('href')).toBe('/settings')
})
