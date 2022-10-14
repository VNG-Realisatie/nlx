// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import {
  DirectoryServiceApi,
  ManagementServiceApi,
} from '../../../../../../api'
import { RootStore, StoreProvider } from '../../../../../../stores'
import OutgoingOrderModel from '../../../../../../stores/models/OutgoingOrderModel'
import OrderDetailView from './index'

test('display order details', () => {
  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [],
    })

  const managementApiClient = new ManagementServiceApi()

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  const order = new OutgoingOrderModel({
    orderStore: rootStore.orderStore,
    orderData: {
      reference: 'my-reference',
      delegatee: 'delegatee',
      validFrom: '2021-01-01T00:00:00.000Z',
      validUntil: '3000-01-01T00:00:00.000Z',
      revokedAt: null,
    },
  })

  renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <OrderDetailView order={order} revokeHandler={() => {}} />
    </StoreProvider>,
  )

  expect(screen.getByTestId('status')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'state-up.svg' + 'Order is active',
  )
  expect(screen.getByTestId('start-end-date')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'timer.svg' + 'Valid until date' + 'Since date',
  )
  expect(screen.getByText('my-reference')).toBeInTheDocument()
  expect(screen.getByText('Requestable services')).toBeInTheDocument()
})
