// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { waitFor } from '@testing-library/react'
import { renderWithProviders } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import OrdersPage from './index'

jest.mock('../../../components/LoadingMessage', () => () => <p>loading</p>)
jest.mock('./OrdersViewPage', () => () => <p>orders view</p>)
jest.mock('./OrdersEmptyView', () => () => <p>orders empty view</p>)

test('no orders present', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListIssuedOrders = jest
    .fn()
    .mockResolvedValue({ orders: [] })

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { getByText, findByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByText('loading')).toBeInTheDocument()

  const emptyView = await findByText('orders empty view')
  expect(emptyView).toBeInTheDocument()
})

test('rendering the orders page', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListIssuedOrders = jest.fn().mockResolvedValue({
    orders: [
      {
        reference: 'ref1',
        description: 'my own description',
        delegatee: 'delegatee',
        services: [{ organization: 'organization X', service: 'service Y' }],
        validFrom: '2021-05-04',
        validUntil: '2021-05-10',
      },
    ],
  })

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { getByText, getByLabelText, findByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByText('loading')).toBeInTheDocument()

  const ordersOverview = await findByText('orders view')
  expect(ordersOverview).toBeInTheDocument()

  const linkAddOrder = getByLabelText(/Add order/)
  expect(linkAddOrder.getAttribute('href')).toBe('/orders/add-order')
})

test('failed to load orders', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListIssuedOrders = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const history = createMemoryHistory({ initialEntries: ['/orders'] })
  const store = new RootStore({
    managementApiClient,
  })

  const { getByText, findByText, queryByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <OrdersPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  await waitFor(() => {
    expect(queryByText('loading')).not.toBeInTheDocument()
  })

  expect(() => getByText('orders view')).toThrow()

  expect(await findByText(/^Failed to load orders$/)).toBeInTheDocument()
  expect(await findByText(/^arbitrary error$/)).toBeInTheDocument()
})
