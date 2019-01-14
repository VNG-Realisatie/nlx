import React from 'react'
import ReactDOM from 'react-dom'
import Services from './Services'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Services serviceList={[]} />, div)
  }).not.toThrow()
})
