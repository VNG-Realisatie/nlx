// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import ParticipantRow from './index'

const participantData = {
  id: 'my-service',
  organization: {
    name: 'Test Organization',
    serialNumber: '00000000000000000001',
  },
  createdAt: '2021-01-01T00:00:00',
  serviceCount: 1,
  inwayCount: 2,
  outwayCount: 3,
}

const renderComponent = ({ participant }) => {
  return renderWithProviders(
    <table>
      <tbody>
        <ParticipantRow participant={participant} />
      </tbody>
    </table>,
  )
}

test('display participant information', () => {
  const participant = participantData
  const { container } = renderComponent({ participant })

  expect(container).toHaveTextContent('Test Organization')
  expect(container).toHaveTextContent('1')
  expect(container).toHaveTextContent('2')
  expect(container).toHaveTextContent('3')
})
