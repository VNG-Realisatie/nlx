// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import OrganizationInwayCheck from './index'

test('shows warning message when inway is not set and there are services', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetService = jest.fn().mockResolvedValue({
    name: 'service',
  })

  managementApiClient.managementListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [],
    })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [],
    })

  const rootStore = new RootStore({
    managementApiClient,
  })
  const applicationStore = rootStore.applicationStore
  const servicesStore = rootStore.servicesStore

  applicationStore.update({
    isOrganizationInwaySet: false,
  })

  servicesStore.fetch({ name: 'service ' })

  const { findByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck />
      </StoreProvider>
    </Router>,
  )

  expect(
    await findByText(
      'Access requests can not be received. Please specify which inway should handle access requests.',
    ),
  ).toBeInTheDocument()
})

test('does not show warning message when inway is not set and services are not set', () => {
  const managementApiClient = new ManagementApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  rootStore.applicationStore.update({
    isOrganizationInwaySet: false,
  })

  const { queryByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck />
      </StoreProvider>
    </Router>,
  )

  expect(
    queryByText(
      'Access requests can not be received. Please specify which inway should handle access requests.',
    ),
  ).not.toBeInTheDocument()
})

test('does not show warning message when inway is set and there are services', () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetService = jest.fn().mockResolvedValue({
    name: 'service',
  })

  managementApiClient.managementListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [],
    })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [],
    })

  const rootStore = new RootStore({
    managementApiClient,
  })
  const applicationStore = rootStore.applicationStore
  const servicesStore = rootStore.servicesStore

  applicationStore.update({
    isOrganizationInwaySet: true,
  })

  servicesStore.fetch({ name: 'service ' })

  const { queryByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck />
      </StoreProvider>
    </Router>,
  )

  expect(
    queryByText(
      'Access requests can not be received. Please specify which inway should handle access requests.',
    ),
  ).not.toBeInTheDocument()
})
