// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders, screen } from '../../../../../../../test-utils'
import Status from './index'

const day = 86400000

test('display status', () => {
  renderWithProviders(
    <Status
      order={{
        reference: 'my-reference',
        delegatee: 'delegatee',
        validFrom: new Date().getTime() - day,
        validUntil: new Date().getTime() + day,
        revokedAt: null,
      }}
    />,
  )

  expect(screen.queryByText('Order is active')).toBeInTheDocument()

  renderWithProviders(
    <Status
      order={{
        reference: 'my-reference',
        delegatee: 'delegatee',
        validFrom: new Date(new Date().getTime() - day),
        validUntil: new Date(new Date().getTime() - day),
        revokedAt: null,
      }}
    />,
  )

  expect(screen.queryByText('Order is expired')).toBeInTheDocument()

  renderWithProviders(
    <Status
      order={{
        reference: 'my-reference',
        delegatee: 'delegatee',
        validFrom: new Date(new Date().getTime() + day),
        validUntil: new Date(new Date().getTime() + 2 * day),
        revokedAt: null,
      }}
    />,
  )

  expect(screen.queryByText('Order is not yet active')).toBeInTheDocument()

  renderWithProviders(
    <Status
      order={{
        reference: 'my-reference',
        delegatee: 'delegatee',
        validFrom: new Date(new Date().getTime() - day),
        validUntil: new Date(new Date().getTime() + day),
        revokedAt: new Date(),
      }}
    />,
  )

  expect(screen.queryByText('Order is revoked')).toBeInTheDocument()
})
