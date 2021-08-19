// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, within } from '@testing-library/react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { RootStore, StoreProvider } from '../../../../stores'
import {
  renderWithProviders,
  renderWithAllProviders,
} from '../../../../test-utils'
import { ManagementApi } from '../../../../api'
import Outgoing from './index'

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

  const history = createMemoryHistory({
    initialEntries: ['/orders'],
  })

  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { getAllByText } = renderWithProviders(
    <Router history={history}>
      <StoreProvider rootStore={rootStore}>
        <Outgoing orders={orders} />
      </StoreProvider>
    </Router>,
  )
  expect(getAllByText('my own description')).toHaveLength(2)
})

test('displays text to indicate there are no orders', () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Outgoing orders={[]} />
    </StoreProvider>,
  )
  expect(getByText("You don't have any issued orders yet")).toBeInTheDocument()
})

test('content should render expected data', () => {
  const day = 86400000

  const orders = [
    {
      reference: 'ref1',
      description: 'my own description',
      delegatee: 'delegatee',
      services: [
        { organization: 'organization X', service: 'service Y' },
        { organization: 'organization Y', service: 'service Z' },
      ],
      validFrom: new Date(new Date().getTime() - day),
      validUntil: new Date(new Date().getTime() + day),
      revokedAt: null,
    },
    {
      reference: 'ref2',
      description: 'my own description',
      delegatee: 'delegatee',
      services: [{ organization: 'organization X', service: 'service Y' }],
      validFrom: new Date(new Date().getTime() - day),
      validUntil: new Date(new Date().getTime() - day),
      revokedAt: null,
    },
    {
      reference: 'ref3',
      description: 'my own description',
      delegatee: 'delegatee',
      services: [{ organization: 'organization X', service: 'service Y' }],
      validFrom: new Date(new Date().getTime() + day),
      validUntil: new Date(new Date().getTime() + 2 * day),
      revokedAt: null,
    },
    {
      reference: 'ref4',
      description: 'my own description',
      delegatee: 'delegatee',
      services: [{ organization: 'organization X', service: 'service Y' }],
      validFrom: new Date(new Date().getTime() - day),
      validUntil: new Date(new Date().getTime() + day),
      revokedAt: new Date(),
    },
  ]

  const history = createMemoryHistory({
    initialEntries: ['/orders'],
  })

  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { container } = renderWithAllProviders(
    <Router history={history}>
      <StoreProvider rootStore={rootStore}>
        <Outgoing orders={orders} />
      </StoreProvider>
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

  const thirdOrder = container.querySelectorAll('tbody tr')[2]
  expect(within(thirdOrder).getByTitle('Inactive')).toBeInTheDocument()

  const fourthOrder = container.querySelectorAll('tbody tr')[3]
  expect(within(fourthOrder).getByTitle('Inactive')).toBeInTheDocument()

  fireEvent.click(firstOrderEl)
  expect(history.location.pathname).toEqual('/orders/outgoing/delegatee/ref1')
})
