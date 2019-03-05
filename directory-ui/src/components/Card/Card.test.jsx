import React from 'react'
import ReactDOM from 'react-dom'
import Card from './Card'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Card />, div)
  }).not.toThrow()
})
