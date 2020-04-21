// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { act } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import InwaysPage from './index'

test('listing all inways', async () => {
  let resolveFetchInways
  const fetchInwaysPromise = new Promise((resolve) => {
    resolveFetchInways = resolve
  })
  const fetchInwaysHandler = jest.fn(() => fetchInwaysPromise)

  const {
    getByRole,
    getByTestId,
    findByTestId,
    getByText,
  } = renderWithProviders(
    <Router>
      <InwaysPage getInways={fetchInwaysHandler} />
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('inways-list')).toThrow()

  await act(async () => {
    resolveFetchInways([
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

  expect(await findByTestId('inways-list')).toBeInTheDocument()
  expect(getByText('name')).toBeInTheDocument()
  expect(getByText('hostname')).toBeInTheDocument()
  expect(getByText('self-address')).toBeInTheDocument()
  expect(getByTestId('services-count')).toHaveTextContent('1')
  expect(getByText('version')).toBeInTheDocument()
})

test('no inways', async () => {
  const fetchInwaysHandler = jest.fn(() => Promise.resolve([]))

  const { findByText, getByTestId } = renderWithProviders(
    <Router>
      <InwaysPage getInways={fetchInwaysHandler} />
    </Router>,
  )

  await act(async () => {
    expect(
      await findByText(/^There are no inways registered yet\.$/),
    ).toBeInTheDocument()
    expect(() => getByTestId('inways-list')).toThrow()
  })
})

test('failed to load inways', async () => {
  const fetchInwaysHandler = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const { findByText, getByTestId } = renderWithProviders(
    <Router>
      <InwaysPage getInways={fetchInwaysHandler} />
    </Router>,
  )

  expect(() => getByTestId('inways-list')).toThrow()
  expect(await findByText(/^Failed to load the inways\.$/)).toBeInTheDocument()
})
