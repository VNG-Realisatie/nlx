// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import { waitFor, fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithAllProviders } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementServiceApi } from '../../../api'
import OrdersPage from './index'

const createOrdersPage = (managementApiClient) => {
  const store = new RootStore({
    managementApiClient,
  })

  const history = createMemoryHistory()

  renderWithAllProviders(
    <HistoryRouter history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </HistoryRouter>,
  )

  return { history }
}

test('rendering the orders page', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          reference: 'ref1',
          description: 'description of the first outgoing order',
          delegatee: '01234567890123456789',
          validFrom: '2021-05-04',
          validUntil: '2021-05-10',
        },
      ],
    })

  managementApiClient.managementServiceListIncomingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  createOrdersPage(managementApiClient)

  expect(screen.getByRole('progressbar')).toBeInTheDocument()

  const ordersOverview = await screen.findByText(
    'description of the first outgoing order',
  )
  expect(ordersOverview).toBeInTheDocument()

  const linkAddOrder = screen.getByLabelText(/Add order/)
  expect(linkAddOrder.getAttribute('href')).toBe('/add-order')
})

test('no outgoing orders present', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  managementApiClient.managementServiceListIncomingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  createOrdersPage(managementApiClient)

  expect(screen.getByRole('progressbar')).toBeInTheDocument()

  const emptyView = await screen.findByText(
    "You don't have any issued orders yet",
  )
  expect(emptyView).toBeInTheDocument()
})

test('failed to load outgoing orders', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  createOrdersPage(managementApiClient)

  await waitFor(() => {
    expect(screen.queryByRole('progressbar')).not.toBeInTheDocument()
  })

  expect(await screen.findByText(/^Failed to load orders$/)).toBeInTheDocument()
  expect(await screen.findByText(/^arbitrary error$/)).toBeInTheDocument()
})

test('no incoming orders present', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  managementApiClient.managementServiceListIncomingOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  const { history } = createOrdersPage(managementApiClient)

  expect(screen.getByRole('progressbar')).toBeInTheDocument()

  const incomingOrdersButton = await screen.findByText(/Received/)

  fireEvent.click(incomingOrdersButton)

  const emptyView = await screen.findByText(
    "You haven't received any orders yet",
  )

  expect(history.location.pathname).toEqual('/incoming')

  expect(emptyView).toBeInTheDocument()
})

test('failed to load incoming orders', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceListIncomingOrders = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  createOrdersPage(managementApiClient)

  await waitFor(() => {
    expect(screen.queryByRole('progressbar')).not.toBeInTheDocument()
  })

  expect(await screen.findByText(/^Failed to load orders$/)).toBeInTheDocument()
  expect(await screen.findByText(/^arbitrary error$/)).toBeInTheDocument()
})
