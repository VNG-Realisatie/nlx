// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import OrdersView from './index'

jest.mock('./OrderRow', () => (props) => (
  <tr {...props}>
    <td>order row</td>
  </tr>
))

test('displays an order row for each order', () => {
  const orders = [
    {
      reference: 'ref1',
      description: 'my own description',
      delegatee: 'delegatee',
      services: [{ organization: 'organization X', service: 'service Y' }],
      validUntil: '2021-05-10',
    },
    {
      reference: 'ref2',
      description: 'my other description',
      delegatee: 'goatadelee',
      services: [{ organization: 'organization Z', service: 'service S' }],
      validUntil: '2021-05-05',
    },
  ]

  const { getAllByText } = renderWithProviders(<OrdersView orders={orders} />)
  expect(getAllByText('order row')).toHaveLength(2)
})

test('displays text to indicate there are no orders', () => {
  const { getByText } = renderWithProviders(<OrdersView ordersMap={[]} />)
  expect(getByText('There are no active orders')).toBeInTheDocument()
})
