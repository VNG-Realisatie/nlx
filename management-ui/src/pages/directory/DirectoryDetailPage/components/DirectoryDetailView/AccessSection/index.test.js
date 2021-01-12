// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders, fireEvent } from '../../../../../../test-utils'
import {
  SHOW_REQUEST_ACCESS,
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
  SHOW_ACCESS_REVOKED,
} from '../../../../directoryServiceAccessState'
import AccessSection from './index'

test('Correctly renders when there is no access', () => {
  const requestAccessSpy = jest.fn()
  const { getByText } = renderWithProviders(
    <AccessSection
      displayState={SHOW_REQUEST_ACCESS}
      latestAccessRequest={null}
      latestAccessProof={null}
      requestAccess={requestAccessSpy}
    />,
  )

  expect(getByText('You have no access')).toBeInTheDocument()

  const button = getByText('Request access')
  fireEvent.click(button)

  expect(requestAccessSpy).toHaveBeenCalled()
})

test('Correctly renders the other states of access', () => {
  // Only `updatedAt` is used by this component, based on displayState
  const latestAccessRequest = {
    id: 'id',
    organizationName: 'foo',
    serviceName: 'bar',
    state: 'CREATED',
    createdAt: new Date('2020-10-01T12:00:00Z'),
    updatedAt: new Date('2020-10-02T12:00:00Z'),
  }

  const { getByText, rerender } = renderWithProviders(
    <AccessSection
      displayState={SHOW_REQUEST_CREATED}
      latestAccessRequest={latestAccessRequest}
      latestAccessProof={null}
      requestAccess={jest.fn()}
    />,
  )
  expect(getByText('Sending request…')).toBeInTheDocument()

  rerender(<AccessSection displayState={SHOW_REQUEST_FAILED} />)
  expect(getByText('Request could not be sent')).toBeInTheDocument()

  rerender(
    <AccessSection
      latestAccessRequest={latestAccessRequest}
      displayState={SHOW_REQUEST_RECEIVED}
    />,
  )
  expect(getByText('Access requested')).toBeInTheDocument()

  rerender(
    <AccessSection
      latestAccessRequest={latestAccessRequest}
      displayState={SHOW_REQUEST_REJECTED}
    />,
  )
  expect(getByText('Access request rejected')).toBeInTheDocument()
  expect(getByText('Request access')).toBeInTheDocument()

  // accessProof.createdAt required in this displayState
  rerender(
    <AccessSection
      displayState={SHOW_HAS_ACCESS}
      latestAccessRequest={latestAccessRequest}
      latestAccessProof={{
        id: 'id',
        organizationName: 'foo',
        serviceName: 'bar',
        createdAt: new Date('2020-10-02T12:01:00Z'),
        revokedAt: null,
      }}
    />,
  )
  expect(getByText('You have access')).toBeInTheDocument()

  rerender(
    <AccessSection
      displayState={SHOW_ACCESS_REVOKED}
      latestAccessRequest={latestAccessRequest}
      latestAccessProof={{
        id: 'id',
        organizationName: 'foo',
        serviceName: 'bar',
        createdAt: new Date('2020-10-02T12:01:00Z'),
        revokedAt: new Date('2020-10-03T12:01:00Z'),
      }}
    />,
  )
  expect(getByText('Your access was revoked')).toBeInTheDocument()
})
