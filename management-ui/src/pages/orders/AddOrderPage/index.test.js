// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { configure } from 'mobx'
import userEvent from '@testing-library/user-event'
import selectEvent from 'react-select-event'
import { renderWithAllProviders, waitFor } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { DirectoryServiceApi, ManagementServiceApi } from '../../../api'
import { ACCESS_REQUEST_STATES } from '../../../stores/models/OutgoingAccessRequestModel'
import AddOrderPage from './index'

test('rendering the add order page', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceCreateOutgoingOrder = jest
    .fn()
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'outway-a',
          publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
        },
      ],
    })

  managementApiClient.managementServiceSynchronizeAllOutgoingAccessRequests =
    jest.fn().mockResolvedValue()

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceListServices = jest
    .fn()
    .mockResolvedValue({
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

  const history = createMemoryHistory()

  renderWithAllProviders(
    <HistoryRouter history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <AddOrderPage />
        </StoreProvider>
      </UserContextProvider>
    </HistoryRouter>,
  )

  jest.spyOn(rootStore.orderStore, 'create')

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
    '01234567890123456789',
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

  await userEvent.click(screen.getByText('Add order'))

  expect(rootStore.orderStore.create).toHaveBeenCalledWith({
    description: 'my-description',
    reference: 'my-reference',
    delegatee: '01234567890123456789',
    publicKeyPem: 'my-public-key-pem',
    validFrom: new Date('2021-01-01T00:00:00.000Z'),
    validUntil: new Date('2021-01-31T00:00:00.000Z'),
    accessProofIds: ['1'],
  })

  expect(screen.queryByRole('alert').textContent).toBe(
    "Failed to add orderYou don't have the required permission.",
  )

  await userEvent.click(screen.getByText('Add order'))

  expect(history.location.pathname).toEqual('/orders')
  expect(history.location.search).toEqual('?lastAction=added')
})
