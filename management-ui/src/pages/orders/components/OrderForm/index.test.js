// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import userEvent from '@testing-library/user-event'
import { waitFor, fireEvent } from '@testing-library/react'
import selectEvent from 'react-select-event'
import { renderWithProviders } from '../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../stores'
import { DirectoryApi, ManagementApi } from '../../../../api'
import { ACCESS_REQUEST_STATES } from '../../../../stores/models/OutgoingAccessRequestModel'
import OrderForm from './index'

test('the form values of the onSubmitHandler', async () => {
  const onSubmitHandlerMock = jest.fn()

  const managementApiClient = new ManagementApi()

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

  const { getByLabelText, getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <OrderForm
        submitButtonText="Add order"
        onSubmitHandler={onSubmitHandlerMock}
      />
    </StoreProvider>,
  )

  userEvent.type(getByLabelText(/Order description/), 'my-description')
  userEvent.type(getByLabelText(/Reference/), 'my-reference')
  userEvent.type(
    getByLabelText(/Delegated organization/),
    '01234567890123456789',
  )
  userEvent.type(getByLabelText(/Public key PEM/), 'my-public-key-pem')
  fireEvent.change(getByLabelText(/Valid from/), {
    target: { value: '2021-01-01' },
  })
  fireEvent.change(getByLabelText(/Valid until/), {
    target: { value: '2021-01-31' },
  })
  await selectEvent.select(
    getByLabelText(/Services/),
    /service-a - organization-a \(00000000000000000001\) - via outway-a \(h\+jpuLAMFzM09tOZpb0Ehslhje4S\/IsIxSWsS4E16Yc=\)/,
  )

  userEvent.click(getByText('Add order'))

  await waitFor(() =>
    expect(onSubmitHandlerMock).toHaveBeenCalledWith({
      description: 'my-description',
      reference: 'my-reference',
      delegatee: '01234567890123456789',
      publicKeyPEM: 'my-public-key-pem',
      validFrom: new Date('2021-01-01T00:00:00.000Z'),
      validUntil: new Date('2021-01-31T00:00:00.000Z'),
      accessProofIds: ['1'],
    }),
  )
})
