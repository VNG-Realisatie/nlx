// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { waitFor, fireEvent } from '@testing-library/react'
import { renderWithProviders } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import OrdersPage from './index'

test('rendering the orders page', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          reference: 'ref1',
          description: 'description of the first outgoing order',
          delegatee: 'delegatee',
          services: [{ organization: 'organization X', service: 'service Y' }],
          validFrom: '2021-05-04',
          validUntil: '2021-05-10',
        },
      ],
    })

  managementApiClient.managementListIncomingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { getByRole, getByLabelText, findByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()

  const ordersOverview = await findByText(
    'description of the first outgoing order',
  )
  expect(ordersOverview).toBeInTheDocument()

  const linkAddOrder = getByLabelText(/Add order/)
  expect(linkAddOrder.getAttribute('href')).toBe('/orders/add-order')
})

test('no outgoing orders present', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  managementApiClient.managementListIncomingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { getByRole, findByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()

  const emptyView = await findByText("You don't have any issued orders yet")
  expect(emptyView).toBeInTheDocument()
})

test('failed to load outgoing orders', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { findByText, queryByRole } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  await waitFor(() => {
    expect(queryByRole('progressbar')).not.toBeInTheDocument()
  })

  expect(await findByText(/^Failed to load orders$/)).toBeInTheDocument()
  expect(await findByText(/^arbitrary error$/)).toBeInTheDocument()
})

test('no incoming orders present', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  managementApiClient.managementListIncomingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { getByRole, findByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()

  const incomingOrdersButton = await findByText(/Received/)

  fireEvent.click(incomingOrdersButton)

  const emptyView = await findByText("You haven't received any orders yet")
  expect(emptyView).toBeInTheDocument()
})

test('failed to load incoming orders', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListIncomingOrders = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { findByText, queryByRole } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  await waitFor(() => {
    expect(queryByRole('progressbar')).not.toBeInTheDocument()
  })

  expect(await findByText(/^Failed to load orders$/)).toBeInTheDocument()
  expect(await findByText(/^arbitrary error$/)).toBeInTheDocument()
})
