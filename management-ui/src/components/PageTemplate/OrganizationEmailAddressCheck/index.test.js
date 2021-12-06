// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import OrganizationEmailAddressCheck from './index'

test('no organization email address', async () => {
  const managementApiClient = new ManagementApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  rootStore.applicationStore.updateOrganizationEmailAddress({
    isOrganizationEmailAddressSet: false,
  })

  const { findByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationEmailAddressCheck />
      </StoreProvider>
    </Router>,
  )

  expect(
    await findByText('Please set an organization email address.'),
  ).toBeInTheDocument()
})

test('organization email address is set', () => {
  const managementApiClient = new ManagementApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  rootStore.applicationStore.updateOrganizationEmailAddress({
    isOrganizationEmailAddressSet: true,
  })

  const { queryByText } = renderWithProviders(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <OrganizationEmailAddressCheck />
      </StoreProvider>
    </Router>,
  )

  expect(
    queryByText('Please set an organization email address.'),
  ).not.toBeInTheDocument()
})
