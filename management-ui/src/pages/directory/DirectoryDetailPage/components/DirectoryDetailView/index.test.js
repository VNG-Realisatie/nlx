// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { makeAutoObservable } from 'mobx'
import { Route, StaticRouter as Router } from 'react-router-dom'
import { fireEvent, within } from '@testing-library/react'
import { renderWithProviders } from '../../../../../test-utils'
import { clickConfirmButtonAndAssert } from '../../../../../components/ConfirmationModal/testUtils'
import DirectoryDetailPage from '../../index'

let service

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

test('can request access', async () => {
  const requestAccessFn = jest.fn()
  service.requestAccess = requestAccessFn

  const { findByText, findByRole } = renderWithProviders(
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage service={service} />
      </Route>
    </Router>,
  )

  const requestAccessButton = await findByText('Request access')
  fireEvent.click(requestAccessButton)

  const dialog = await findByRole('dialog')
  const okButton = within(dialog).getByText('Send')
  await clickConfirmButtonAndAssert(okButton, () =>
    expect(requestAccessFn).toHaveBeenCalled(),
  )
})
