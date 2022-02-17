// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import {
  renderWithProviders,
  fireEvent,
} from '../../../../../../../../test-utils'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../../../../../../../stores/models/OutgoingAccessRequestModel'
import AccessProofModel from '../../../../../../../../stores/models/AccessProofModel'
import AccessState from './index'

test('No access', () => {
  const requestAccessSpy = jest.fn()

  renderWithProviders(
    <AccessState
      accessRequest={null}
      accessProof={null}
      onRequestAccess={requestAccessSpy}
    />,
  )

  expect(screen.getByText('You have no access')).toBeInTheDocument()

  fireEvent.click(screen.getByText('Request access'))
  expect(requestAccessSpy).toHaveBeenCalled()
})

test('Request access failed', () => {
  const retryRequestAccessSpy = jest.fn()

  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.FAILED,
      errorDetails: {
        cause: 'cause of failed access request',
      },
    },
    outgoingAccessRequestStore: null,
  })

  renderWithProviders(
    <AccessState
      accessRequest={accessRequest}
      accessProof={null}
      onRetryRequestAccess={retryRequestAccessSpy}
    />,
  )

  expect(screen.getByText('Request could not be sent')).toBeInTheDocument()
  expect(screen.getByText('cause of failed access request')).toBeInTheDocument()

  fireEvent.click(screen.getByText('Retry'))
  expect(retryRequestAccessSpy).toHaveBeenCalled()
})

test('Request access created', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.CREATED,
    },
    outgoingAccessRequestStore: null,
  })

  renderWithProviders(
    <AccessState accessRequest={accessRequest} accessProof={null} />,
  )

  expect(screen.getByText('Sending request…')).toBeInTheDocument()
})

test('Request access received', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.RECEIVED,
    },
    outgoingAccessRequestStore: null,
  })

  renderWithProviders(
    <AccessState accessRequest={accessRequest} accessProof={null} />,
  )

  expect(screen.getByText('Access requested')).toBeInTheDocument()
})

test('Has access', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.APPROVED,
    },
    outgoingAccessRequestStore: null,
  })

  const accessProof = new AccessProofModel({
    accessProofData: {
      revokedAt: null,
    },
  })

  renderWithProviders(
    <AccessState accessRequest={accessRequest} accessProof={accessProof} />,
  )

  expect(screen.getByText('You have access')).toBeInTheDocument()
})

test('Access rejected', () => {
  const requestAccessSpy = jest.fn()

  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.REJECTED,
    },
    outgoingAccessRequestStore: null,
  })

  renderWithProviders(
    <AccessState
      accessRequest={accessRequest}
      accessProof={null}
      onRequestAccess={requestAccessSpy}
    />,
  )

  expect(screen.getByText('Access request rejected')).toBeInTheDocument()

  fireEvent.click(screen.getByText('Request access'))
  expect(requestAccessSpy).toHaveBeenCalled()
})

test('Access revoked', () => {
  const requestAccessSpy = jest.fn()

  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.REJECTED,
    },
    outgoingAccessRequestStore: null,
  })

  const accessProof = new AccessProofModel({
    accessProofData: {
      revokedAt: new Date(),
    },
  })

  renderWithProviders(
    <AccessState
      accessRequest={accessRequest}
      accessProof={accessProof}
      onRequestAccess={requestAccessSpy}
    />,
  )

  expect(screen.getByText('Your access was revoked')).toBeInTheDocument()

  fireEvent.click(screen.getByText('Request access'))
  expect(requestAccessSpy).toHaveBeenCalled()
})
