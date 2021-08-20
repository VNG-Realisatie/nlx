// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Router, StaticRouter } from 'react-router-dom'
import { fireEvent, waitFor, within } from '@testing-library/react'
import { configure } from 'mobx'
import { createMemoryHistory } from 'history'
import { renderWithAllProviders, screen } from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { ManagementApi } from '../../../../../api'
import OrderDetailPage from './index'

jest.mock('../../../../../components/Modal')

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  jest.useRealTimers()
})

test('display order details', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegatee: 'delegatee',
          reference: 'my-reference',
          description: 'description',
          validFrom: '2020-01-01',
          validUntil: '3000-01-01',
          revokedAt: null,
          services: [],
        },
      ],
    })
  const rootStore = new RootStore({ managementApiClient })
  const orderStore = rootStore.orderStore

  await orderStore.fetchOutgoing()

  const history = createMemoryHistory({
    initialEntries: ['/delegatee/reference'],
  })

  renderWithAllProviders(
    <Router history={history}>
      <Route path="/:delegatee/:reference">
        <StoreProvider rootStore={rootStore}>
          <OrderDetailPage order={orderStore.outgoingOrders[0]} />
        </StoreProvider>
      </Route>
    </Router>,
  )

  expect(screen.getByText('Issued to delegatee')).toBeInTheDocument()
  expect(screen.getByText('description')).toBeInTheDocument()

  const orderModel = orderStore.outgoingOrders[0]
  jest.spyOn(orderModel, 'revoke')

  const revokeButton = await screen.findByText('Revoke')
  fireEvent.click(revokeButton)

  const confirmModal = screen.getByRole('dialog')
  const okButton = within(confirmModal).getByText('Revoke')

  managementApiClient.managementRevokeOutgoingOrder = jest
    .fn()
    .mockResolvedValue()

  fireEvent.click(okButton)
  await waitFor(() => expect(orderModel.revoke).toHaveBeenCalledTimes(1))

  expect(screen.getByText('Order is revoked')).toBeInTheDocument()
  expect(screen.getByText('Revoked on date')).toBeInTheDocument()

  fireEvent.click(screen.getByTestId('close-button'))

  await waitFor(() =>
    expect(history.location.pathname).toEqual('/orders/outgoing'),
  )
})

test('display error for a non-existing order', async () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { findByTestId, getByText } = renderWithAllProviders(
    <StaticRouter location="/delegatee/reference">
      <Route path="/:delegatee/:reference">
        <StoreProvider rootStore={rootStore}>
          <OrderDetailPage revokeHandler={() => {}} />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )
  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the order for delegatee')

  expect(getByText('Order not found')).toBeInTheDocument()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
