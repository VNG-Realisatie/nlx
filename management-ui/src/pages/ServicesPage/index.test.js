// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { act } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import ServicesPage from './index'

test('listing all services', async () => {
  let resolveFetchServices
  const fetchServicesPromise = new Promise((resolve) => {
    resolveFetchServices = resolve
  })
  const fetchServicesHandler = jest.fn(() => fetchServicesPromise)

  const {
    getByRole,
    getByTestId,
    getByLabelText,
    findByTestId,
  } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <ServicesPage getServices={fetchServicesHandler} />
      </UserContextProvider>
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
  expect(getByTestId('service-count')).toHaveTextContent('1Services')

  const linkAddService = getByLabelText(/Add service/)
  expect(linkAddService.getAttribute('href')).toBe('/services/add-service')
})

test('no services', async () => {
  const fetchServicesHandler = jest.fn(() => Promise.resolve([]))

  const { findByText, getByTestId } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <ServicesPage getServices={fetchServicesHandler} />
      </UserContextProvider>
    </Router>,
  )

  await act(async () => {
    expect(
      await findByText(/^There are no services yet\.$/),
    ).toBeInTheDocument()
    expect(() => getByTestId('services-list')).toThrow()
    expect(getByTestId('service-count')).toHaveTextContent('0Services')
  })
})

test('failed to load services', async () => {
  const fetchServicesHandler = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  const { findByText, getByTestId } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <ServicesPage getServices={fetchServicesHandler} />
      </UserContextProvider>
    </Router>,
  )

  expect(() => getByTestId('services-list')).toThrow()
  expect(
    await findByText(/^Failed to load the services\.$/),
  ).toBeInTheDocument()
  expect(getByTestId('service-count')).toHaveTextContent('0Services')
})
