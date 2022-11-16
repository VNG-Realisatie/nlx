// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders, screen } from '../../../../../../../test-utils'
import OutgoingOrderModel from '../../../../../../../stores/models/OutgoingOrderModel'
import Status from './index'

const day = 86400000

test('display status', async () => {
  const order = new OutgoingOrderModel({
    orderData: {
      reference: 'my-reference',
      delegatee: 'delegatee',
      validFrom: new Date(new Date().getTime() - day),
      validUntil: new Date(new Date().getTime() + day),
      revokedAt: null,
    },
  })

  renderWithProviders(<Status order={order} revokeHandler={() => {}} />)

  expect(screen.queryByText('Order is active')).toBeInTheDocument()

  order.update({
    orderData: {
      validFrom: new Date(new Date().getTime() - day),
      validUntil: new Date(new Date().getTime() - day),
    },
  })

  expect(await screen.findByText('Order is expired')).toBeInTheDocument()

  order.update({
    orderData: {
      validFrom: new Date(new Date().getTime() + day),
      validUntil: new Date(new Date().getTime() + 2 * day),
    },
  })

  expect(await screen.findByText('Order is not yet active')).toBeInTheDocument()

  order.update({
    orderData: {
      revokedAt: new Date(),
    },
  })

  expect(await screen.findByText('Order is revoked')).toBeInTheDocument()
})
