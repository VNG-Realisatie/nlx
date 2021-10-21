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
import IncomingOrderModel from '../../../../stores/models/IncomingOrderModel'
import Incoming from './index'

test('displays an order row for each order', () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const orders = [
    new IncomingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref1',
        description: 'description',
        delegator: '01234567890123456789',
        services: [],
        validUntil: '2021-05-10',
        validFrom: '2021-05-10',
      },
    }),
    new IncomingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref2',
        description: 'description',
        delegator: '01234567890123456789',
        services: [],
        validUntil: '2021-05-05',
        validFrom: '2021-05-05',
      },
    }),
  ]

  const history = createMemoryHistory()

  const { getAllByText } = renderWithProviders(
    <Router history={history}>
      <StoreProvider rootStore={rootStore}>
        <Incoming orders={orders} />
      </StoreProvider>
    </Router>,
  )
  expect(getAllByText('description')).toHaveLength(2)
})

test('displays text to indicate there are no orders', () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Incoming orders={[]} />
    </StoreProvider>,
  )
  expect(getByText("You haven't received any orders yet")).toBeInTheDocument()
})

test('content should render expected data', () => {
  const day = 86400000

  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const orders = [
    new IncomingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref1',
        description: 'my own description',
        delegator: '10000000000000000000',
        services: [
          {
            organization: {
              serialNumber: '00000000000000000001',
              name: 'organization X',
            },
            service: 'service Y',
          },
          {
            organization: {
              serialNumber: '00000000000000000002',
              name: 'organization Y',
            },
            service: 'service Z',
          },
        ],
        validFrom: new Date(new Date().getTime() - day).toISOString(),
        validUntil: new Date(new Date().getTime() + day).toISOString(),
        revokedAt: null,
      },
    }),
    new IncomingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref2',
        description: 'my own description',
        delegator: '10000000000000000000',
        services: [
          {
            organization: {
              serialNumber: '00000000000000000002',
              name: 'organization X',
            },
            service: 'service Y',
          },
        ],
        validFrom: new Date(new Date().getTime() - day).toISOString(),
        validUntil: new Date(new Date().getTime() - day).toISOString(),
        revokedAt: null,
      },
    }),
    new IncomingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref3',
        description: 'my own description',
        delegator: '10000000000000000000',
        services: [
          {
            organization: {
              serialNumber: '00000000000000000003',
              name: 'organization X',
            },
            service: 'service Y',
          },
        ],
        validFrom: new Date(new Date().getTime() + day).toISOString(),
        validUntil: new Date(new Date().getTime() + 2 * day).toISOString(),
        revokedAt: null,
      },
    }),
    new IncomingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref4',
        description: 'my own description',
        delegator: '10000000000000000000',
        services: [
          {
            organization: {
              serialNumber: '00000000000000000004',
              name: 'organization X',
            },
            service: 'service Y',
          },
        ],
        validFrom: new Date(new Date().getTime() - day).toISOString(),
        validUntil: new Date(new Date().getTime() + day).toISOString(),
        revokedAt: new Date(),
      },
    }),
  ]

  const history = createMemoryHistory({})

  const { container } = renderWithAllProviders(
    <Router history={history}>
      <StoreProvider rootStore={rootStore}>
        <Incoming orders={orders} />
      </StoreProvider>
    </Router>,
  )

  const firstOrderEl = container.querySelectorAll('tbody tr')[0]
  const firstOrder = within(firstOrderEl)
  expect(firstOrder.getByTitle('Active')).toBeInTheDocument()
  expect(firstOrder.getByText('my own description')).toBeInTheDocument()
  expect(firstOrder.getByText('10000000000000000000')).toBeInTheDocument()
  expect(
    firstOrder.getByTitle('organization X (00000000000000000001) - service Y'),
  ).toBeInTheDocument()
  expect(
    firstOrder.getByTitle('organization Y (00000000000000000002) - service Z'),
  ).toHaveTextContent('organization Y (00000000000000000002) - service Z')

  const secondOrder = container.querySelectorAll('tbody tr')[1]
  expect(within(secondOrder).getByTitle('Inactive')).toBeInTheDocument()

  const thirdOrder = container.querySelectorAll('tbody tr')[2]
  expect(within(thirdOrder).getByTitle('Inactive')).toBeInTheDocument()

  const fourthOrder = container.querySelectorAll('tbody tr')[3]
  expect(within(fourthOrder).getByTitle('Inactive')).toBeInTheDocument()

  fireEvent.click(firstOrderEl)
  expect(history.location.pathname).toEqual(
    '/orders/incoming/10000000000000000000/ref1',
  )
})
