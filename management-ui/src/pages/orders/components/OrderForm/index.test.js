// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import userEvent from '@testing-library/user-event'
import { waitFor, fireEvent, screen } from '@testing-library/react'
import selectEvent, { openMenu } from 'react-select-event'
import { renderWithProviders } from '../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../stores'
import { DirectoryApi, ManagementApi } from '../../../../api'
import { ACCESS_REQUEST_STATES } from '../../../../stores/models/OutgoingAccessRequestModel'
import OrderForm from './index'

test('with initial values', async () => {
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

  const { container } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <OrderForm
        initialValues={{
          description: 'my-description',
          reference: 'my-reference',
          delegatee: '01234567890123456789',
          publicKeyPEM: 'my-public-key-pem',
          validFrom: new Date('2022-03-01T00:00:00.000Z'),
          validUntil: new Date('2022-03-02T00:00:00.000Z'),
          accessProofIds: ['1'],
        }}
        submitButtonText="Edit order"
        onSubmitHandler={onSubmitHandlerMock}
      />
    </StoreProvider>,
  )

  expect(screen.getByLabelText(/Order description/).value).toBe(
    'my-description',
  )

  expect(screen.getByLabelText(/Reference/).value).toBe('my-reference')

  expect(screen.getByLabelText(/Delegated organization/).value).toBe(
    '01234567890123456789',
  )

  expect(screen.getByLabelText(/Public key PEM/).value).toBe(
    'my-public-key-pem',
  )

  expect(screen.getByLabelText(/Valid from/).value).toBe('2022-03-01')

  expect(screen.getByLabelText(/Valid until/).value).toBe('2022-03-02')

  // we're asserting the selected value like this, since React-Select does not
  // set the input's value properly
  await waitFor(async () => {
    expect(
      container.querySelectorAll('.ReactSelect__multi-value__label')[0]
        .textContent,
    ).toBe(
      'service-a - organization-a (00000000000000000001) - via outway-a (h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=)',
    )
  })
})

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

  await userEvent.type(getByLabelText(/Order description/), 'my-description')
  await userEvent.type(getByLabelText(/Reference/), 'my-reference')
  await userEvent.type(
    getByLabelText(/Delegated organization/),
    '01234567890123456789',
  )
  await userEvent.type(getByLabelText(/Public key PEM/), 'my-public-key-pem')
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

  await userEvent.click(getByText('Add order'))

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

test('access proofs for which the same service has already been selected should be disabled', async () => {
  const onSubmitHandlerMock = jest.fn()

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'outway-a',
        publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
      },
      {
        name: 'outway-b',
        publicKeyFingerprint: 'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yd=',
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
          {
            accessRequest: {
              id: '2',
              state: ACCESS_REQUEST_STATES.APPROVED,
              publicKeyFingerprint:
                'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yd=',
            },
            accessProof: {
              id: '2',
              organization: {
                serialNumber: '00000000000000000001',
                name: 'organization-a',
              },
              serviceName: 'service-a',
            },
          },
        ],
      },
      {
        organization: {
          serialNumber: '00000000000000000001',
          name: 'organization-a',
        },
        serviceName: 'service-b',
        accessStates: [
          {
            accessRequest: {
              id: '3',
              state: ACCESS_REQUEST_STATES.APPROVED,
              publicKeyFingerprint:
                'h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=',
            },
            accessProof: {
              id: '3',
              organization: {
                serialNumber: '00000000000000000001',
                name: 'organization-a',
              },
              serviceName: 'service-b',
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

  const { getByLabelText, container } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <OrderForm
        submitButtonText="Add order"
        onSubmitHandler={onSubmitHandlerMock}
      />
    </StoreProvider>,
  )

  // open the dropdown and verify there are no disabled options
  openMenu(getByLabelText(/Services/))

  await waitFor(() => {
    expect(
      container.querySelector('.ReactSelect__menu-list'),
    ).toBeInTheDocument()
  })

  let options = container.querySelectorAll('.ReactSelect__option')
  expect(options).toHaveLength(3)

  for (const option of options) {
    expect(option.getAttribute('aria-disabled')).toEqual('false')
  }

  // select an option
  await selectEvent.select(
    getByLabelText(/Services/),
    /service-a - organization-a \(00000000000000000001\) - via outway-b \(h\+jpuLAMFzM09tOZpb0Ehslhje4S\/IsIxSWsS4E16Yd=\)/,
  )

  // open the dropdown and verify the remaining options
  openMenu(getByLabelText(/Services/))

  await waitFor(() => {
    expect(
      container.querySelector('.ReactSelect__menu-list'),
    ).toBeInTheDocument()
  })

  options = container.querySelectorAll('.ReactSelect__option')
  expect(options).toHaveLength(2)

  const optionServiceA = options[0]
  const optionServiceB = options[1]

  expect(optionServiceA.getAttribute('aria-disabled')).toEqual('true')
  expect(optionServiceA.innerHTML).toEqual(
    'service-a - organization-a (00000000000000000001) - via outway-a (h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=)',
  )

  expect(optionServiceB.getAttribute('aria-disabled')).toEqual('false')
  expect(optionServiceB.innerHTML).toEqual(
    'service-b - organization-a (00000000000000000001) - via outway-a (h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=)',
  )
})
