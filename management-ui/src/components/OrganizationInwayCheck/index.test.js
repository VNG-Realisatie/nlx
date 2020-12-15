// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../test-utils'
import { useApplicationStore, useServicesStore } from '../../hooks/use-stores'
import OrganizationInwayCheck from './index'

jest.mock('../../hooks/use-stores', () => ({
  useApplicationStore: jest.fn(),
  useServicesStore: jest.fn(),
}))

test('fetches settings if not set in store', () => {
  const getSettings = jest
    .fn()
    .mockResolvedValue({ organizationInway: 'inway' })
  const updateFn = jest.fn()

  useApplicationStore.mockImplementation(() => ({
    isOrganizationInwaySet: null,
    update: updateFn,
  }))

  useServicesStore.mockImplementation(() => ({
    services: [{ serviceName: 'service' }],
  }))

  renderWithProviders(
    <Router>
      <OrganizationInwayCheck getSettings={getSettings} />
    </Router>,
  )

  expect(getSettings).toHaveBeenCalled()
  // TODO: how to await this?
  // expect(updateFn).toHaveBeenCalledWith({ isOrganizationInwaySet: 'inway' })
})

test('shows warning message when inway is not set and there are services', () => {
  const getSettings = jest.fn()

  useApplicationStore.mockImplementation(() => ({
    isOrganizationInwaySet: false,
    update: jest.fn(),
  }))

  useServicesStore.mockImplementation(() => ({
    services: [{ serviceName: 'service' }],
  }))

  const { getByText } = renderWithProviders(
    <Router>
      <OrganizationInwayCheck getSettings={getSettings} />
    </Router>,
  )

  expect(
    getByText(
      'Access requests can not be received. Set which inway handles access requests.',
    ),
  ).toBeInTheDocument()
})

test('does not show warning message when inway is not set and services are not set', () => {
  const getSettings = jest.fn()

  useApplicationStore.mockImplementation(() => ({
    isOrganizationInwaySet: false,
    update: jest.fn(),
  }))

  useServicesStore.mockImplementation(() => ({
    services: [],
  }))

  const { queryByText } = renderWithProviders(
    <Router>
      <OrganizationInwayCheck getSettings={getSettings} />
    </Router>,
  )

  expect(
    queryByText(
      'Access requests can not be received. Set which inway handles access requests.',
    ),
  ).not.toBeInTheDocument()
})
