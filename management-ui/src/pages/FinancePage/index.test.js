// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter } from 'react-router-dom'

import { renderWithProviders } from '../../test-utils'
import { ManagementApi } from '../../api'
import { RootStore, StoreProvider } from '../../stores'
import FinancePage from './index'

jest.mock('../../components/PageTemplate')

test('it shows a message when finance is disabled', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementIsFinanceEnabled = jest.fn().mockResolvedValue({
    enabled: false,
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  await rootStore.financeStore.fetch()

  const { getByText } = renderWithProviders(
    <StaticRouter>
      <StoreProvider rootStore={rootStore}>
        <FinancePage />
      </StoreProvider>
    </StaticRouter>,
  )

  expect(await getByText('Configure the transaction log')).toBeInTheDocument()
})

test('it shows download link when finance is enabled', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementIsFinanceEnabled = jest.fn().mockResolvedValue({
    enabled: true,
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  await rootStore.financeStore.fetch()

  const { getByText } = renderWithProviders(
    <StaticRouter>
      <StoreProvider rootStore={rootStore}>
        <FinancePage />
      </StoreProvider>
    </StaticRouter>,
  )

  expect(await getByText('Export report')).toBeInTheDocument()
})
