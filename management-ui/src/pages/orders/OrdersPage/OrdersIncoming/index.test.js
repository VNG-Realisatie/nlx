// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import OrdersView from './index'

test('displays an order row for each order', () => {
  const orders = [
    {
      reference: 'ref1',
      description: 'my own description',
      delegator: 'delegator',
      services: [{ organization: 'organization X', service: 'service Y' }],
      validUntil: '2021-05-10',
    },
    {
      reference: 'ref2',
      description: 'my own description',
      delegator: 'goatadelee',
      services: [{ organization: 'organization Z', service: 'service S' }],
      validUntil: '2021-05-05',
    },
  ]

  const { getAllByText } = renderWithProviders(<OrdersView orders={orders} />)
  expect(getAllByText('my own description')).toHaveLength(2)
})

test('displays text to indicate there are no orders', () => {
  const { getByText } = renderWithProviders(<OrdersView ordersMap={[]} />)
  expect(getByText("You haven't received any orders yet")).toBeInTheDocument()
})

test('content should render expected data', () => {
  const orders = [
    {
      reference: 'ref1',
      description: 'my own description',
      delegator: 'delegator',
      services: [
        { organization: 'organization X', service: 'service Y' },
        { organization: 'organization Y', service: 'service Z' },
      ],
      validUntil: '2021-05-10',
    },
  ]

  const { getByText, getByTitle } = renderWithProviders(
    <OrdersView orders={orders} />,
  )

  expect(getByText('my own description')).toBeInTheDocument()
  expect(getByText('delegator')).toBeInTheDocument()
  expect(getByTitle('organization X - service Y')).toBeInTheDocument()
  expect(getByTitle('organization Y - service Z')).toHaveTextContent(
    'organization Y - service Z',
  )
})
