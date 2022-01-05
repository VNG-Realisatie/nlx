// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, screen, waitFor, within } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithProviders } from '../../../../../test-utils'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../../../../stores/models/OutgoingAccessRequestModel'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import { RootStore } from '../../../../../stores'
import { ManagementApi } from '../../../../../api'
import DirectoryDetailView from './index'

jest.mock('../../../../../components/Modal')

test('can request access', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  const serviceModel = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {
      id: 'Test Organization/Test Service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Test Organization',
      },
      serviceName: 'Test Service',
      state: 'degraded',
      apiSpecificationType: 'API',
      latestAccessRequest: null,
    },
  })

  jest.spyOn(serviceModel, 'requestAccess').mockResolvedValue()

  renderWithProviders(<DirectoryDetailView service={serviceModel} />)

  const requestAccessButton = await screen.findByText('Request access')
  fireEvent.click(requestAccessButton)

  const dialog = await screen.findByRole('dialog')
  const okButton = within(dialog).getByText('Send')

  fireEvent.click(okButton)
  await waitFor(() => expect(serviceModel.requestAccess).toHaveBeenCalled())
})

test('display stacktrace when requesting access failed', () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  const serviceModel = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {
      id: 'my-service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Test Organization',
      },
      serviceName: 'Test Service',
    },
    latestAccessRequest: new OutgoingAccessRequestModel({
      outgoingAccessRequestStore: rootStore.outgoingAccessRequestStore,
      accessRequestData: {
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
          stackTrace: ['Go main panic'],
        },
      },
    }),
  })

  renderWithProviders(<DirectoryDetailView service={serviceModel} />)

  const failedMessages = screen.getAllByText('Request could not be sent')
  const stacktraceButton = screen.getByText('Show stacktrace')

  expect(failedMessages).toHaveLength(2)

  fireEvent.click(stacktraceButton)

  expect(screen.getByTestId('stacktrace')).toBeVisible()
  expect(screen.getByText('Go main panic')).toBeInTheDocument()
})
