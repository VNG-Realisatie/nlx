// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, within } from '@testing-library/react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../../test-utils'
import OrdersView from '../OrdersOutgoing'

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
      delegatee: 'delegatee',
      services: [
        { organization: 'organization X', service: 'service Y' },
        { organization: 'organization Y', service: 'service Z' },
      ],
      validUntil: '2021-05-10',
      revokedAt: null,
    },
    {
      reference: 'ref2',
      description: 'my own description',
      delegatee: 'delegatee',
      services: [{ organization: 'organization X', service: 'service Y' }],
      validUntil: '2021-05-10',
      revokedAt: '2021-04-10',
    },
  ]

  const history = createMemoryHistory({
    initialEntries: ['/orders'],
  })

  const { container } = renderWithProviders(
    <Router history={history}>
      <OrdersView orders={orders} />
    </Router>,
  )

  const firstOrderEl = container.querySelectorAll('tbody tr')[0]
  const firstOrder = within(firstOrderEl)
  expect(firstOrder.getByTitle('Active')).toBeInTheDocument()
  expect(firstOrder.getByText('my own description')).toBeInTheDocument()
  expect(firstOrder.getByText('delegatee')).toBeInTheDocument()
  expect(
    firstOrder.getByTitle('organization X - service Y'),
  ).toBeInTheDocument()
  expect(firstOrder.getByTitle('organization Y - service Z')).toHaveTextContent(
    'organization Y - service Z',
  )

  const secondOrder = container.querySelectorAll('tbody tr')[1]
  expect(within(secondOrder).getByTitle('Inactive')).toBeInTheDocument()

  fireEvent.click(firstOrderEl)
  expect(history.location.pathname).toEqual('/orders/outgoing/delegatee/ref1')
})
