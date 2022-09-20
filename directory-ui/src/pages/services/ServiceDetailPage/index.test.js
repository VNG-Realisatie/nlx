// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter, Routes, Route } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../test-utils'
import ServiceDetailPage from './index'

/* eslint-disable react/display-name */
jest.mock('./components/DirectoryDetailView', () => () => (
  <div data-testid="directory-service-details" />
))

const services = [
  {
    id: 'Test Organization/my-service',
    organization: {
      name: 'Test Organization',
      serialNumber: '01234567890123456789',
    },
    name: 'my-service',
    apiType: 'API',
  },
]

test('display directory service details', async () => {
  renderWithProviders(
    // Router & Route still required for hooks
    // Note not they, but the service data is tested
    <MemoryRouter initialEntries={['/organization/my-service']}>
      <Routes>
        <Route
          path="/:organizationName/:serviceName"
          element={<ServiceDetailPage services={services} />}
        />
      </Routes>
    </MemoryRouter>,
  )

  expect(await screen.findByText('Test Organization')).toBeInTheDocument()
  expect(screen.getByText('my-service')).toBeInTheDocument()
  expect(screen.getByText('state-unknown.svg')).toBeInTheDocument()
  expect(screen.getByTestId('directory-service-details')).toBeInTheDocument()
})

test('service does not exist', () => {
  const { getByTestId, getByText, queryByText } = renderWithProviders(
    <MemoryRouter initialEntries={['/organization/unknown-service']}>
      <Routes>
        <Route
          path="/:organizationName/:serviceName"
          element={<ServiceDetailPage services={services} />}
        />
      </Routes>
    </MemoryRouter>,
  )

  const message = getByTestId('error-message')
  expect(message).toBeInTheDocument()
  expect(message.textContent).toBe(
    "Kan de service 'unknown-service' niet vinden.",
  )

  expect(getByText('unknown-service')).toBeInTheDocument()
  expect(queryByText('organization')).toBeNull()

  const closeButton = getByTestId('close-button')
  expect(closeButton).toBeInTheDocument()
})
