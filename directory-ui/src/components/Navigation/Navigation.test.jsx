import React from 'react'
import ReactDOM from 'react-dom'
import { MemoryRouter } from 'react-router-dom';
import Navigation from './Navigation'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(
      <MemoryRouter>
        <Navigation />
      </MemoryRouter>
      , div)
  }).not.toThrow()
})
