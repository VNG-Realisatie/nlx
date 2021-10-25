// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import OutwayModel from '../../../../stores/models/OutwayModel'
import Outways from './index'

test('no outways', () => {
  renderWithProviders(<Outways outways={[]} />)

  expect(
    screen.getByText(/^There are no outways registered yet$/),
  ).toBeInTheDocument()
})

test('list all outways', () => {
  const outways = [
    new OutwayModel({ outwayData: { name: 'outway1' } }),
    new OutwayModel({ outwayData: { name: 'outway2' } }),
  ]

  renderWithProviders(
    <MemoryRouter>
      <Outways outways={outways} />
    </MemoryRouter>,
  )

  expect(screen.getByTestId('outways-list')).toBeInTheDocument()
  expect(screen.getAllByTestId('outway-row')).toHaveLength(2)
})
