// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../test-utils'

import ServiceRow from './index'

test('service row should render expected data', () => {
  const service = {
    name: 'service',
    internal: true,
    inways: ['inway2'],
  }
  const { getByText, queryByTestId } = renderWithProviders(
    <table>
      <tbody>
        <ServiceRow service={service} />
      </tbody>
    </table>,
  )

  expect(getByText('service')).toBeInTheDocument()
  expect(queryByTestId('warning-cell')).toBeEmptyDOMElement()
})
