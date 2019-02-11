import React from 'react'
import ReactDOM from 'react-dom'
import ServicesOverviewPage from './ServicesOverviewPage'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<ServicesOverviewPage />, div)
  }).not.toThrow()
})
