// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import Table from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <Table>
        <thead>
          <tr>
            <Table.Th>Heading</Table.Th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <Table.Td>Cell</Table.Td>
          </tr>
        </tbody>
      </Table>,
    ),
  ).not.toThrow()
})
