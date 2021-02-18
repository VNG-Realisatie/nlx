// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter } from 'react-router-dom'

import { renderWithProviders } from '../../test-utils'
import FinancePage from './index'

jest.mock('../../components/PageTemplate')

test('it shows download link', async () => {
  const { getByText } = renderWithProviders(
    <StaticRouter>
      <FinancePage />
    </StaticRouter>,
  )

  expect(await getByText('Export report')).toBeInTheDocument()
})
