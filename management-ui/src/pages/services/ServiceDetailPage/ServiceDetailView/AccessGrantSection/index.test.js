// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, screen, waitForElementToBeRemoved } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import {
  fireEvent,
  renderWithProviders,
  renderWithAllProviders,
} from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { ManagementServiceApi } from '../../../../../api'
import { INTERVAL } from '../../../../../hooks/use-polling'
import AccessGrantSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  jest.useRealTimers()
})

test('polling with access grant section collapsed', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListServices = jest
    .fn()
    .mockResolvedValue({
      services: [
        {
          name: 'service-a',
          organization: {
            name: 'organization-a',
            serialNumber: '00000000000000000001',
          },
        },
      ],
    })

  managementApiClient.managementServiceListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [
        {
          id: '1',
          serviceName: 'service-a',
          organization: {
            name: 'organization-a',
            serialNumber: '00000000000000000001',
          },
          createdAt: new Date(),
          updatedAt: new Date(),
        },
      ],
    })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.servicesStore.fetchAll()

  const { getByTestId, findByTestId } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessGrantSection
        service={rootStore.servicesStore.getByName('service-a')}
      />
    </StoreProvider>,
  )

  expect(getByTestId('amount')).toHaveTextContent('0')

  jest.advanceTimersByTime(INTERVAL)

  const amountAccented = await findByTestId('amount')
  expect(amountAccented).toHaveTextContent('1')
})

test('polling with access grant section expanded', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceGetService = jest
    .fn()
    .mockResolvedValue({
      name: 'service-a',
    })

  managementApiClient.managementServiceListAccessGrantsForService = jest
    .fn()
    .mockResolvedValueOnce({
      accessGrants: [
        {
          id: '1',
          serviceName: 'service-a',
          organization: {
            name: 'organization-a',
            serialNumber: '00000000000000000001',
          },
          createdAt: new Date(),
          updatedAt: new Date(),
        },
      ],
    })
    .mockResolvedValue({
      accessGrants: [],
    })

  managementApiClient.managementServiceListIncomingAccessRequests = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [],
    })

  const rootStore = new RootStore({ managementApiClient })
  await act(async () => {
    await rootStore.servicesStore.fetch({ name: 'service-a' })
  })
  const service = rootStore.servicesStore.getByName('service-a')

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <MemoryRouter>
        <AccessGrantSection service={service} />
      </MemoryRouter>
    </StoreProvider>,
  )

  const toggler = screen.getByText(/Organizations with access/i)

  fireEvent.click(toggler)

  expect(screen.getByText('organization-a')).toBeInTheDocument()
  expect(screen.queryByText('Show updates')).not.toBeInTheDocument()

  act(() => {
    jest.advanceTimersByTime(INTERVAL)
  })

  expect(await screen.findByText('organization-a')).toBeInTheDocument()

  expect(screen.getByText('Show updates')).toBeInTheDocument()

  fireEvent.click(screen.getByText('Show updates'))

  await waitForElementToBeRemoved(() => screen.getByText('organization-a'))
  expect(screen.queryByText('Show updates')).not.toBeInTheDocument()
})
