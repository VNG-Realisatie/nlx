// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter, Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { act } from '@testing-library/react'
import { renderWithProviders, waitFor } from '../../test-utils'
import deferredPromise from '../../test-utils/deferred-promise'
import { UserContextProvider } from '../../user-context'
import ServicesPage from './index'

jest.mock('./ServicesPageView', () => () => (
  <p data-testid="services-list">mock-services</p>
))

test('fetching all services', async () => {
  const history = createMemoryHistory({ initialEntries: ['/services'] })

  const fetchServicesPromise = deferredPromise()
  const fetchServicesHandler = jest.fn(() => fetchServicesPromise)

  const { getByRole, getByTestId, getByLabelText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <ServicesPage getServices={fetchServicesHandler} />
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('services-list')).toThrow()

  await act(async () => {
    fetchServicesPromise.resolve([
      {
        name: 'my-first-service',
        authorizationSettings: {
          mode: 'none',
          authorizations: [],
        },
        inways: [],
        internal: false,
      },
    ])
  })

  waitFor(() =>
    expect(getByTestId('services-list')).toHaveTextContent('mock-services'),
  )
  expect(getByTestId('service-count')).toHaveTextContent('1Services')

  const linkAddService = getByLabelText(/Add service/)
  expect(linkAddService.getAttribute('href')).toBe('/services/add-service')
})

test('failed to load services', async () => {
  const fetchServicesHandler = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const { findByText, getByTestId } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider user={{}}>
        <ServicesPage getServices={fetchServicesHandler} />
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(() => getByTestId('services-list')).toThrow()
  expect(
    await findByText(/^Failed to load the services\.$/),
  ).toBeInTheDocument()
  expect(getByTestId('service-count')).toHaveTextContent('0Services')
})
