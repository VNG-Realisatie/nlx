// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, waitForElementToBeRemoved } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import {
  fireEvent,
  renderWithProviders,
  renderWithAllProviders,
} from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { STATES } from '../../../../../stores/models/IncomingAccessRequestModel'
import { ManagementApi } from '../../../../../api'
import { INTERVAL } from '../../../../../hooks/use-polling'
import AccessRequestsSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('polling with access request section collapsed', async () => {
  jest.useFakeTimers()
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [
      {
        name: 'service-a',
      },
    ],
  })

  managementApiClient.managementListIncomingAccessRequests = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [
        {
          id: '1',
          serviceName: 'service-a',
          state: STATES.RECEIVED,
          createdAt: new Date(),
          updatedAt: new Date(),
          publicKeyFingerprint: 'public-key-fingerprint',
        },
      ],
    })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.servicesStore.fetchAll()

  const { getByTestId, findByTestId } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessRequestsSection
        service={rootStore.servicesStore.getByName('service-a')}
        onApproveOrRejectCallbackHandler={() => {}}
      />
    </StoreProvider>,
  )

  expect(getByTestId('amount')).toHaveTextContent('0')

  jest.advanceTimersByTime(INTERVAL)

  const amountAccented = await findByTestId('amount')

  expect(amountAccented).toHaveTextContent('1')

  jest.useRealTimers()
})

test('polling with access request section expanded', async () => {
  jest.useFakeTimers()

  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetService = jest.fn().mockResolvedValue({
    name: 'service-a',
  })

  managementApiClient.managementListIncomingAccessRequests = jest
    .fn()
    .mockResolvedValueOnce({
      accessRequests: [
        {
          id: '1',
          serviceName: 'service-a',
          organization: {
            serialNumber: '00000000000000000001',
            name: 'organization-a',
          },
          state: STATES.RECEIVED,
          createdAt: new Date(),
          updatedAt: new Date(),
          publicKeyFingerprint: 'public-key-fingerprint',
        },
      ],
    })
    .mockResolvedValue({
      accessRequests: [],
    })

  managementApiClient.managementListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [],
    })

  const rootStore = new RootStore({ managementApiClient })
  await act(async () => {
    await rootStore.servicesStore.fetch({ name: 'service-a' })
  })
  const service = rootStore.servicesStore.getByName('service-a')

  const { getByText, queryByText, findByText } = renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <MemoryRouter>
        <AccessRequestsSection
          service={service}
          onApproveOrRejectCallbackHandler={() => {}}
        />
      </MemoryRouter>
    </StoreProvider>,
  )

  const toggler = getByText(/Access requests/i)

  fireEvent.click(toggler)

  expect(getByText(/organization-a/)).toBeInTheDocument()
  expect(getByText(/OIN serialNumber/)).toBeInTheDocument()
  expect(
    getByText(/Public Key Fingerprint publicKeyFingerprint/),
  ).toBeInTheDocument()
  expect(queryByText('Show updates')).not.toBeInTheDocument()

  act(() => {
    jest.advanceTimersByTime(INTERVAL)
  })

  expect(await findByText('organization-a')).toBeInTheDocument()

  expect(getByText('Show updates')).toBeInTheDocument()

  fireEvent.click(getByText('Show updates'))

  await waitForElementToBeRemoved(() => getByText('organization-a'))
  expect(queryByText('Show updates')).not.toBeInTheDocument()

  jest.useRealTimers()
})
