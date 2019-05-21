// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import DocumentationPage from './DocumentationPage'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<DocumentationPage />, div)
  }).not.toThrow()
})
