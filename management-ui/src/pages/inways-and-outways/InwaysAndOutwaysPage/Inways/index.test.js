// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import InwayModel from '../../../../stores/models/InwayModel'
import Inways from './index'

test('no inways', () => {
  renderWithProviders(<Inways inways={[]} />)
  expect(
    screen.getByText(/^There are no inways registered yet$/),
  ).toBeInTheDocument()
})

test('list all inways', () => {
  const inways = [
    new InwayModel({ inway: { name: 'inway1' } }),
    new InwayModel({ inway: { name: 'inway2' } }),
  ]

  renderWithProviders(
    <MemoryRouter>
      <Inways inways={inways} />
    </MemoryRouter>,
  )

  expect(screen.getByTestId('inways-list')).toBeInTheDocument()
  expect(screen.getAllByTestId('inway-row')).toHaveLength(2)
})
