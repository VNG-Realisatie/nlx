// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
// NOTE: about unstable_*: https://github.com/remix-run/react-router/issues/8264#issuecomment-1003781044
import {
  MemoryRouter,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
import { act, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithAllProviders, renderWithProviders } from '../test-utils'
import { UserContextProvider } from '../user-context'
import Routes from './index'

jest.mock('../pages/LoginOIDCPage', () => () => (
  <div data-testid="login-page" />
))
jest.mock('../pages/services/ServicesPage', () => () => (
  <div data-testid="services-page" />
))
jest.mock('../pages/inways-and-outways/InwaysAndOutwaysPage', () => () => (
  <div data-testid="inways-and-outways-page" />
))
jest.mock('../pages/services/AddServicePage', () => () => (
  <div data-testid="add-service-page" />
))
jest.mock('../pages/AuditLogPage', () => () => (
  <div data-testid="audit-log-page" />
))
jest.mock('../pages/TransactionLogPage', () => () => (
  <div data-testid="transaction-log-page" />
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
    <HistoryRouter history={history}>
      <UserContextProvider fetchAuthenticatedUser={fetchUser}>
        <Routes authorizationPageElement={<div>hoi</div>} />
      </UserContextProvider>
    </HistoryRouter>,
  )

  expect(history.location.pathname).toEqual('/login')
})

test('when authenticated and not accepted the ToS', async () => {
  const history = createMemoryHistory()

  await act(async () => {
    renderWithAllProviders(
      <HistoryRouter history={history}>
        <UserContextProvider user={{ id: '42' }}>
          <Routes
            tos={{ enabled: true, url: 'https://example.com', accepted: false }}
          />
        </UserContextProvider>
      </HistoryRouter>,
    )
  })

  expect(history.location.pathname).toEqual('/terms-of-service')
})

test('redirects to /inways-and-outways when navigating to /', async () => {
  const history = createMemoryHistory()

  renderWithProviders(
    <HistoryRouter history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes
          authorizationPageElement={<div>Please login</div>}
          tos={{ enabled: false }}
        />
      </UserContextProvider>
    </HistoryRouter>,
  )

  expect(history.location.pathname).toEqual('/inways-and-outways')
})

test('the /login route renders the LoginOIDCPage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/login']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes authorizationPageElement={<div>my-login-page</div>} />
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(screen.getByText('my-login-page')).toBeInTheDocument()
})

test('the /services route renders the ServicesPage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/services']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('services-page')).toBeInTheDocument()
})

test('the /services/add-service route renders the AddServicePage', () => {
  const { getByTestId } = renderWithProviders(
    <MemoryRouter initialEntries={['/services/add-service']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(getByTestId('add-service-page')).toBeInTheDocument()
})

test('the /inways-and-outways route renders the OverviewPage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/inways-and-outways']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('inways-and-outways-page')).toBeInTheDocument()
})

test('the /audit-log route renders the AuditLogPage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/audit-log']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('audit-log-page')).toBeInTheDocument()
})

test('the /transaction-log route renders the TransactionLogPage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/transaction-log']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('transaction-log-page')).toBeInTheDocument()
})

test('the /finances route renders the FinancePage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/finances']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('finances-page')).toBeInTheDocument()
})

test('the /orders route redirects to /orders/outgoing', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/orders']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('orders-page')).toBeInTheDocument()
})

test('the /orders/outgoing route renders the OrdersPage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/orders/outgoing']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('orders-page')).toBeInTheDocument()
})

test('the /orders/incoming route renders the OrdersPage', () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/orders/incoming']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(screen.getByTestId('orders-page')).toBeInTheDocument()
})

test('the /orders/add route renders the AddOrderPage', () => {
  const { getByTestId } = renderWithProviders(
    <MemoryRouter initialEntries={['/orders/add-order']}>
      <UserContextProvider user={{ id: '42' }}>
        <Routes tos={{ enabled: false }} />
      </UserContextProvider>
    </MemoryRouter>,
  )
  expect(getByTestId('add-order-page')).toBeInTheDocument()
})
