// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { fireEvent, screen, waitFor } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithProviders } from '../../../../../../../test-utils'
import { DirectoryApi, ManagementApi } from '../../../../../../../api'
import { RootStore, StoreProvider } from '../../../../../../../stores'
import OutwaysWithoutAccessSection from './index'

jest.mock('../../../../../../../components/Modal')

test('Outways without access section', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'outway-1',
        publicKeyFingerprint: 'public-key-fingerprint-1',
        publicKeyPEM: 'public-key-pem-1',
      },
      {
        name: 'outway-2',
        publicKeyFingerprint: 'public-key-fingerprint-1',
        publicKeyPEM: 'public-key-pem-2',
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

  jest.spyOn(service, 'requestAccess').mockResolvedValue()

  const onShowConfirmRequestAccessModalHandler = jest.fn()
  const onHideConfirmRequestAccessModalHandler = jest.fn()

  const { container } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OutwaysWithoutAccessSection
          service={service}
          onShowConfirmRequestAccessModalHandler={
            onShowConfirmRequestAccessModalHandler
          }
          onHideConfirmRequestAccessModalHandler={
            onHideConfirmRequestAccessModalHandler
          }
        />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(container).toHaveTextContent(/None/)

  await rootStore.outwayStore.fetchAll()

  fireEvent.click(screen.getByText(/Outways without access/i))

  expect(screen.getByText('outway-1')).toBeInTheDocument()
  expect(screen.getByText('outway-2')).toBeInTheDocument()

  fireEvent.click(screen.getByText('Request access'))

  expect(onShowConfirmRequestAccessModalHandler).toHaveBeenCalledTimes(1)

  fireEvent.click(await screen.findByText('Send'))

  await waitFor(() => {
    expect(service.requestAccess).toHaveBeenCalledWith('public-key-pem-1')
  })
  expect(onHideConfirmRequestAccessModalHandler).toHaveBeenCalledTimes(1)
})
