// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import OrderDetailView from './index'

test('display order details', () => {
  renderWithProviders(
    <OrderDetailView
      order={{
        reference: 'my-reference',
        delegatee: 'delegatee',
        validFrom: new Date('2021-01-01T00:00:00.000Z'),
        validUntil: new Date('3000-01-01T00:00:00.000Z'),
        services: [],
        revokedAt: null,
      }}
      revokeHandler={() => {}}
    />,
  )

  expect(screen.getByTestId('status')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'up.svg' + 'Order is active' + 'Revoke',
  )
  expect(screen.getByTestId('start-end-date')).toHaveTextContent(
    // eslint-disable-next-line no-useless-concat
    'timer.svg' + 'Valid until date' + 'Since date',
  )
  expect(screen.getByText('my-reference')).toBeInTheDocument()
  expect(screen.getByText('Requestable services')).toBeInTheDocument()
})
