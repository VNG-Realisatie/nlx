// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Routes, MemoryRouter } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../test-utils'
import { DirectoryServiceApi, ManagementServiceApi } from '../../../api'
import { RootStore, StoreProvider } from '../../../stores'
import { SERVICE_STATE_DEGRADED } from '../../../components/StateIndicator'
import DirectoryDetailPage from './index'

jest.mock('./components/DirectoryDetailView', () => () => (
  <div data-testid="directory-service-details" />
))

test('display directory service details', async () => {
  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [
        {
          organization: {
            serialNumber: '00000000000000000001',
            name: 'Test Organization',
          },
          serviceName: 'Test Service',
          state: SERVICE_STATE_DEGRADED,
          apiSpecificationType: 'API',
          latestAccessRequest: null,
        },
      ],
    })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [],
    })

  const rootStore = new RootStore({
    directoryApiClient,
    managementApiClient,
  })

  await rootStore.directoryServicesStore.fetchAll()

  renderWithProviders(
    <MemoryRouter initialEntries={['/00000000000000000001/Test Service']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route
            path=":organizationSerialNumber/:serviceName"
            element={<DirectoryDetailPage />}
          />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(await screen.findByText('Test Organization')).toBeInTheDocument()
  expect(screen.getByText('Test Service')).toBeInTheDocument()
  expect(screen.getByText('state-degraded.svg')).toBeInTheDocument()
  expect(screen.getByTestId('directory-service-details')).toBeInTheDocument()

  directoryApiClient.directoryServiceGetOrganizationService = jest
    .fn()
    .mockRejectedValue({
      status: 404,
    })

  await rootStore.directoryServicesStore.fetch(
    '00000000000000000001',
    'Test Service',
  )

  expect(screen.getByText('Failed to load the service')).toBeInTheDocument()
})

test('service does not exist', () => {
  const directoryApiClient = new DirectoryServiceApi()

  directoryApiClient.directoryServiceGetOrganizationService = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListOutways = jest
    .fn()
    .mockResolvedValue({
      outways: [],
    })

  const rootStore = new RootStore({
    directoryApiClient,
    managementApiClient,
  })

  renderWithProviders(
    <MemoryRouter initialEntries={['/00000000000000000001/Test Service']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route
            path=":organizationSerialNumber/:serviceName"
            element={<DirectoryDetailPage />}
          />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )

  const message = screen.getByTestId('error-message')
  expect(message).toBeInTheDocument()
  expect(message.textContent).toBe('Failed to load the service')

  expect(screen.getByText('Test Service')).toBeInTheDocument()
  expect(screen.queryByText('organization')).toBeNull()

  const closeButton = screen.getByTestId('close-button')
  expect(closeButton).toBeInTheDocument()
})
