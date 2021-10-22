// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter, Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { screen } from '@testing-library/react'
import { renderWithProviders, waitFor } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import InwaysAndOutwaysPage from './index'

jest.mock('../../../components/PageTemplate')
jest.mock('./Inways', () => () => <p data-testid="inways-list">mock inways</p>)

test('the Overviews page', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [],
  })

  const history = createMemoryHistory({})
  const rootStore = new RootStore({
    managementApiClient,
  })

  renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <InwaysAndOutwaysPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  const showAllButton = screen.getByLabelText('Show all')
  expect(showAllButton.getAttribute('href')).toBe('/inways-and-outways')

  const showInwaysButton = screen.getByLabelText('Show Inways')
  expect(showInwaysButton.getAttribute('href')).toBe('/inways-and-outways')

  const showOutwaysButton = screen.getByLabelText('Show Outways')
  expect(showOutwaysButton.getAttribute('href')).toBe('/inways-and-outways')
})

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

  const history = createMemoryHistory({
    initialEntries: ['/inways-and-outways'],
  })
  const rootStore = new RootStore({
    managementApiClient,
  })

  const { getByRole, getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <InwaysAndOutwaysPage />
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
          <InwaysAndOutwaysPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(() => getByTestId('inways-list')).toThrow()
  expect(await findByText(/^Failed to load the inways$/)).toBeInTheDocument()
})
