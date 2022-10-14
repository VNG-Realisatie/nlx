// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { TXLogServiceApi, ManagementServiceApi } from '../../api'
import { RootStore, StoreProvider } from '../../stores'
import { renderWithProviders, waitFor } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import { DIRECTION_IN } from '../../stores/models/TransactionLogModel'
import TransactionLogPage from './index'

test('fetching the transaction logs', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceIsTXLogEnabled = jest
    .fn()
    .mockResolvedValue({
      enabled: true,
    })

  const txLogApiClient = new TXLogServiceApi()

  txLogApiClient.tXLogServiceListRecords = jest.fn().mockResolvedValue({
    records: [
      {
        source: {
          serialNumber: '00000000000000000001',
        },
        destination: {
          serialNumber: '00000000000000000002',
        },
        direction: DIRECTION_IN,
        serviceName: 'my-service',
        createdAt: new Date(),
      },
    ],
  })

  const store = new RootStore({
    txLogApiClient: txLogApiClient,
    managementApiClient: managementApiClient,
  })

  const { getByRole, findAllByTestId } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <TransactionLogPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()

  const transactionLogElements = await findAllByTestId('transaction-log-list')
  expect(transactionLogElements).toHaveLength(1)
})

test('failed to load transaction logs', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceIsTXLogEnabled = jest
    .fn()
    .mockResolvedValue({
      enabled: true,
    })

  const txLogApiClient = new TXLogServiceApi()

  txLogApiClient.managementServiceListTransactionRecords = jest
    .fn()
    .mockRejectedValue(Error('arbitrary error'))

  const store = new RootStore({
    txLogApiClient: txLogApiClient,
    managementApiClient: managementApiClient,
  })

  const { queryByRole, getByTestId, findByText } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <TransactionLogPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  await waitFor(() => {
    expect(queryByRole('progressbar')).not.toBeInTheDocument()
  })

  expect(() => getByTestId('transaction-log-record')).toThrow()

  expect(
    await findByText(/^Failed to load the transaction logs$/),
  ).toBeInTheDocument()
  expect(await findByText(/^something went wrong$/)).toBeInTheDocument()
})

test('it shows a message when txlog is disabled', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceIsTXLogEnabled = jest
    .fn()
    .mockResolvedValue({
      enabled: false,
    })

  const txLogApiClient = new TXLogServiceApi()

  txLogApiClient.tXLogServiceListRecords = jest.fn().mockResolvedValue({
    records: [
      {
        source: {
          serialNumber: '00000000000000000001',
        },
        destination: {
          serialNumber: '00000000000000000002',
        },
        direction: DIRECTION_IN,
        serviceName: 'my-service',
        createdAt: new Date(),
      },
    ],
  })

  const store = new RootStore({
    txLogApiClient: txLogApiClient,
    managementApiClient: managementApiClient,
  })

  const { getByText, queryByRole } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={store}>
          <TransactionLogPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  await waitFor(() => {
    expect(queryByRole('progressbar')).not.toBeInTheDocument()
  })

  expect(await getByText('Configure the transaction log')).toBeInTheDocument()
})
