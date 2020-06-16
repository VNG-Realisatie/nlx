// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import DirectoryServices from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <DirectoryServices directoryServices={() => ({ services: [] })} />,
    ),
  ).not.toThrow()
})

test('show a empty services message', () => {
  const { getByTestId } = renderWithProviders(
    <DirectoryServices directoryServices={() => ({ services: [] })} />,
  )
  expect(getByTestId('directory-no-services')).toHaveTextContent(
    'There are no services yet.',
  )
})
test('show a table with rows for every service', () => {
  const { getByTestId, getByRole } = renderWithProviders(
    <MemoryRouter>
      <DirectoryServices
        directoryServices={() => ({
          services: [
            {
              organizationName: 'Test Organization',
              serviceName: 'Test Service',
              status: 'degraded',
              apiSpecificationType: 'API',
            },
          ],
        })}
      />
    </MemoryRouter>,
  )
  expect(getByRole('grid')).toBeTruthy()
  const serviceRow = getByTestId('directory-service-row-0')
  expect(serviceRow).toHaveTextContent('Test Organization')
  expect(serviceRow).toHaveTextContent('Test Service')
  expect(serviceRow).toHaveTextContent('status-degraded.svg')
  expect(serviceRow).toHaveTextContent('API')
})
