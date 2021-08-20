// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../../test-utils'
import StartEndDate from './index'

test('display active order dates', () => {
  renderWithProviders(
    <StartEndDate
      validFrom={new Date('2021-01-01T00:00:00.000Z')}
      validUntil={new Date('3000-01-01T00:00:00.000Z')}
      revokedAt={null}
    />,
  )

  expect(screen.queryByText('Valid until date')).toBeInTheDocument()
  expect(screen.queryByText('Since date')).toBeInTheDocument()
})

test('display revoked order dates', () => {
  renderWithProviders(
    <StartEndDate
      validFrom={new Date('2021-01-01T00:00:00.000Z')}
      validUntil={new Date('3000-01-01T00:00:00.000Z')}
      revokedAt={new Date('2021-01-01T00:00:00.000Z')}
    />,
  )

  expect(screen.queryByText('Revoked on date')).toBeInTheDocument()
})
