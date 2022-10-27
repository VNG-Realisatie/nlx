// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { fireEvent, screen, waitFor, within } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../../../../test-utils'
import { ACCESS_REQUEST_STATES } from '../../../../../stores/models/OutgoingAccessRequestModel'
import { RootStore, StoreProvider } from '../../../../../stores'
import { DirectoryServiceApi, ManagementServiceApi } from '../../../../../api'
import DirectoryDetailView from './index'

jest.mock('../../../../../components/Modal')

test('can request access', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'outway-1',
          publicKeyFingerprint: 'public-key-fingerprint-1',
        },
      ],
    })

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceGetOrganizationService = jest
    .fn()
    .mockResolvedValue({
      directoryService: {
        id: 'Test Organization/Test Service',
        organization: {
          serialNumber: '00000000000000000001',
          name: 'Test Organization',
        },
        serviceName: 'Test Service',
        state: 'degraded',
        apiSpecificationType: 'API',
        accessStates: [],
      },
    })

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  const serviceModel = await rootStore.directoryServicesStore.fetch(
    '00000000000000000001',
    'Test organization',
  )

  jest.spyOn(serviceModel, 'requestAccess').mockResolvedValue()

  await rootStore.outwayStore.fetchAll()

  renderWithAllProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <DirectoryDetailView service={serviceModel} />
      </StoreProvider>
    </MemoryRouter>,
  )

  const outwaysWithoutAccessSection = await screen.findByText(
    /Outways without access/i,
  )

  fireEvent.click(outwaysWithoutAccessSection)

  const requestAccessButton = await screen.findByText('Request access')
  fireEvent.click(requestAccessButton)

  const dialog = await screen.findByRole('dialog')
  const okButton = within(dialog).getByText('Send')

  fireEvent.click(okButton)
  await waitFor(() => expect(serviceModel.requestAccess).toHaveBeenCalled())
})

test('display error when existing access request failed', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [
        {
          name: 'outway-1',
          publicKeyFingerprint: 'public-key-fingerprint',
        },
      ],
    })

  managementApiClient.managementServiceSynchronizeOutgoingAccessRequests = jest
    .fn()
    .mockResolvedValue({})

  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceGetOrganizationService = jest
    .fn()
    .mockResolvedValue({
      directoryService: {
        id: 'Test Organization/Test Service',
        organization: {
          serialNumber: '00000000000000000001',
          name: 'Test Organization',
        },
        serviceName: 'Test Service',
        state: 'degraded',
        apiSpecificationType: 'API',
        accessStates: [
          {
            accessRequest: {
              id: 'my-latest-access-request',
              organization: {
                serialNumber: '00000000000000000002',
                name: 'organization',
              },
              serviceName: 'service',
              state: ACCESS_REQUEST_STATES.FAILED,
              createdAt: new Date('2020-06-30T08:31:41.106Z'),
              updatedAt: new Date('2020-06-30T08:31:41.106Z'),
              errorDetails: {
                cause: 'Something went wrong',
              },
              publicKeyFingerprint: 'public-key-fingerprint',
            },
            accessProof: null,
          },
        ],
      },
    })

  const rootStore = new RootStore({
    managementApiClient,
    directoryApiClient,
  })

  await rootStore.outwayStore.fetchAll()

  const serviceModel = await rootStore.directoryServicesStore.fetch(
    '00000000000000000001',
    'Test organization',
  )

  renderWithAllProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <DirectoryDetailView service={serviceModel} />
      </StoreProvider>
    </MemoryRouter>,
  )

  fireEvent.click(await screen.findByText(/Outways without access/i))

  expect(screen.getByRole('alert')).toBeInTheDocument()
  expect(screen.getByText('Something went wrong')).toBeInTheDocument()
})
