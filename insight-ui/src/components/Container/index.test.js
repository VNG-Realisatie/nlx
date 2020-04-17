// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import ReactDOM from 'react-dom'
import Container from './index'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Container />, div)
  }).not.toThrow()
})
