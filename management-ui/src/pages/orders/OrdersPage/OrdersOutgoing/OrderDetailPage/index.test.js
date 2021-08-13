// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { ManagementApi } from '../../../../../api'
import OrderDetailPage from './index'

test('display service details', () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })
  const { getByText } = renderWithProviders(
    <StaticRouter location="/orders/delegatee/reference">
      <Route path="/orders/:delegatee/:reference">
        <StoreProvider rootStore={rootStore}>
          <OrderDetailPage
            order={{
              delegatee: 'delegatee',
              reference: 'reference',
              description: 'description',
            }}
          />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )

  expect(getByText('description')).toBeInTheDocument()
})

test('display error for a non-existing order', async () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { findByTestId, getByText } = renderWithProviders(
    <StaticRouter location="/orders/delegatee/reference">
      <Route path="/orders/:delegatee/:reference">
        <StoreProvider rootStore={rootStore}>
          <OrderDetailPage />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )
  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the order')

  expect(getByText('Order not found')).toBeInTheDocument()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
