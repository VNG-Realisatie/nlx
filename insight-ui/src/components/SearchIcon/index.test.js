import React from 'react'
import ReactDOM from 'react-dom'
import SearchIcon from './index'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<SearchIcon />, div)
  }).not.toThrow()
})
