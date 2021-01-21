// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter, Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders, waitFor } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import InwaysPage from './index'

jest.mock('../../../components/PageTemplate')
jest.mock('./InwaysPageView', () => () => (
  <p data-testid="inways-list">mock inways</p>
))

test('fetching all inways', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [
      {
        name: 'name',
        version: 'version',
        hostname: 'hostname',
        selfAddress: 'self-address',
        services: [
          {
            name: 'service-1',
          },
        ],
      },
    ],
  })

  const history = createMemoryHistory({ initialEntries: ['/inways'] })
  const rootStore = new RootStore({
    managementApiClient,
  })

  const { getByRole, getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <InwaysPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('inways-list')).toThrow()

  await waitFor(() =>
    expect(getByTestId('inways-list')).toHaveTextContent('mock inways'),
  )
})

test('failed to load inways', async () => {
  const rootStore = new RootStore({
    inwayRepository: {
      getAll: jest.fn().mockRejectedValue(new Error('arbitrary error')),
    },
  })

  const { findByText, getByTestId } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <InwaysPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(() => getByTestId('inways-list')).toThrow()
  expect(await findByText(/^Failed to load the inways$/)).toBeInTheDocument()
})
