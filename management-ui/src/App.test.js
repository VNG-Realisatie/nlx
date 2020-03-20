import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { render } from '@testing-library/react'
import App from './App'

jest.mock('./pages/LoginPage', () => () => <div data-testid="login-page" />)
jest.mock('./pages/ServicesPage', () => () => (
  <div data-testid="services-page" />
))

test('redirects to /login when navigating to /', async () => {
  const history = createMemoryHistory()
  render(
    <Router history={history}>
      <App />
    </Router>,
  )
  expect(history.location.pathname).toEqual('/login')
})

test('the /login route renders the LoginPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/login'] })
  const { getByTestId } = render(
    <Router history={history}>
      <App />
    </Router>,
  )
  expect(getByTestId('login-page')).toBeInTheDocument()
})

test('the /services route renders the ServicesPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/services'] })
  const { getByTestId } = render(
    <Router history={history}>
      <App />
    </Router>,
  )
  expect(getByTestId('services-page')).toBeInTheDocument()
})
