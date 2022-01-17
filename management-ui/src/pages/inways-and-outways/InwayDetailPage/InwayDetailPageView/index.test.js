// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import InwayModel from '../../../../stores/models/InwayModel'
import { RootStore } from '../../../../stores'
import InwayDetails from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('should display inway details', () => {
  const rootStore = new RootStore({})

  const inway = new InwayModel({
    store: rootStore.inwayStore,
    inway: {
      name: 'name',
      ipAddress: '127.0.0.1',
      hostname: 'host.name',
      selfAddress: 'self.address',
      services: [],
    },
  })

  renderWithProviders(
    <Router>
      <InwayDetails inway={inway} />
    </Router>,
  )

  expect(screen.getByTestId('gateway-type')).toHaveTextContent('inway')
  expect(screen.getByText('127.0.0.1')).toBeInTheDocument()
  expect(screen.getByText('host.name')).toBeInTheDocument()
  expect(screen.getByText('self.address')).toBeInTheDocument()

  fireEvent.click(screen.getByTestId('inway-services'))

  expect(
    screen.getByText('No services have been connected'),
  ).toBeInTheDocument()

  inway.update({
    services: [
      {
        name: 'my-service',
      },
    ],
  })

  fireEvent.click(screen.getByTestId('inway-services'))

  expect(screen.getByText('my-service')).toBeInTheDocument()
  expect(
    screen.queryByText('No services have been connected'),
  ).not.toBeInTheDocument()
})
