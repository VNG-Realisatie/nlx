// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../test-utils'
import OrderRow from './index'

test('order row should render expected data', () => {
  const order = {
    description: 'my own description',
    delegatee: 'delegatee',
    services: [
      { organization: 'organization X', service: 'service Y' },
      { organization: 'organization Y', service: 'service Z' },
    ],
  }
  const { getByText, getByTitle } = renderWithProviders(
    <table>
      <tbody>
        <OrderRow order={order} />
      </tbody>
    </table>,
  )

  expect(getByText('my own description')).toBeInTheDocument()
  expect(getByText('delegatee')).toBeInTheDocument()
  expect(getByTitle('organization X - service Y')).toBeInTheDocument()
  expect(getByTitle('organization Y - service Z')).toHaveTextContent(
    'organization Y - service Z',
  )
})
