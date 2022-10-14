// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, screen, within } from '@testing-library/react'
import {
  Route,
  Routes,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { RootStore, StoreProvider } from '../../../../stores'
import {
  renderWithProviders,
  renderWithAllProviders,
} from '../../../../test-utils'
import { ManagementServiceApi } from '../../../../api'
import OutgoingOrderModel from '../../../../stores/models/OutgoingOrderModel'
import Outgoing from './index'

test('displays an order row for each order', () => {
  const managementApiClient = new ManagementServiceApi()
  const rootStore = new RootStore({ managementApiClient })

  const orders = [
    new OutgoingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref1',
        description: 'my own description',
        delegatee: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        validUntil: '2021-05-10',
        accessProofIds: [],
      },
    }),
    new OutgoingOrderModel({
      orderStore: rootStore.orderStore,
      orderData: {
        reference: 'ref2',
        description: 'my own description',
        delegatee: {
          serialNumber: '00000000000000000001',
          name: 'Organization One',
        },
        validUntil: '2021-05-05',
        accessProofIds: [],
      },
    }),
  ]

  const history = createMemoryHistory()

  renderWithProviders(
    <HistoryRouter history={history}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path="/*" element={<Outgoing orders={orders} />} />
        </Routes>
      </StoreProvider>
    </HistoryRouter>,
  )
  expect(screen.getAllByText('my own description')).toHaveLength(2)
})

test('displays text to indicate there are no orders', () => {
  const managementApiClient = new ManagementServiceApi()
  const rootStore = new RootStore({ managementApiClient })

  renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Outgoing orders={[]} />
    </StoreProvider>,
  )
  expect(
    screen.getByText("You don't have any issued orders yet"),
  ).toBeInTheDocument()
})

test('content should render expected data', async () => {
  const oneDay = 86400000

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          reference: 'ref1',
          description: 'my own description',
          delegatee: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
          validFrom: new Date(new Date().getTime() - oneDay).toISOString(),
          validUntil: new Date(new Date().getTime() + oneDay).toISOString(),
          revokedAt: null,
          accessProofs: [
            {
              id: 1,
              organization: {
                serialNumber: '0000000000000000000',
                name: 'Organization Two',
              },
              serviceName: 'service A',
              publicKeyFingerprint: 'public-key-fingerprint-a',
            },
            {
              id: 2,
              organization: {
                serialNumber: '00000000000000000003',
                name: 'Organization Three',
              },
              serviceName: 'service B',
              publicKeyFingerprint: 'public-key-fingerprint-b',
            },
          ],
        },
        {
          reference: 'ref2',
          description: 'my own description',
          delegatee: '20000000000000000000',
          validFrom: new Date(new Date().getTime() - oneDay).toISOString(),
          validUntil: new Date(new Date().getTime() - oneDay).toISOString(),
          revokedAt: null,
          accessProofs: [],
        },
      ],
    })

  const rootStore = new RootStore({ managementApiClient })
  const orderStore = rootStore.orderStore

  await orderStore.fetchOutgoing()

  const history = createMemoryHistory()

  const { container } = renderWithAllProviders(
    <HistoryRouter history={history}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route
            path="/*"
            element={<Outgoing orders={orderStore.outgoingOrders} />}
          />
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
    firstOrder.getByTitle('Organization Two - service A'),
  ).toBeInTheDocument()
  expect(
    firstOrder.getByTitle('Organization Three - service B'),
  ).toBeInTheDocument()

  const secondOrder = container.querySelectorAll('tbody tr')[1]
  expect(within(secondOrder).getByTitle('Inactive')).toBeInTheDocument()

  fireEvent.click(firstOrderEl)
  expect(history.location.pathname).toEqual(
    '/orders/outgoing/00000000000000000001/ref1',
  )
})
