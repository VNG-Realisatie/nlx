// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../test-utils'
import Navigation from './index'

test('Settings Navigation', () => {
  const history = createMemoryHistory({ initialEntries: ['/settings'] })
  const { getByLabelText } = renderWithProviders(
    <Router history={history}>
      <Navigation />
    </Router>,
  )

  const linkHome = getByLabelText('General settings')
  expect(linkHome.getAttribute('href')).toBe('/settings/general')
})
