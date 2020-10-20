// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders, fireEvent } from '../../../../../../test-utils'
import AccessRequestSection from './index'

test('Should throw error when state is unknown', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})

  const latestAccessRequest = {
    id: 'id',
    state: '?',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  expect(() =>
    renderWithProviders(
      <AccessRequestSection latestAccessRequest={latestAccessRequest} />,
    ),
  ).toThrow()

  errorSpy.mockRestore()
})

test('Correctly renders when there is no access request', () => {
  const requestAccessSpy = jest.fn()
  const { getByText } = renderWithProviders(
    <AccessRequestSection
      latestAccessRequest={undefined}
      requestAccess={requestAccessSpy}
    />,
  )

  expect(getByText('You have no access')).toBeInTheDocument()

  const button = getByText('Request Access')
  fireEvent.click(button)

  expect(requestAccessSpy).toHaveBeenCalled()
})

test('Correctly renders the access request states', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'FAILED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }

  const { getByText, rerender } = renderWithProviders(
    <AccessRequestSection latestAccessRequest={latestAccessRequest} />,
  )
  expect(getByText('Request could not be sent')).toBeInTheDocument()

  rerender(
    <AccessRequestSection
      latestAccessRequest={{ ...latestAccessRequest, state: 'CREATED' }}
    />,
  )
  expect(getByText('Sending request…')).toBeInTheDocument()

  rerender(
    <AccessRequestSection
      latestAccessRequest={{ ...latestAccessRequest, state: 'RECEIVED' }}
    />,
  )
  expect(getByText('Access requested')).toBeInTheDocument()

  rerender(
    <AccessRequestSection
      latestAccessRequest={{ ...latestAccessRequest, state: 'APPROVED' }}
    />,
  )
  expect(getByText('You have access')).toBeInTheDocument()

  rerender(
    <AccessRequestSection
      latestAccessRequest={{ ...latestAccessRequest, state: 'REJECTED' }}
    />,
  )
  expect(getByText('Access request rejected')).toBeInTheDocument()
})
