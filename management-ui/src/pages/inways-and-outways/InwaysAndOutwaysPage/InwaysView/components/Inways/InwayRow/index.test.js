// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../../../../../test-utils'
import InwayModel from '../../../../../../../stores/models/InwayModel'
import InwayRow from './index'

test('rendering the InwayRow', () => {
  const inwayModel = new InwayModel({
    inway: {
      name: 'inway-name',
      hostname: 'MyComputer.local',
      selfAddress: 'inway.organization-a.nlx.local:7913',
      services: [
        {
          name: 'service1',
        },
        {
          name: 'service2',
        },
      ],
      version: 'v0.0.42',
    },
  })

  const { getByText, getByTestId } = renderWithProviders(
    <MemoryRouter>
      <table>
        <tbody>
          <InwayRow inway={inwayModel} />
        </tbody>
      </table>
    </MemoryRouter>,
  )

  expect(getByText('inway-name')).toBeInTheDocument()
  expect(getByText('MyComputer.local')).toBeInTheDocument()
  expect(getByText('inway.organization-a.nlx.local:7913')).toBeInTheDocument()
  expect(getByTestId('services-count')).toHaveTextContent('2')
  expect(getByText('v0.0.42')).toBeInTheDocument()
})
