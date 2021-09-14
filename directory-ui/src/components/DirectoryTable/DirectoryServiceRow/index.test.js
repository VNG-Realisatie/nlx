// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import { SERVICE_STATE_DEGRADED } from '../../StateIndicator'
import DirectoryServiceRow from './index'

const serviceData = {
  id: 'my-service',
  organization: 'Test Organization',
  name: 'Test Service',
  status: SERVICE_STATE_DEGRADED,
  apiType: 'API',
}

const renderComponent = ({ service }) => {
  return renderWithProviders(
    <table>
      <tbody>
        <DirectoryServiceRow service={service} />
      </tbody>
    </table>,
  )
}

test('display service information', () => {
  const service = serviceData
  const { container } = renderComponent({ service })

  expect(container).toHaveTextContent('Test Organization')
  expect(container).toHaveTextContent('Test Service')
  expect(container).toHaveTextContent('state-degraded.svg')
  expect(container).toHaveTextContent('API')
})
