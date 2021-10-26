// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../../test-utils'
import OutwayModel from '../../../../stores/models/OutwayModel'
import { RootStore, StoreProvider } from '../../../../stores'
import { ManagementApi } from '../../../../api'
import Outways from './index'

const render = (outways) => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  return renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <MemoryRouter>
        <Outways outways={outways} />
      </MemoryRouter>
    </StoreProvider>,
  )
}

test('no outways', () => {
  render([])

  expect(
    screen.getByText(/^There are no outways registered yet$/),
  ).toBeInTheDocument()
})

test('list all outways', () => {
  const outways = [
    new OutwayModel({ outwayData: { name: 'outway1' } }),
    new OutwayModel({ outwayData: { name: 'outway2' } }),
  ]

  render(outways)

  expect(screen.getByTestId('outways-list')).toBeInTheDocument()
  expect(screen.getAllByTestId('outway-row')).toHaveLength(2)
})
