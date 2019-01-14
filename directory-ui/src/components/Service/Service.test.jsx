import React from 'react'
import ReactDOM from 'react-dom'
import Service from './Service'

it('renders without crashing', () => {
  expect(() => {
    const tbody = document.createElement('tbody')
    ReactDOM.render(<Service data={{
      organization_name: 'Organization Name',
      service_name: 'Service Name',
    }} />, tbody)
  }).not.toThrow()
})
