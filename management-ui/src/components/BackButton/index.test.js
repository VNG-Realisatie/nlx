// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { StaticRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import BackButton from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <Router>
        {' '}
        <BackButton to="/link" />
      </Router>,
    ),
  ).not.toThrow()
})
