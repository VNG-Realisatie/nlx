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
import { configure } from 'mobx'
import { createMemoryHistory } from 'history'
import { renderWithAllProviders, screen } from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { ManagementServiceApi } from '../../../../../api'
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

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListIncomingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegator: '01234567890123456789',
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

  await orderStore.fetchIncoming()

  const history = createMemoryHistory({
    initialEntries: ['/01234567890123456789/my-reference'],
  })

  renderWithAllProviders(
    <HistoryRouter history={history}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":delegator/:reference" element={<OrderDetailPage />} />
        </Routes>
      </StoreProvider>
    </HistoryRouter>,
  )

  expect(screen.getByText('Issued by delegator')).toBeInTheDocument()
  expect(screen.getByText('description')).toBeInTheDocument()
})

test('display error for a non-existing order', async () => {
  const managementApiClient = new ManagementServiceApi()
  const rootStore = new RootStore({ managementApiClient })

  renderWithAllProviders(
    <MemoryRouter initialEntries={['/delegator/reference']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":delegator/:reference" element={<OrderDetailPage />} />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )
  const message = await screen.findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe(
    'Failed to load the order issued by delegator',
  )

  expect(screen.getByText('Order not found')).toBeInTheDocument()

  const closeButton = await screen.findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
