// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../../../test-utils'
import AccessRequestMessage from './index'

test('Should throw error when state is unknown', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})

  expect(() => renderWithProviders(<AccessRequestMessage />)).toThrow()

  errorSpy.mockRestore()
})

test('Correctly renders the access request states', () => {
  const latestAccessRequest = {
    id: 'id',
    organizationName: 'organization',
    serviceName: 'service',
    state: 'FAILED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }

  const { getByText, getByTitle, rerender } = renderWithProviders(
    <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
  )
  expect(getByText('Request could not be sent')).toBeInTheDocument()

  rerender(
    <AccessRequestMessage
      latestAccessRequest={{ ...latestAccessRequest, state: 'CREATED' }}
    />,
  )
  expect(getByText('Sending request')).toBeInTheDocument()

  rerender(
    <AccessRequestMessage
      latestAccessRequest={{ ...latestAccessRequest, state: 'RECEIVED' }}
    />,
  )
  expect(getByText('Requested')).toBeInTheDocument()

  rerender(
    <AccessRequestMessage
      latestAccessRequest={{ ...latestAccessRequest, state: 'APPROVED' }}
    />,
  )
  expect(getByText('check.svg')).toBeInTheDocument()
  expect(getByTitle('Approved')).toBeInTheDocument()

  rerender(
    <AccessRequestMessage
      latestAccessRequest={{ ...latestAccessRequest, state: 'REJECTED' }}
    />,
  )
  expect(getByText('Rejected')).toBeInTheDocument()
})
