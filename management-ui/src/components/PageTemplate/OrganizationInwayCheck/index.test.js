// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementServiceApi } from '../../../api'
import OrganizationInwayCheck from './index'

test('providing services but no organization inway', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClientWithService(managementApiClient)

  const rootStore = new RootStore({
    managementApiClient,
  })

  rootStore.servicesStore.fetch({ name: 'service ' })

  rootStore.applicationStore.updateOrganizationInway({
    isOrganizationInwaySet: false,
  })

  const { findByText } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(
    await findByText(
      'Please select an organization inway. At the moment access requests can not be received and outgoing orders can not be retrieved by other organizations.',
    ),
  ).toBeInTheDocument()
})

test('having outgoing orders but no organization Inway', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClientWithOutgoingOrder(managementApiClient)

  const rootStore = new RootStore({
    managementApiClient,
  })

  rootStore.orderStore.fetchOutgoing()

  rootStore.applicationStore.updateOrganizationInway({
    isOrganizationInwaySet: false,
  })

  const { findByText } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(
    await findByText(
      'Please select an organization inway. At the moment access requests can not be received and outgoing orders can not be retrieved by other organizations.',
    ),
  ).toBeInTheDocument()
})

test('organization inway is set with services and outgoing orders present', () => {
  const managementApiClient = new ManagementServiceApi()

  const rootStore = new RootStore({
    managementApiClient,
  })

  managementApiClientWithOutgoingOrder(managementApiClient)
  managementApiClientWithService(managementApiClient)

  rootStore.orderStore.fetchOutgoing()
  rootStore.servicesStore.fetch({ name: 'service ' })

  rootStore.applicationStore.updateOrganizationInway({
    isOrganizationInwaySet: false,
  })

  const { queryByText } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <OrganizationInwayCheck />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(
    queryByText(
      'Please select an organization inway. At the moment access requests can not be received and outgoing orders can not be retrieved by other organizations.',
    ),
  ).not.toBeInTheDocument()
})

function managementApiClientWithService(managementApiClient) {
  managementApiClient.managementServiceGetService = jest
    .fn()
    .mockResolvedValue({
      name: 'service',
    })

  managementApiClient.managementServiceListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [],
    })

  managementApiClient.managementServiceListIncomingAccessRequests = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [],
    })
}

function managementApiClientWithOutgoingOrder(managementApiClient) {
  managementApiClient.managementServiceListOutgoingOrders = jest
    .fn()
    .mockResolvedValue({
      orders: [
        {
          reference: 'reference',
          delegatee: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
        },
      ],
    })
}
