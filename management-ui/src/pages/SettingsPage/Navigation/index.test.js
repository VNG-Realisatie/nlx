// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../test-utils'
import Navigation from './index'

test('Settings Navigation', () => {
  renderWithProviders(
    <MemoryRouter>
      <Navigation />
    </MemoryRouter>,
  )

  const linkHome = screen.getByLabelText('General settings')
  expect(linkHome.getAttribute('href')).toBe('/general')
})
