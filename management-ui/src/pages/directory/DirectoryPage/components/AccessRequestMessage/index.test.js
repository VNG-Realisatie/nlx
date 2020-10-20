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

test('Correctly renders FAILED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'FAILED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
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
    <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('Sending request')).toBeInTheDocument()
})

test('Correctly renders RECEIVED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'RECEIVED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('Requested')).toBeInTheDocument()
})

test('Correctly renders APPROVED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'APPROVED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText, getByTitle } = renderWithProviders(
    <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('check.svg')).toBeInTheDocument()
  expect(getByTitle('Approved')).toBeInTheDocument()
})

test('Correctly renders REJECTED state', () => {
  const latestAccessRequest = {
    id: 'id',
    state: 'REJECTED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:21Z',
  }
  const { getByText } = renderWithProviders(
    <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
  )

  expect(getByText('Rejected')).toBeInTheDocument()
})
