import React from 'react'
import ReactDOM from 'react-dom'
import Item from './index'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Item />, div)
  }).not.toThrow()
})
