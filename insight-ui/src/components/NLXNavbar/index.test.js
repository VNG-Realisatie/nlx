import React from 'react'
import ReactDOM from 'react-dom'
import { MemoryRouter } from 'react-router-dom'
import NLXNavbar from './index'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(
      <MemoryRouter>
        <NLXNavbar
          homePageURL="https://www.nlx.io"
          aboutPageURL="https://nlx.io/about/"
          docsPageURL="https://docs.nlx.io/"
        />
      </MemoryRouter>,
      div,
    )
  }).not.toThrow()
})
