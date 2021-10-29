// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'

import { renderWithProviders } from '../test-utils'
import { UserContextProvider } from '../user-context'
import Routes from '.'

jest.mock('../pages/LoginOIDCPage', () => () => (
  <div data-testid="login-page" />
))
jest.mock('../pages/services/ServicesPage', () => () => (
  <div data-testid="services-page" />
))
jest.mock('../pages/inways/OverviewPage', () => () => (
  <div data-testid="overview-page" />
))
jest.mock('../pages/services/AddServicePage', () => () => (
  <div data-testid="add-service-page" />
))
jest.mock('../pages/AuditLogPage', () => () => (
  <div data-testid="audit-log-page" />
))
jest.mock('../pages/FinancePage', () => () => (
  <div data-testid="finances-page" />
))
jest.mock('../pages/orders/OrdersPage', () => () => (
  <div data-testid="orders-page" />
))
jest.mock('../pages/orders/AddOrderPage', () => () => (
  <div data-testid="add-order-page" />
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

test('the /login route renders the LoginOIDCPage', () => {
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

test('the /inways-and-outways route renders the OverviewPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/inways'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('overview-page')).toBeInTheDocument()
})

test('the /audit-log route renders the AuditLogPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/audit-log'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('audit-log-page')).toBeInTheDocument()
})

test('the /finances route renders the FinancePage', () => {
  const history = createMemoryHistory({ initialEntries: ['/finances'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('finances-page')).toBeInTheDocument()
})

test('the /orders route redirects to /orders/outgoing', () => {
  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('orders-page')).toBeInTheDocument()
})

test('the /orders/outgoing route renders the OrdersPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/orders/outgoing'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('orders-page')).toBeInTheDocument()
})

test('the /orders/incoming route renders the OrdersPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/orders/incoming'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('orders-page')).toBeInTheDocument()
})

test('the /orders/add route renders the AddOrderPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/orders/add-order'] })
  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('add-order-page')).toBeInTheDocument()
})
