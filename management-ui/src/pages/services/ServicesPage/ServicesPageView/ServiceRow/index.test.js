// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../../../test-utils'
import ServiceRow from './index'

test('service row should render expected data', () => {
  const service = {
    name: 'Service name',
    internal: true,
    inways: ['inway2'],
    incomingAccessRequestCount: 0,
  }
  const { queryByTestId, queryByText, rerender } = renderWithProviders(
    <MemoryRouter>
      <table>
        <tbody>
          <ServiceRow service={service} />
        </tbody>
      </table>
    </MemoryRouter>,
  )

  expect(queryByText('Service name')).toBeInTheDocument()
  expect(queryByTestId('warning-cell')).toBeEmptyDOMElement()
  expect(queryByText('requestWithCount')).not.toBeInTheDocument()

  const serviceWithIncomingAccessRequest = Object.assign({}, service, {
    incomingAccessRequestCount: 1,
  })

  rerender(
    <MemoryRouter>
      <table>
        <tbody>
          <ServiceRow service={serviceWithIncomingAccessRequest} />
        </tbody>
      </table>
    </MemoryRouter>,
  )

  expect(queryByText('requestWithCount')).toBeInTheDocument()
})
