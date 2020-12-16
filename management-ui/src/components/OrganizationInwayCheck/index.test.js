// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import { RootStore, StoreProvider } from '../../stores'
import OrganizationInwayCheck from './index'

test('fetches settings if not set in store', () => {
  const getSettings = jest
    .fn()
    .mockResolvedValue({ organizationInway: 'inway' })

  const rootStore = new RootStore()
  rootStore.servicesStore.isInitiallyFetched = true
  rootStore.applicationStore.isOrganizationInwaySet = null

  renderWithProviders(
    <Router>
      <StoreProvider store={rootStore}>
        <OrganizationInwayCheck getSettings={getSettings} />
      </StoreProvider>
    </Router>,
  )

  expect(getSettings).toHaveBeenCalled()
})

test('shows warning message when inway is not set and there are services', async () => {
  const getSettings = jest.fn().mockResolvedValue({ organizationInway: null })

  const rootStore = new RootStore()
  rootStore.servicesStore.isInitiallyFetched = true
  rootStore.servicesStore.services.push({ serviceName: 'service' })

  const { findByText } = renderWithProviders(
    <Router>
      <StoreProvider store={rootStore}>
        <OrganizationInwayCheck getSettings={getSettings} />
      </StoreProvider>
    </Router>,
  )

  expect(
    await findByText(
      'Access requests can not be received. Set which inway handles access requests.',
    ),
  ).toBeInTheDocument()
})

test('does not show warning message when inway is not set and services are not set', () => {
  const getSettings = jest.fn().mockResolvedValue({ organizationInway: null })

  const rootStore = new RootStore()
  rootStore.servicesStore.isInitiallyFetched = true

  const { queryByText } = renderWithProviders(
    <Router>
      <StoreProvider store={rootStore}>
        <OrganizationInwayCheck getSettings={getSettings} />
      </StoreProvider>
    </Router>,
  )

  expect(
    queryByText(
      'Access requests can not be received. Set which inway handles access requests.',
    ),
  ).not.toBeInTheDocument()
})
