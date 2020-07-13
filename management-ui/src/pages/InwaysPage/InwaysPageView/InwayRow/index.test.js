// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../test-utils'

import InwayRow from './index'

test('row shows expected data', () => {
  const inway = {
    name: 'inway',
    hostname: 'hostname',
    selfAddress: 'selfie',
    services: ['service1', 'service2'],
    version: 'test',
  }
  const { getByText, getByTestId } = renderWithProviders(
    <table>
      <tbody>
        <InwayRow inway={inway} />
      </tbody>
    </table>,
  )

  expect(getByText('inway')).toBeInTheDocument()
  expect(getByText('hostname')).toBeInTheDocument()
  expect(getByText('selfie')).toBeInTheDocument()
  expect(getByTestId('services-count')).toHaveTextContent('2')
  expect(getByText('test')).toBeInTheDocument()
})
