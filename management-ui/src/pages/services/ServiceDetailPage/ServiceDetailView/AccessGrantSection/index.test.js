// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, waitForElementToBeRemoved } from '@testing-library/react'
import { fireEvent, renderWithProviders } from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { ManagementApi } from '../../../../../api'
import AccessGrantSection, { POLLING_INTERVAL } from './index'

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
          organizationName: 'organization-a',
          createdAt: new Date(),
          updatedAt: new Date(),
        },
      ],
    })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.servicesStore.fetchAll()

  const { getByTestId, findByTestId, debug } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessGrantSection
        service={rootStore.servicesStore.getService('service-a')}
      />
    </StoreProvider>,
  )

  expect(getByTestId('amount')).toHaveTextContent('0')

  jest.advanceTimersByTime(POLLING_INTERVAL)

  const amountAccented = await findByTestId('amount')

  debug()
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
          organizationName: 'organization-a',
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

  const { getByText, queryByText, findByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessGrantSection service={service} />
    </StoreProvider>,
  )

  const toggler = getByText(/Organizations with access/i)

  fireEvent.click(toggler)

  expect(getByText('organization-a')).toBeInTheDocument()
  expect(queryByText('New organizations')).not.toBeInTheDocument()

  act(() => {
    jest.advanceTimersByTime(POLLING_INTERVAL)
  })

  expect(await findByText('organization-a')).toBeInTheDocument()

  expect(getByText('New organizations')).toBeInTheDocument()

  fireEvent.click(getByText('New organizations'))

  await waitForElementToBeRemoved(() => getByText('organization-a'))
  expect(queryByText('New organizations')).not.toBeInTheDocument()
})
