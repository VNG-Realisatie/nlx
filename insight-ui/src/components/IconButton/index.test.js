import React from 'react'
import ReactDOM from 'react-dom'
import IconButton from './index'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<IconButton />, div)
  }).not.toThrow()
})
