// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { act } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import ServicesPage from './index'

test('listing all services', async () => {
  let resolveFetchServices
  const fetchServicesPromise = new Promise((resolve) => {
    resolveFetchServices = resolve
  })
  const fetchServicesHandler = jest.fn(() => fetchServicesPromise)

  const { getByRole, getByTestId, findByTestId } = renderWithProviders(
    <Router>
      <ServicesPage getServices={fetchServicesHandler} />
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('services-list')).toThrow()

  await act(async () => {
    resolveFetchServices([
      {
        name: 'My First Service',
        authorizationSettings: {
          mode: 'none',
          authorizations: [],
        },
      },
    ])
  })

  expect(await findByTestId('services-list')).toBeInTheDocument()
})

test('no services', async () => {
  const fetchServicesHandler = jest.fn(() => Promise.resolve([]))

  const { findByText, getByTestId } = renderWithProviders(
    <Router>
      <ServicesPage getServices={fetchServicesHandler} />
    </Router>,
  )

  await act(async () => {
    expect(
      await findByText(/^There are no services yet\.$/),
    ).toBeInTheDocument()
    expect(() => getByTestId('services-list')).toThrow()
  })
})

test('failed to load services', async () => {
  const fetchServicesHandler = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const { findByText, getByTestId } = renderWithProviders(
    <Router>
      <ServicesPage getServices={fetchServicesHandler} />
    </Router>,
  )

  expect(() => getByTestId('services-list')).toThrow()
  expect(
    await findByText(/^Failed to load the services\.$/),
  ).toBeInTheDocument()
})
