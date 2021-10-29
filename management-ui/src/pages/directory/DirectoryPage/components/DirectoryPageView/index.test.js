// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../test-utils'
import DirectoryServices from './index'

// eslint-disable-next-line react/prop-types
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
    'There are no services yet',
  )
})

test('show a table with rows for every service', () => {
  const { getByTestId, getByRole } = renderWithProviders(
    <DirectoryServices
      services={[
        {
          organization: {
            name: 'Test Organization',
            serialNumber: '00000000000000000001',
          },
          serviceName: 'Test Service',
        },
      ]}
    />,
  )

  expect(getByRole('grid')).toBeTruthy()
  expect(getByTestId('testrow')).toBeInTheDocument()
})
