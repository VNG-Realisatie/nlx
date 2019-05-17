// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import IconButton from './IconButton'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<IconButton />, div)
  }).not.toThrow()
})
