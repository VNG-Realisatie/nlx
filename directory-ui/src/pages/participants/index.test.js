// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import { BrowserRouter as Router } from 'react-router-dom'
import ParticipantsPage from './index'

test('renders without crashing', () => {
  expect(() =>
    render(
      <Router>
        <ParticipantsPage />
      </Router>,
    ),
  ).not.toThrow()
})
