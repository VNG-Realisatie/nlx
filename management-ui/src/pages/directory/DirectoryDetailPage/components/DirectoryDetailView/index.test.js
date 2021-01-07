// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { makeAutoObservable } from 'mobx'
import { Route, StaticRouter as Router } from 'react-router-dom'
import { fireEvent, renderWithProviders } from '../../../../../test-utils'
import DirectoryDetailPage from '../../index'

let service

global.confirm = jest.fn(() => true)

beforeEach(() => {
  service = makeAutoObservable({
    id: 'Test Organization/Test Service',
    organizationName: 'Test Organization',
    serviceName: 'Test Service',
    state: 'degraded',
    apiSpecificationType: 'API',
    latestAccessRequest: null,
    fetch: jest.fn(),
    requestAccess: jest.fn(),
    retryRequestAccess: jest.fn(),
  })
})

test('A service with failed latestAccessRequest', () => {
  service.latestAccessRequest = {
    id: 'string',
    organizationName: 'organization',
    serviceName: 'service',
    state: 'FAILED',
    createdAt: new Date('2020-06-30T08:31:41.106Z'),
    updatedAt: new Date('2020-06-30T08:31:41.106Z'),
    errorDetails: {
      cause: 'Something went wrong',
      stackTrace: ['Go main panic'],
    },
  }

  const { getAllByText, getByText, getByTestId } = renderWithProviders(
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage service={service} />
      </Route>
    </Router>,
  )

  const failedMessages = getAllByText('Request could not be sent')
  const stacktraceButton = getByText('Show stacktrace')

  expect(failedMessages).toHaveLength(2)

  fireEvent.click(stacktraceButton)

  expect(getByTestId('stacktrace')).toBeVisible()
  expect(getByText('Go main panic')).toBeInTheDocument()
})
