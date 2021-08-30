// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import { BrowserRouter as Router } from 'react-router-dom'
import ServicesPage from './index'

test('renders without crashing', () => {
  expect(() =>
    render(
      <Router>
        <ServicesPage />
      </Router>,
    ),
  ).not.toThrow()
})
