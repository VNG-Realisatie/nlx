// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import Header from './index'

test('renders without crashing', () => {
  expect(() =>
    render(
      <MemoryRouter>
        <Header />
      </MemoryRouter>,
    ),
  ).not.toThrow()
})
