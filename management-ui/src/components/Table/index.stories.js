// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import Table from './index'

const tableStory = {
  title: 'Components/Table',
  parameters: {
    componentSubtitle: 'Table component.',
  },
  component: Table,
}

export default tableStory

export const intro = () => (
  <Table>
    <thead>
      <Table.Tr>
        <Table.Th>Heading A</Table.Th>
        <Table.Th>Heading B</Table.Th>
        <Table.Th>Heading C</Table.Th>
      </Table.Tr>
    </thead>
    <tbody>
      <Table.Tr>
        <Table.Td>Cell 1</Table.Td>
        <Table.Td>Cell 2</Table.Td>
        <Table.Td>Cell 3</Table.Td>
      </Table.Tr>

      <Table.Tr>
        <Table.Td>Cell 4</Table.Td>
        <Table.Td>Cell 5</Table.Td>
        <Table.Td>Cell 6</Table.Td>
      </Table.Tr>

      <Table.Tr>
        <Table.Td>Cell 7</Table.Td>
        <Table.Td>Cell 8</Table.Td>
        <Table.Td>Cell 9</Table.Td>
      </Table.Tr>
    </tbody>
  </Table>
)

export const selected = () => (
  <Router>
    <p>Highlight selected rows using the `selected` prop.</p>
    <Table>
      <tbody>
        <Table.Tr selected>
          <Table.Td>Item A</Table.Td>
        </Table.Tr>

        <Table.Tr>
          <Table.Td>Item B</Table.Td>
        </Table.Tr>

        <Table.Tr>
          <Table.Td>Item C</Table.Td>
        </Table.Tr>
      </tbody>
    </Table>
  </Router>
)

export const withLinks = () => (
  <Router>
    <p>Displays rows as links using React Router's Link component.</p>
    <Table withLinks>
      <thead>
        <Table.TrHead>
          <Table.Th>My items</Table.Th>
        </Table.TrHead>
      </thead>
      <tbody>
        <Table.Tr to="/item-a">
          <Table.Td>Item A</Table.Td>
        </Table.Tr>

        <Table.Tr to="/item-b">
          <Table.Td>Item B</Table.Td>
        </Table.Tr>

        <Table.Tr to="/item-c">
          <Table.Td>Item C</Table.Td>
        </Table.Tr>
      </tbody>
    </Table>
  </Router>
)
