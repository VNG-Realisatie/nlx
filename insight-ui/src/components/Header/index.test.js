import React from 'react'
import ReactDOM from 'react-dom'
import Header from './index'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Header />, div)
  }).not.toThrow()
})
