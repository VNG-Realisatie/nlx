// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Router, StaticRouter } from 'react-router-dom'
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

  managementApiClient.managementListIncomingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegator: 'delegator',
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
    initialEntries: ['/delegator/reference'],
  })

  renderWithAllProviders(
    <Router history={history}>
      <Route path="/:delegator/:reference">
        <StoreProvider rootStore={rootStore}>
          <OrderDetailPage order={orderStore.incomingOrders[0]} />
        </StoreProvider>
      </Route>
    </Router>,
  )

  expect(screen.getByText('Issued by delegator')).toBeInTheDocument()
  expect(screen.getByText('description')).toBeInTheDocument()
})

test('display error for a non-existing order', async () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { findByTestId, getByText } = renderWithAllProviders(
    <StaticRouter location="/delegator/reference">
      <Route path="/:delegator/:reference">
        <StoreProvider rootStore={rootStore}>
          <OrderDetailPage revokeHandler={() => {}} />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )
  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe(
    'Failed to load the order issued by delegator',
  )

  expect(getByText('Order not found')).toBeInTheDocument()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
