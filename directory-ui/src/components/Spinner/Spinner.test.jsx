import React from 'react'
import ReactDOM from 'react-dom'
import Spinner from './Spinner'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Spinner />, div)
  }).not.toThrow()
})
