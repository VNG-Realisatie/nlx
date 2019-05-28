// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import { MemoryRouter } from 'react-router-dom';
import Header from './Header'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(
      <MemoryRouter>
        <Header />
      </MemoryRouter>
      , div)
  }).not.toThrow()
})
