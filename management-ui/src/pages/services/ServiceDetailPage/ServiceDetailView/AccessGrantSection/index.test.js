// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, waitForElementToBeRemoved } from '@testing-library/react'
import {
  fireEvent,
  renderWithProviders,
  renderWithAllProviders,
} from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { ManagementApi } from '../../../../../api'
import { INTERVAL } from '../../../../../hooks/use-polling'
import AccessGrantSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  jest.useRealTimers()
})

test('polling with access grant section collapsed', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [
      {
        name: 'service-a',
      },
    ],
  })

  managementApiClient.managementListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [
        {
          id: '1',
          serviceName: 'service-a',
          organization: {
            name: 'organization-a',
            serviceName: '00000000000000000001',
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
        service={rootStore.servicesStore.getService('service-a')}
      />
    </StoreProvider>,
  )

  expect(getByTestId('amount')).toHaveTextContent('0')

  jest.advanceTimersByTime(INTERVAL)

  const amountAccented = await findByTestId('amount')
  expect(amountAccented).toHaveTextContent('1')
})

test('polling with access grant section expanded', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetService = jest.fn().mockResolvedValue({
    name: 'service-a',
  })

  managementApiClient.managementListAccessGrantsForService = jest
    .fn()
    .mockResolvedValueOnce({
      accessGrants: [
        {
          id: '1',
          serviceName: 'service-a',
          organization: {
            name: 'organization-a',
            serviceName: '00000000000000000001',
          },
          createdAt: new Date(),
          updatedAt: new Date(),
        },
      ],
    })
    .mockResolvedValue({
      accessGrants: [],
    })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [],
    })

  const rootStore = new RootStore({ managementApiClient })
  await act(async () => {
    await rootStore.servicesStore.fetch({ name: 'service-a' })
  })
  const service = rootStore.servicesStore.getService('service-a')

  const { getByText, queryByText, findByText } = renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessGrantSection service={service} />
    </StoreProvider>,
  )

  const toggler = getByText(/Organizations with access/i)

  fireEvent.click(toggler)

  expect(getByText('organization-a')).toBeInTheDocument()
  expect(queryByText('Show updates')).not.toBeInTheDocument()

  act(() => {
    jest.advanceTimersByTime(INTERVAL)
  })

  expect(await findByText('organization-a')).toBeInTheDocument()

  expect(getByText('Show updates')).toBeInTheDocument()

  fireEvent.click(getByText('Show updates'))

  await waitForElementToBeRemoved(() => getByText('organization-a'))
  expect(queryByText('Show updates')).not.toBeInTheDocument()
})
