// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../test-utils'
import OutwayModel from '../../../../../stores/models/OutwayModel'
import OutwayRow from './index'

test('rendering the OutwayRow', () => {
  const outwayModel = new OutwayModel({
    outwayData: {
      name: 'outway-name',
      version: 'v0.0.42',
    },
  })

  renderWithProviders(
    <table>
      <tbody>
        <OutwayRow outway={outwayModel} />
      </tbody>
    </table>,
  )

  expect(screen.getByText('outway-name')).toBeInTheDocument()
  expect(screen.getByText('v0.0.42')).toBeInTheDocument()
})
