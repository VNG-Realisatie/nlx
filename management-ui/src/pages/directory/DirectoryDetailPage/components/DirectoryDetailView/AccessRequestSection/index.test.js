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

test('Correctly renders FAILED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'FAILED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestSection latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('Request could not be sent')).toBeInTheDocument()
})

test('Correctly renders CREATED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'CREATED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestSection latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('Sending request...')).toBeInTheDocument()
})

test('Correctly renders RECEIVED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'RECEIVED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestSection latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('Access requested')).toBeInTheDocument()
})

test('Correctly renders APPROVED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'APPROVED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestSection latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('You have access')).toBeInTheDocument()
})

test('Correctly renders REJECTED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'REJECTED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestSection latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('Access request rejected')).toBeInTheDocument()
})
