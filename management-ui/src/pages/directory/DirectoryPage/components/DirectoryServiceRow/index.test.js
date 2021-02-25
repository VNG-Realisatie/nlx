// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { configure } from 'mobx'
import { fireEvent, act, waitFor } from '@testing-library/react'
import { renderWithProviders } from '../../../../../test-utils'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import {
  SERVICE_STATE_DEGRADED,
  SERVICE_STATE_UP,
} from '../../../../../components/StateIndicator'
import DirectoryServiceRow from './index'

jest.mock('../../../../../stores/models/OutgoingAccessRequestModel')
jest.mock('../../../../../components/Modal')

const buildServiceModel = () => {
  return new DirectoryServiceModel({
    serviceData: {
      id: 'my-service',
      organizationName: 'Test Organization',
      serviceName: 'Test Service',
      state: SERVICE_STATE_DEGRADED,
      apiSpecificationType: 'API',
    },
  })
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
  configure({ safeDescriptors: false })
  const service = buildServiceModel()
  const { container, getByText } = renderComponent({ service })

  expect(container).toHaveTextContent('Test Organization')
  expect(container).toHaveTextContent('Test Service')
  expect(container).toHaveTextContent('state-degraded.svg')
  expect(container).toHaveTextContent('API')

  const button = getByText('Request')
  expect(button).not.toBeVisible()
})

test('requesting access', async () => {
  configure({ safeDescriptors: false })

  const service = buildServiceModel()
  service.requestAccess = jest.fn()

  const { getByText } = renderComponent({ service })

  fireEvent.click(getByText('Request'))
  fireEvent.click(getByText('Send'))

  await waitFor(() => expect(service.requestAccess).toHaveBeenCalled())
})

test('display changes to the service', () => {
  const service = buildServiceModel()
  const { getByTestId } = renderComponent({ service })

  act(() => {
    service.state = SERVICE_STATE_UP
  })

  const serviceRow = getByTestId('directory-service-row')
  expect(serviceRow).not.toHaveTextContent('state-degraded.svg')
  expect(serviceRow).toHaveTextContent('state-up.svg')
})
