// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, waitFor, within } from '@testing-library/react'
import { renderWithProviders } from '../../../../../test-utils'
import { ACCESS_REQUEST_STATES } from '../../../../../stores/models/OutgoingAccessRequestModel'
import DirectoryDetailView from './index'

jest.mock('../../../../../components/Modal')

test('can request access', async () => {
  const service = {
    id: 'Test Organization/Test Service',
    organizationName: 'Test Organization',
    serviceName: 'Test Service',
    state: 'degraded',
    apiSpecificationType: 'API',
    latestAccessRequest: null,
    requestAccess: jest.fn(),
    retryRequestAccess: () => {},
  }

  const { findByText, findByRole } = renderWithProviders(
    <DirectoryDetailView service={service} />,
  )

  const requestAccessButton = await findByText('Request access')
  fireEvent.click(requestAccessButton)

  const dialog = await findByRole('dialog')
  const okButton = within(dialog).getByText('Send')

  fireEvent.click(okButton)
  await waitFor(() => expect(service.requestAccess).toHaveBeenCalled())
})

test('display stacktrace when requesting access failed', () => {
  const service = {
    id: 'my-service',
    organizationName: 'Test Organization',
    serviceName: 'Test Service',
    latestAccessRequest: {
      id: 'my-latest-access-request',
      organizationName: 'organization',
      serviceName: 'service',
      state: ACCESS_REQUEST_STATES.FAILED,
      createdAt: new Date('2020-06-30T08:31:41.106Z'),
      updatedAt: new Date('2020-06-30T08:31:41.106Z'),
      errorDetails: {
        cause: 'Something went wrong',
        stackTrace: ['Go main panic'],
      },
    },
    requestAccess: () => {},
    retryRequestAccess: () => {},
  }

  const { getAllByText, getByText, getByTestId } = renderWithProviders(
    <DirectoryDetailView service={service} />,
  )

  const failedMessages = getAllByText('Request could not be sent')
  const stacktraceButton = getByText('Show stacktrace')

  expect(failedMessages).toHaveLength(2)

  fireEvent.click(stacktraceButton)

  expect(getByTestId('stacktrace')).toBeVisible()
  expect(getByText('Go main panic')).toBeInTheDocument()
})
