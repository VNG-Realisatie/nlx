// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import { ManagementServiceApi } from '../../api'
import { RootStore, StoreProvider } from '../../stores'
import FinancePage from './index'

jest.mock('../../components/PageTemplate')

test('it shows a message when finance is disabled', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceIsFinanceEnabled = jest
    .fn()
    .mockResolvedValue({
      enabled: false,
    })

  const rootStore = new RootStore({
    managementApiClient,
  })

  await rootStore.financeStore.fetch()

  const { getByText } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <FinancePage />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(await getByText('Configure the transaction log')).toBeInTheDocument()
})

test('it shows download link when finance is enabled', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceIsFinanceEnabled = jest
    .fn()
    .mockResolvedValue({
      enabled: true,
    })

  const rootStore = new RootStore({
    managementApiClient,
  })

  await rootStore.financeStore.fetch()

  const { getByText } = renderWithProviders(
    <MemoryRouter>
      <StoreProvider rootStore={rootStore}>
        <FinancePage />
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(await getByText('Export report')).toBeInTheDocument()
})
