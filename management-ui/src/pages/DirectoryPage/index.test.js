// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { act, renderWithProviders, waitFor } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import deferredPromise from '../../test-utils/deferred-promise'
import DirectoryPage from './index'

jest.mock('./DirectoryServices', () => ({ directoryServices }) => {
  const { services } = directoryServices
  return (
    <div data-testid="mock-directory-services">
      {services.map((o, i) => (
        <span key={i} data-testid={`mock-directory-service-${i}`}>
          {o.serviceName}
        </span>
      ))}
    </div>
  )
})

test('listing all services', async () => {
  const directoryServices = deferredPromise()
  const getDirectoryServices = jest.fn(() => directoryServices)

  const { getByRole, getByTestId } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <DirectoryPage getDirectoryServices={getDirectoryServices} />
      </UserContextProvider>
    </Router>,
  )

  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('mock-directory-services')).toThrow()
  expect(getByTestId('directory-description')).toHaveTextContent(
    /^List of all available services$/,
  )

  await act(async () => {
    directoryServices.resolve({ services: [{ serviceName: 'Test Service' }] })
  })

  waitFor(() =>
    expect(getByTestId('mock-directory-services')).toBeInTheDocument(),
  )
  expect(() => getByRole('progressbar')).toThrow()

  expect(getByTestId('mock-directory-service-0')).toHaveTextContent(
    'Test Service',
  )
  expect(getByTestId('directory-description')).toHaveTextContent(
    /^List of all available services \(1\)$/,
  )
})

test('no services', async () => {
  const getDirectoryServices = jest.fn(() => Promise.resolve({ services: [] }))

  const { findByTestId, getByTestId } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <DirectoryPage getDirectoryServices={getDirectoryServices} />
      </UserContextProvider>
    </Router>,
  )

  await act(async () => {
    expect(await findByTestId('mock-directory-services')).toBeInTheDocument()
    expect(() => getByTestId('mock-directory-service-0')).toThrow()
    expect(getByTestId('directory-description')).toHaveTextContent(
      /^List of all available services \(0\)$/,
    )
  })
})

test('failed to load services', async () => {
  jest.spyOn(console, 'error').mockImplementation(() => undefined)
  const getDirectoryServices = jest.fn(async () => {
    throw new Error('arbitrary error')
  })

  const { findByTestId, getByTestId } = renderWithProviders(
    <Router>
      <UserContextProvider user={{}}>
        <DirectoryPage getDirectoryServices={getDirectoryServices} />
      </UserContextProvider>
    </Router>,
  )

  await act(async () => {
    expect(await findByTestId('error-message')).toBeInTheDocument()

    expect(getByTestId('directory-description')).toHaveTextContent(
      /^List of all available services$/,
    )
    expect(() => getByTestId('mock-directory-services')).toThrow()
    expect(getByTestId('error-message')).toHaveTextContent(
      /^Failed to load the directory\.$/,
    )
  })
  console.error.mockRestore()
})
