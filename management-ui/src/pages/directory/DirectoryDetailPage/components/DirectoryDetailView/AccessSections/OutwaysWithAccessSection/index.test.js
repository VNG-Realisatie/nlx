// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { fireEvent, screen, waitFor, within } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../../../../../../test-utils'
import {
  DirectoryServiceApi,
  ManagementServiceApi,
} from '../../../../../../../api'
import { RootStore, StoreProvider } from '../../../../../../../stores'
import { ACCESS_REQUEST_STATES } from '../../../../../../../stores/models/OutgoingAccessRequestModel'
import OutwaysWithAccessSection from './index'

test('Outways with access section', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'outway-1',
          publicKeyFingerprint: 'public-key-fingerprint-1',
          publicKeyPem: 'public-key-pem-1',
        },
        {
          name: 'outway-2',
          publicKeyFingerprint: 'public-key-fingerprint-1',
          publicKeyPem: 'public-key-pem-1',
        },
      ],
    })

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceGetOrganizationService = jest
    .fn()
    .mockResolvedValue({
      directoryService: {
        name: 'my-service',
        organization: {
          organizationName: 'my-organization',
          organizationSerialNumber: '00000000000000000001',
        },
        accessStates: [
          {
            accessRequest: {
              state: ACCESS_REQUEST_STATES.APPROVED,
              publicKeyFingerprint: 'public-key-fingerprint-1',
            },
            accessProof: {},
          },
        ],
      },
    })

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  const service = await rootStore.directoryServicesStore.fetch(
    '00000000000000000001',
    'my-service',
  )

  const { container } = renderWithAllProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OutwaysWithAccessSection service={service} />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(container).toHaveTextContent(/None/)

  await rootStore.outwayStore.fetchAll()

  fireEvent.click(screen.getByText(/Outways with access/i))

  expect(screen.getByText('outway-1')).toBeInTheDocument()
  expect(screen.getByText('outway-2')).toBeInTheDocument()
})

test('Terminate access', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'outway-1',
          publicKeyFingerprint: 'public-key-fingerprint-1',
          publicKeyPem: 'public-key-pem-1',
        },
        {
          name: 'outway-2',
          publicKeyFingerprint: 'public-key-fingerprint-1',
          publicKeyPem: 'public-key-pem-1',
        },
      ],
    })

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceGetOrganizationService = jest
    .fn()
    .mockResolvedValue({
      directoryService: {
        name: 'my-service',
        organization: {
          organizationName: 'my-organization',
          organizationSerialNumber: '00000000000000000001',
        },
        accessStates: [
          {
            accessRequest: {
              id: 42,
              state: ACCESS_REQUEST_STATES.APPROVED,
              publicKeyFingerprint: 'public-key-fingerprint-1',
            },
            accessProof: {},
          },
        ],
      },
    })

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  const service = await rootStore.directoryServicesStore.fetch(
    '00000000000000000001',
    'my-service',
  )

  renderWithAllProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OutwaysWithAccessSection service={service} />
      </StoreProvider>
    </MemoryRouter>,
  )

  const outgoingAccessRequest =
    rootStore.outgoingAccessRequestStore.outgoingAccessRequests.get(42)

  const terminateSpy = jest
    .spyOn(outgoingAccessRequest, 'terminate')
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  await rootStore.outwayStore.fetchAll()

  fireEvent.click(screen.getByText(/Outways with access/i))

  fireEvent.click(screen.getByText('Terminate'))

  let confirmModal = screen.getByRole('dialog')
  const cancelButton = within(confirmModal).getByText('Cancel')
  fireEvent.click(cancelButton)

  await waitFor(() => expect(terminateSpy).not.toHaveBeenCalled())

  fireEvent.click(
    screen.getByTitle(
      'Terminate access for Outways with public key fingerprint {{publicKeyFingerprint}}',
    ),
  )

  confirmModal = screen.getByRole('dialog')
  let okButton = within(confirmModal).getByText('Terminate')
  fireEvent.click(okButton)

  await waitFor(() => expect(terminateSpy).toHaveBeenCalled())

  expect(screen.queryByRole('alert')).toHaveTextContent(
    "Failed to terminate accessYou don't have the required permission.",
  )

  fireEvent.click(screen.getByText('Terminate'))

  confirmModal = screen.getByRole('dialog')
  okButton = within(confirmModal).getByText('Terminate')
  fireEvent.click(okButton)

  // toast
  expect(await screen.findByText('Access terminated')).toBeInTheDocument()
})
