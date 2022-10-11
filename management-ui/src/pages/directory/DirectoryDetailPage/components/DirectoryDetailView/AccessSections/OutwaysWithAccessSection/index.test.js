// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../../test-utils'
import { DirectoryApi, ManagementApi } from '../../../../../../../api'
import { RootStore, StoreProvider } from '../../../../../../../stores'
import { ACCESS_REQUEST_STATES } from '../../../../../../../stores/models/OutgoingAccessRequestModel'
import OutwaysWithAccessSection from './index'

test('Outways with access section', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
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

  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryGetOrganizationService = jest
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

  const { container } = renderWithProviders(
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
