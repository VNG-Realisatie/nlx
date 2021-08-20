// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import { ManagementApi } from '../../../../../../api'
import { RootStore } from '../../../../../../stores'
import OutgoingOrderModel from '../../../../../../stores/models/OutgoingOrderModel'
import OrderDetailView from './index'

test('display order details', () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const order = new OutgoingOrderModel({
    orderStore: rootStore.orderStore,
    orderData: {
      reference: 'my-reference',
      delegatee: 'delegatee',
      validFrom: '2021-01-01T00:00:00.000Z',
      validUntil: '3000-01-01T00:00:00.000Z',
      services: [],
      revokedAt: null,
    },
  })

  renderWithProviders(
    <OrderDetailView order={order} revokeHandler={() => {}} />,
  )

  expect(screen.getByTestId('status')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'up.svg' + 'Order is active' + 'Revoke',
  )
  expect(screen.getByTestId('start-end-date')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'timer.svg' + 'Valid until date' + 'Since date',
  )
  expect(screen.getByText('my-reference')).toBeInTheDocument()
  expect(screen.getByText('Requestable services')).toBeInTheDocument()
})
