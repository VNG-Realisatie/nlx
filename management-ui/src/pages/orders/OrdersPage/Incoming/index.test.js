// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, within } from '@testing-library/react'
import {
  Routes,
  Route,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
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
        delegator: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
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
        delegator: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        services: [],
        validUntil: '2021-05-05',
        validFrom: '2021-05-05',
      },
    }),
  ]

  const history = createMemoryHistory()

  const { getAllByText } = renderWithProviders(
    <HistoryRouter history={history}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path="*" element={<Incoming orders={orders} />} />
        </Routes>
      </StoreProvider>
    </HistoryRouter>,
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
        delegator: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        services: [
          {
            organization: {
              serialNumber: '00000000000000000002',
              name: 'Organization Two',
            },
            service: 'service Y',
          },
          {
            organization: {
              serialNumber: '00000000000000000003',
              name: 'Organization Three',
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
        delegator: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        services: [
          {
            organization: {
              serialNumber: '00000000000000000002',
              name: 'Organization Two',
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
        delegator: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        services: [
          {
            organization: {
              serialNumber: '00000000000000000003',
              name: 'Organization Three',
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
        delegator: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        services: [
          {
            organization: {
              serialNumber: '00000000000000000004',
              name: 'Organization Four',
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
    <HistoryRouter history={history}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path="*" element={<Incoming orders={orders} />} />
        </Routes>
      </StoreProvider>
    </HistoryRouter>,
  )

  const firstOrderEl = container.querySelectorAll('tbody tr')[0]
  const firstOrder = within(firstOrderEl)
  expect(firstOrder.getByTitle('Active')).toBeInTheDocument()
  expect(firstOrder.getByText('my own description')).toBeInTheDocument()
  expect(firstOrder.getByText('Organization One')).toBeInTheDocument()
  expect(
    firstOrder.getByTitle('Organization Two - service Y'),
  ).toBeInTheDocument()
  expect(
    firstOrder.getByTitle('Organization Three - service Z'),
  ).toHaveTextContent('Organization Three - service Z')

  const secondOrder = container.querySelectorAll('tbody tr')[1]
  expect(within(secondOrder).getByTitle('Inactive')).toBeInTheDocument()

  const thirdOrder = container.querySelectorAll('tbody tr')[2]
  expect(within(thirdOrder).getByTitle('Inactive')).toBeInTheDocument()

  const fourthOrder = container.querySelectorAll('tbody tr')[3]
  expect(within(fourthOrder).getByTitle('Inactive')).toBeInTheDocument()

  fireEvent.click(firstOrderEl)
  expect(history.location.pathname).toEqual(
    '/orders/incoming/00000000000000000001/ref1',
  )
})
