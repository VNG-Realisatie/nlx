// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import ReactDOM from 'react-dom'
import Navigation from './index'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Navigation />, div)
  }).not.toThrow()
})
