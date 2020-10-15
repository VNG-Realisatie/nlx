// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter, Router } from 'react-router-dom'
import { act } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithProviders, waitFor } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { mockInwaysStore } from '../../../stores/InwaysStore.mock'
import { StoreProvider } from '../../../stores'
import InwaysPage from './index'

jest.mock('./InwaysPageView', () => () => (
  <p data-testid="inways-list">mock inways</p>
))

test('fetching all inways', async () => {
  const history = createMemoryHistory({ initialEntries: ['/inways'] })

  const store = mockInwaysStore({ inways: null, isInitiallyFetched: false })

  const { getByRole, getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider store={store}>
          <InwaysPage />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('inways-list')).toThrow()

  await act(async () => {
    store.inwaysStore.inways = [
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
    ]
    store.inwaysStore.isReady = true
  })

  waitFor(() =>
    expect(getByTestId('inways-list')).toHaveTextContent('mock-inways'),
  )
})

test('failed to load inways', async () => {
  const store = mockInwaysStore({ inways: null, error: 'arbitrary error' })

  const { findByText, getByTestId } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <StoreProvider store={store}>
          <InwaysPage />
        </StoreProvider>
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(() => getByTestId('inways-list')).toThrow()
  expect(await findByText(/^Failed to load the inways\.$/)).toBeInTheDocument()
})
