// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import ParticipantsTable from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <Router>
        <ParticipantsTable participants={[]} />
      </Router>,
    ),
  ).not.toThrow()
})

test('show a empty participants message', () => {
  const { getByTestId } = renderWithProviders(
    <Router>
      <ParticipantsTable participants={[]} />
    </Router>,
  )
  expect(getByTestId('directory-no-participants')).toHaveTextContent(
    'Geen deelnemers gevonden',
  )
})

test('show a table with rows for every participant', () => {
  const { getByTestId, getByRole } = renderWithProviders(
    <Router>
      <ParticipantsTable
        participants={[
          {
            organization: {
              name: 'Test Organization',
              serialNumber: '00000000000000000001',
            },
            createdAt: '2021-01-01T00:00:00',
            serviceCount: 5,
            inwayCount: 4,
            outwayCount: 3,
          },
        ]}
      />
    </Router>,
  )

  expect(getByRole('grid')).toBeTruthy()
  expect(getByTestId('directory-participant-row')).toBeInTheDocument()
})
