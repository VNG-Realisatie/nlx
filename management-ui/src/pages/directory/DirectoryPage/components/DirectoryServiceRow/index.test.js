// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observable } from 'mobx'

import { renderWithProviders, fireEvent, act } from '../../../../../test-utils'
import DirectoryServiceRow from './index'

jest.mock('../../../../../models/OutgoingAccessRequestModel')

describe('a service we do not have access to', () => {
  let service

  global.confirm = jest.fn(() => true)

  beforeEach(() => {
    service = observable({
      id: 'Test Organization/Test Service',
      organizationName: 'Test Organization',
      serviceName: 'Test Service',
      state: 'degraded',
      apiSpecificationType: 'API',
      latestAccessRequest: null,
      requestAccess: jest.fn(),

      isLoading: false,
    })
  })

  it('shows the data', () => {
    const { getByTestId, getByText } = renderWithProviders(
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>,
    )

    const serviceRow = getByTestId('directory-service-row')
    expect(serviceRow).toHaveTextContent('Test Organization')
    expect(serviceRow).toHaveTextContent('Test Service')
    expect(serviceRow).toHaveTextContent('state-degraded.svg')
    expect(serviceRow).toHaveTextContent('API')

    const button = getByText('Request')
    expect(button).not.toBeVisible()
  })

  it('should be possible to request access', () => {
    const { getByText } = renderWithProviders(
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>,
    )

    const button = getByText('Request')
    fireEvent.click(button)

    expect(service.requestAccess).toHaveBeenCalled()
  })

  it('should reflect a change of state', () => {
    const { getByTestId } = renderWithProviders(
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>,
    )

    act(() => {
      service.state = 'up'
    })

    const serviceRow = getByTestId('directory-service-row')
    expect(serviceRow).not.toHaveTextContent('state-degraded.svg')
    expect(serviceRow).toHaveTextContent('state-up.svg')
  })

  it('shows the state of the latest access request', () => {
    service.latestAccessRequest = {
      id: 'string',
      state: 'FAILED',
      createdAt: '2020-06-30T08:31:41.106Z',
      updatedAt: '2020-06-30T08:31:41.106Z',
    }

    const { getByTestId } = renderWithProviders(
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>,
    )

    const serviceRow = getByTestId('directory-service-row')
    expect(serviceRow.querySelector('button')).not.toBeInTheDocument()
  })
})
