// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import {
  Route,
  Routes,
  MemoryRouter,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
import { fireEvent, waitFor, within } from '@testing-library/react'
import { configure } from 'mobx'
import { createMemoryHistory } from 'history'
import { renderWithAllProviders, screen } from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { DirectoryServiceApi, ManagementServiceApi } from '../../../../../api'
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

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [],
    })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegatee: '01234567890123456789',
          reference: 'my-reference',
          description: 'description',
          validFrom: '2020-01-01',
          validUntil: '3000-01-01',
          revokedAt: null,
          services: [],
        },
      ],
    })

  managementApiClient.managementServiceRevokeOutgoingOrder = jest
    .fn()
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  const orderStore = rootStore.orderStore

  await orderStore.fetchOutgoing()

  const history = createMemoryHistory({
    initialEntries: ['/01234567890123456789/my-reference'],
  })

  renderWithAllProviders(
    <HistoryRouter history={history}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":delegatee/:reference" element={<OrderDetailPage />} />
        </Routes>
      </StoreProvider>
    </HistoryRouter>,
  )

  expect(screen.getByText('Issued to delegatee')).toBeInTheDocument()
  expect(screen.getByText('description')).toBeInTheDocument()

  const orderModel = orderStore.outgoingOrders[0]
  jest.spyOn(orderModel, 'revoke')

  let revokeButton = await screen.findByText(/Revoke/)
  fireEvent.click(revokeButton)

  let confirmModal = screen.getByRole('dialog')
  let okButton = within(confirmModal).getByText(/Revoke/)

  fireEvent.click(okButton)
  await waitFor(() => expect(orderModel.revoke).toHaveBeenCalledTimes(1))

  expect(screen.queryByRole('alert').textContent).toBe(
    "Failed to revoke the orderYou don't have the required permission.",
  )

  revokeButton = await screen.findByText(/Revoke/)
  fireEvent.click(revokeButton)

  confirmModal = screen.getByRole('dialog')
  okButton = within(confirmModal).getByText(/Revoke/)

  fireEvent.click(okButton)
  await waitFor(() => expect(orderModel.revoke).toHaveBeenCalledTimes(2))

  expect(screen.getByText(/Order is revoked/)).toBeInTheDocument()
  expect(screen.getByText(/Revoked on date/)).toBeInTheDocument()

  fireEvent.click(screen.getByTestId('close-button'))

  await waitFor(() =>
    expect(history.location.pathname).toEqual('/orders/outgoing'),
  )
})

test('display error for a non-existing order', async () => {
  const managementApiClient = new ManagementServiceApi()
  const rootStore = new RootStore({ managementApiClient })

  renderWithAllProviders(
    <MemoryRouter initialEntries={['/delegatee/reference']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":delegatee/:reference" element={<OrderDetailPage />} />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )
  const message = await screen.findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the order for delegatee')

  expect(screen.getByText('Order not found')).toBeInTheDocument()

  const closeButton = await screen.findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
