// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import BackButton from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <MemoryRouter>
        <BackButton to="/link" />
      </MemoryRouter>,
    ),
  ).not.toThrow()
})
