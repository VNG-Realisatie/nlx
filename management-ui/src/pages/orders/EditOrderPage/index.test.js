// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import {
  Route,
  Routes,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { configure } from 'mobx'
import userEvent from '@testing-library/user-event'
import selectEvent from 'react-select-event'
import { renderWithProviders, waitFor } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { DirectoryApi, ManagementApi } from '../../../api'
import { ACCESS_REQUEST_STATES } from '../../../stores/models/OutgoingAccessRequestModel'
import EditOrderPage from './index'

test('rendering the edit order page', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          delegatee: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
          reference: 'my-reference',
        },
      ],
    })

  managementApiClient.managementUpdateOutgoingOrder = jest
    .fn()
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'outway-a',
        publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
      },
    ],
  })

  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryListServices = jest.fn().mockResolvedValue({
    services: [
      {
        organization: {
          serialNumber: '00000000000000000001',
          name: 'organization-a',
        },
        serviceName: 'service-a',
        accessStates: [
          {
            accessRequest: {
              id: '1',
              state: ACCESS_REQUEST_STATES.APPROVED,
              publicKeyFingerprint:
                'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
            },
            accessProof: {
              id: '1',
              organization: {
                serialNumber: '00000000000000000001',
                name: 'organization-a',
              },
              serviceName: 'service-a',
            },
          },
        ],
      },
    ],
  })

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  const history = createMemoryHistory({
    initialEntries: ['/00000000000000000001/my-reference'],
  })

  renderWithProviders(
    <HistoryRouter history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <Routes>
            <Route path="*" element={<div></div>} />
            <Route
              path="/:delegateeSerialNumber/:reference"
              element={<EditOrderPage />}
            />
          </Routes>
        </StoreProvider>
      </UserContextProvider>
    </HistoryRouter>,
  )

  jest.spyOn(rootStore.orderStore, 'updateOutgoing')

  await waitFor(() => {
    expect(screen.queryByRole('progressbar')).not.toBeInTheDocument()
  })

  await userEvent.type(
    screen.getByLabelText(/Order description/),
    'my-description',
  )
  await userEvent.type(screen.getByLabelText(/Reference/), 'my-reference')
  await userEvent.type(
    screen.getByLabelText(/Delegated organization/),
    '00000000000000000001',
  )
  await userEvent.type(
    screen.getByLabelText(/Public key PEM/),
    'my-public-key-pem',
  )
  fireEvent.change(screen.getByLabelText(/Valid from/), {
    target: { value: '2021-01-01' },
  })
  fireEvent.change(screen.getByLabelText(/Valid until/), {
    target: { value: '2021-01-31' },
  })
  await selectEvent.select(
    screen.getByRole('combobox'),
    /service-a - organization-a \(00000000000000000001\) - via outway-a \(h\+jpuLAMFzM09tOZpb0Ehslhje4S\/IsIxSWsS4E16Yc=\)/,
  )

  await userEvent.click(screen.getByText('Update order'))

  expect(rootStore.orderStore.updateOutgoing).toHaveBeenCalledWith({
    description: 'my-description',
    reference: 'my-reference',
    delegatee: '00000000000000000001',
    delegateeSerialNumber: '00000000000000000001',
    publicKeyPem: 'my-public-key-pem',
    validFrom: new Date('2021-01-01T00:00:00.000Z'),
    validUntil: new Date('2021-01-31T00:00:00.000Z'),
    accessProofIds: ['1'],
  })

  expect(screen.queryByRole('alert').textContent).toBe(
    "Failed to update the orderYou don't have the required permission.",
  )

  await userEvent.click(screen.getByText('Update order'))

  expect(history.location.pathname).toEqual(
    '/orders/outgoing/00000000000000000001/my-reference',
  )
})
