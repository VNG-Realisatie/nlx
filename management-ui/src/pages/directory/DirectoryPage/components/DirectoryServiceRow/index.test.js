// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { configure } from 'mobx'
import { act, screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../test-utils'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import {
  SERVICE_STATE_DEGRADED,
  SERVICE_STATE_UP,
} from '../../../../../components/StateIndicator'
import DirectoryServiceRow from './index'

jest.mock('../../../../../stores/models/OutgoingAccessRequestModel')

const buildServiceModel = () => {
  return new DirectoryServiceModel({
    serviceData: {
      id: 'my-service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Test Organization',
      },
      serviceName: 'Test Service',
      state: SERVICE_STATE_DEGRADED,
      apiSpecificationType: 'API',
    },
  })
}

const renderComponent = ({ service }) => {
  return renderWithProviders(
    <MemoryRouter>
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>
    </MemoryRouter>,
  )
}

test('display service information', () => {
  configure({ safeDescriptors: false })

  const service = buildServiceModel()
  const { container } = renderComponent({ service })

  expect(container).toHaveTextContent('Test Organization')
  expect(container).toHaveTextContent('Test Service')
  expect(container).toHaveTextContent('state-degraded.svg')
  expect(container).toHaveTextContent('API')
})

test('display changes to the service', () => {
  const service = buildServiceModel()
  renderComponent({ service })

  act(() => {
    service.update({
      serviceData: {
        state: SERVICE_STATE_UP,
      },
    })
  })

  const serviceRow = screen.getByTestId('directory-service-row')
  expect(serviceRow).not.toHaveTextContent('state-degraded.svg')
  expect(serviceRow).toHaveTextContent('state-up.svg')
})
