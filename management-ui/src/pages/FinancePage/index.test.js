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

test('it shows download link', async () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({
    managementApiClient,
  })

  const { getByText } = renderWithProviders(
    <StaticRouter>
      <StoreProvider rootStore={rootStore}>
        <FinancePage />
      </StoreProvider>
    </StaticRouter>,
  )

  expect(await getByText('Export report')).toBeInTheDocument()
})
