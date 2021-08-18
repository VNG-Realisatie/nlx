// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import Table from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <Router>
        <Table>
          <thead>
            <Table.TrHead>
              <Table.Th>Heading</Table.Th>
            </Table.TrHead>
          </thead>
          <tbody>
            <Table.Tr>
              <Table.Td>Cell</Table.Td>
            </Table.Tr>
          </tbody>
        </Table>
      </Router>,
    ),
  ).not.toThrow()
})

test('adds a column for the links', () => {
  const result = renderWithProviders(
    <Router>
      <Table withLinks>
        <thead>
          <Table.TrHead>
            <Table.Th>Heading</Table.Th>
          </Table.TrHead>
        </thead>
        <tbody>
          <Table.Tr to="cell" name="Cell">
            <Table.Td>Cell</Table.Td>
          </Table.Tr>
        </tbody>
      </Table>
    </Router>,
  )
  expect(result.container.querySelectorAll('thead tr th')).toHaveLength(2)
  expect(result.container.querySelectorAll('tbody tr td')).toHaveLength(2)
  expect(
    result.container.querySelector('tbody tr td:last-child svg'),
  ).toBeTruthy()
})
