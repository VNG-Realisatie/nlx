// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import IncomingOrderModel from '../../../../../../stores/models/IncomingOrderModel'
import { ManagementServiceApi } from '../../../../../../api'
import { RootStore } from '../../../../../../stores'
import OrderDetailView from './index'

test('display order details', () => {
  const managementApiClient = new ManagementServiceApi()
  const rootStore = new RootStore({ managementApiClient })

  const order = new IncomingOrderModel({
    orderStore: rootStore.orderStore,
    orderData: {
      reference: 'my-reference',
      delegator: '01234567890123456789',
      validFrom: '2021-01-01T00:00:00.000Z',
      validUntil: '3000-01-01T00:00:00.000Z',
      services: [],
      revokedAt: null,
    },
  })

  renderWithProviders(<OrderDetailView order={order} />)

  expect(screen.getByTestId('status')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'up.svg' + 'Order is active',
  )
  expect(screen.getByTestId('start-end-date')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'timer.svg' + 'Valid until date' + 'Since date',
  )
  expect(screen.getByText('my-reference')).toBeInTheDocument()
  expect(screen.getByText('Requestable services')).toBeInTheDocument()
})
