// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import OrganizationInwayCheck from './index'

const warningTextKey =
  'Access requests can not be received. Please specify which inway should handle access requests.'

test('fetches settings if not set in store', () => {
  const getSettings = jest
    .fn()
    .mockResolvedValue({ organizationInway: 'inway' })

  const rootStore = new RootStore()
  rootStore.servicesStore.isInitiallyFetched = true
  rootStore.applicationStore.isOrganizationInwaySet = null

  renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck getSettings={getSettings} />
      </StoreProvider>
    </Router>,
  )

  expect(getSettings).toHaveBeenCalled()
})

test('shows warning message when inway is not set and there are services', async () => {
  const getSettings = jest.fn().mockResolvedValue({ organizationInway: false })

  const rootStore = new RootStore()
  rootStore.servicesStore.isInitiallyFetched = true
  rootStore.servicesStore.services.push({ serviceName: 'service' })

  const { findByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck getSettings={getSettings} />
      </StoreProvider>
    </Router>,
  )

  expect(await findByText(warningTextKey)).toBeInTheDocument()
})

test('does not show warning message when inway is not set and services are not set', () => {
  const getSettings = jest.fn().mockResolvedValue({ organizationInway: false })

  const rootStore = new RootStore()
  rootStore.servicesStore.isInitiallyFetched = true

  const { queryByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck getSettings={getSettings} />
      </StoreProvider>
    </Router>,
  )

  expect(queryByText(warningTextKey)).not.toBeInTheDocument()
})

test('does not show warning message when inway is set and there are services', () => {
  const getSettings = jest.fn()

  const rootStore = new RootStore()
  rootStore.applicationStore.isOrganizationInwaySet = true
  rootStore.servicesStore.isInitiallyFetched = true
  rootStore.servicesStore.services.push({ serviceName: 'service' })

  const { queryByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck getSettings={getSettings} />
      </StoreProvider>
    </Router>,
  )

  expect(getSettings).not.toHaveBeenCalled()
  expect(queryByText(warningTextKey)).not.toBeInTheDocument()
})
