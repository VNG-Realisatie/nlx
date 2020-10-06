// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import ReactDOM from 'react-dom'
import Switch from './Switch'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Switch />, div)
  }).not.toThrow()
})
