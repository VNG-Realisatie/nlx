// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { act } from '@testing-library/react'
import { renderWithProviders } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import deferredPromise from '../../../test-utils/deferred-promise'
import InwaysPage from './index'

jest.mock('./InwaysPageView', () => () => (
  <p data-testid="inways-list">mock inways</p>
))

test('listing all inways', async () => {
  const fetchInwaysPromise = deferredPromise()
  const fetchInwaysHandler = jest.fn(() => fetchInwaysPromise)

  const { getByRole, getByTestId, findByTestId } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <InwaysPage getInways={fetchInwaysHandler} />
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('inways-list')).toThrow()

  await act(async () => {
    fetchInwaysPromise.resolve([
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
    ])
  })

  expect(await findByTestId('inways-list')).toHaveTextContent('mock inways')
})

test('failed to load inways', async () => {
  const fetchInwaysHandler = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const { findByText, getByTestId } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <InwaysPage getInways={fetchInwaysHandler} />
      </UserContextProvider>
    </Router>,
  )

  expect(() => getByTestId('inways-list')).toThrow()
  expect(await findByText(/^Failed to load the inways\.$/)).toBeInTheDocument()
})
