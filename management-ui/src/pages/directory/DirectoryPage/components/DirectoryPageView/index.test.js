// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../test-utils'
import DirectoryServices from './index'

jest.mock('../DirectoryServiceRow', () => ({ organizationName }) => (
  <tr data-testid="testrow">
    <td>{organizationName}</td>
  </tr>
))

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(<DirectoryServices services={[]} />),
  ).not.toThrow()
})

test('show a empty services message', () => {
  const { getByTestId } = renderWithProviders(
    <DirectoryServices services={[]} />,
  )
  expect(getByTestId('directory-no-services')).toHaveTextContent(
    'There are no services yet.',
  )
})

test('show a table with rows for every service', () => {
  const { getByTestId, getByRole } = renderWithProviders(
    <DirectoryServices
      services={[
        {
          organizationName: 'Test Organization',
          serviceName: 'Test Service',
          state: 'degraded',
          apiSpecificationType: 'API',
        },
      ]}
    />,
  )

  expect(getByRole('grid')).toBeTruthy()
  expect(getByTestId('testrow')).toBeInTheDocument()
})
