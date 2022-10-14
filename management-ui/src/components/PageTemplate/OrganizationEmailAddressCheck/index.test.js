// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementServiceApi } from '../../../api'
import OrganizationEmailAddressCheck from './index'

test('no organization email address', async () => {
  const managementApiClient = new ManagementServiceApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  rootStore.applicationStore.updateOrganizationEmailAddress({
    isOrganizationEmailAddressSet: false,
  })

  const { findByText } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OrganizationEmailAddressCheck />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(
    await findByText('Please set an organization email address.'),
  ).toBeInTheDocument()
})

test('organization email address is set', () => {
  const managementApiClient = new ManagementServiceApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  rootStore.applicationStore.updateOrganizationEmailAddress({
    isOrganizationEmailAddressSet: true,
  })

  const { queryByText } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OrganizationEmailAddressCheck />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(
    queryByText('Please set an organization email address.'),
  ).not.toBeInTheDocument()
})
