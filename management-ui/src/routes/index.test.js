// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'

import { renderWithProviders } from '../test-utils'
import { UserContextProvider } from '../user-context'
import Routes from '.'

jest.mock('../pages/LoginPage', () => () => <div data-testid="login-page" />)
jest.mock('../pages/services/ServicesPage', () => () => (
  <div data-testid="services-page" />
))
jest.mock('../pages/inways/InwaysPage', () => () => <div data-testid="inways-page" />)
jest.mock('../pages/services/AddServicePage', () => () => (
  <div data-testid="add-service-page" />
))

test('when not authenticated it redirects to /login when navigating to /', async () => {
  const history = createMemoryHistory()
  const fetchUser = () => {
    throw new Error('not authenticated')
  }
  renderWithProviders(
    <Router history={history}>
      <UserContextProvider fetchAuthenticatedUser={fetchUser}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(history.location.pathname).toEqual('/login')
})

test('redirects to /inways when navigating to /', async () => {
  const history = createMemoryHistory()
  renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(history.location.pathname).toEqual('/inways')
})

test('the /login route renders the LoginPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/login'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('login-page')).toBeInTheDocument()
})

test('the /services route renders the ServicesPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/services'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('services-page')).toBeInTheDocument()
})

test('the /services/add-service route renders the AddServicePage', () => {
  const history = createMemoryHistory({
    initialEntries: ['/services/add-service'],
  })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('add-service-page')).toBeInTheDocument()
})

test('the /inways route renders the InwaysPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/inways'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('inways-page')).toBeInTheDocument()
})
