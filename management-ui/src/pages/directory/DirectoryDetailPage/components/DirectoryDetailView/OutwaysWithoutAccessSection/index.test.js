// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import { DirectoryApi, ManagementApi } from '../../../../../../api'
import { RootStore, StoreProvider } from '../../../../../../stores'
import OutwaysWithoutAccessSection from './index'

test('Outways without access section', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'outway-1',
        publicKeyFingerprint: 'public-key-fingerprint-1',
      },
      {
        name: 'outway-2',
        publicKeyFingerprint: 'public-key-fingerprint-1',
      },
    ],
  })

  const directoryApiClient = new DirectoryApi()

  directoryApiClient.directoryGetOrganizationService = jest
    .fn()
    .mockResolvedValue({
      name: 'my-service',
      organization: {
        organizationName: 'my-organization',
        organizationSerialNumber: '00000000000000000001',
      },
      accessStates: [],
    })

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  const service = await rootStore.directoryServicesStore.fetch(
    '00000000000000000001',
    'my-service',
  )

  const requestAccessHandler = jest.fn()

  const { container } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OutwaysWithoutAccessSection
          service={service}
          requestAccessHandler={requestAccessHandler}
        />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(container).toHaveTextContent(/None/)

  await rootStore.outwayStore.fetchAll()

  fireEvent.click(screen.getByText(/Outways without access/i))

  expect(await screen.findByTestId('outway-names')).toHaveTextContent(
    'outway-1, outway-2',
  )

  fireEvent.click(screen.getByText('Request access'))

  expect(requestAccessHandler).toHaveBeenCalledWith('public-key-fingerprint-1')
})
