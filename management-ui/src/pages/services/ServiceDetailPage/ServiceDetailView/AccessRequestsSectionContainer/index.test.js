// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, waitForElementToBeRemoved } from '@testing-library/react'
import { fireEvent, renderWithProviders } from '../../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../../stores'
import { ACCESS_REQUEST_STATES } from '../../../../../stores/models/IncomingAccessRequestModel'
import { ManagementApi } from '../../../../../api'
import AccessRequestsSection, { POLLING_INTERVAL } from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('listing the access requests when no access requests are available', async () => {
  const onApproveOrRejectCallbackHandler = jest.fn().mockResolvedValue(null)

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [
      {
        name: 'service-a',
      },
    ],
  })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [],
    })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.servicesStore.fetchAll()

  const { getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessRequestsSection
        service={rootStore.servicesStore.getService('service-a')}
        onApproveOrRejectCallbackHandler={onApproveOrRejectCallbackHandler}
      />
    </StoreProvider>,
  )

  const toggler = getByText(/Access requests/i)
  fireEvent.click(toggler)

  expect(getByText('There are no access requests')).toBeInTheDocument()
})

test('listing the access requests when an access request is available', async () => {
  global.confirm = jest.fn(() => true)
  const onApproveOrRejectCallbackHandler = jest.fn().mockResolvedValue(null)

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [
      {
        name: 'service-a',
      },
    ],
  })

  managementApiClient.managementGetService = jest.fn().mockResolvedValue({
    name: 'service-a',
  })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [
        {
          id: '1',
          serviceName: 'service-a',
          organizationName: 'organization-a',
          state: ACCESS_REQUEST_STATES.RECEIVED,
          createdAt: new Date(),
          updatedAt: new Date(),
          approve: jest.fn().mockResolvedValue(null),
          reject: jest.fn().mockResolvedValue(null),
        },
      ],
    })

  managementApiClient.managementListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [],
    })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.servicesStore.fetchAll()

  const service = rootStore.servicesStore.getService('service-a')
  await service.fetch()

  const { getByTestId, getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessRequestsSection
        service={service}
        onApproveOrRejectCallbackHandler={onApproveOrRejectCallbackHandler}
      />
    </StoreProvider>,
  )

  const accessRequest = service.incomingAccessRequests[0]

  accessRequest.approve = jest.fn().mockResolvedValue()
  accessRequest.reject = jest.fn().mockResolvedValue()

  const toggler = getByText(/Access requests/i)
  fireEvent.click(toggler)

  expect(
    getByTestId('service-incoming-accessrequests-list'),
  ).toBeInTheDocument()
  expect(getByText('organization-a')).toBeInTheDocument()
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

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [
        {
          id: '1',
          serviceName: 'service-a',
          state: ACCESS_REQUEST_STATES.RECEIVED,
          createdAt: new Date(),
          updatedAt: new Date(),
        },
      ],
    })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.servicesStore.fetchAll()

  const { getByTestId, findByTestId } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessRequestsSection
        service={rootStore.servicesStore.getService('service-a')}
        onApproveOrRejectCallbackHandler={() => {}}
      />
    </StoreProvider>,
  )

  expect(getByTestId('amount')).toHaveTextContent('0')

  jest.advanceTimersByTime(POLLING_INTERVAL)

  const amountAccented = await findByTestId('amount-accented')

  expect(amountAccented).toHaveTextContent('1')

  jest.useRealTimers()
})

test('polling with access request section expanded', async () => {
  jest.useFakeTimers()

  const managementApiClient = new ManagementApi()

  managementApiClient.managementGetService = jest.fn().mockResolvedValue({
    name: 'service-a',
  })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValueOnce({
      accessRequests: [
        {
          id: '1',
          serviceName: 'service-a',
          organizationName: 'organization-a',
          state: ACCESS_REQUEST_STATES.RECEIVED,
          createdAt: new Date(),
          updatedAt: new Date(),
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
  const service = rootStore.servicesStore.getService('service-a')

  const { getByText, queryByText, findByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <AccessRequestsSection
        service={service}
        onApproveOrRejectCallbackHandler={() => {}}
      />
    </StoreProvider>,
  )

  const toggler = getByText(/Access requests/i)

  fireEvent.click(toggler)

  expect(getByText('organization-a')).toBeInTheDocument()
  expect(queryByText('Nieuwe verzoeken')).not.toBeInTheDocument()

  act(() => {
    jest.advanceTimersByTime(POLLING_INTERVAL)
  })

  expect(await findByText('organization-a')).toBeInTheDocument()

  expect(getByText('Nieuwe verzoeken')).toBeInTheDocument()

  fireEvent.click(getByText('Nieuwe verzoeken'))

  await waitForElementToBeRemoved(() => getByText('organization-a'))
  expect(queryByText('Nieuwe verzoeken')).not.toBeInTheDocument()

  jest.useRealTimers()
})
