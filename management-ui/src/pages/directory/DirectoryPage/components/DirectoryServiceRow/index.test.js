// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { makeAutoObservable } from 'mobx'

import { renderWithProviders, fireEvent, act } from '../../../../../test-utils'
import DirectoryServiceRow from './index'

jest.mock('../../../../../stores/models/OutgoingAccessRequestModel')

describe('a service we do not have access to', () => {
  let service

  global.confirm = jest.fn(() => true)

  beforeEach(() => {
    service = makeAutoObservable({
      id: 'Test Organization/Test Service',
      organizationName: 'Test Organization',
      serviceName: 'Test Service',
      state: 'degraded',
      apiSpecificationType: 'API',
      latestAccessRequest: null,
      latestAccessProof: null,
      requestAccess: jest.fn(),
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
    const spy = jest.spyOn(service, 'requestAccess')
    const { getByText } = renderWithProviders(
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>,
    )

    const button = getByText('Request')
    fireEvent.click(button)

    expect(spy).toHaveBeenCalled()
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

  it('has request access button in certain states', () => {
    service.latestAccessRequest = {
      id: 'string',
      state: 'FAILED',
      createdAt: '2020-06-30T08:30:00Z',
      updatedAt: '2020-06-30T08:30:05Z',
    }

    const { getByTestId } = renderWithProviders(
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>,
    )

    const serviceRow = getByTestId('directory-service-row')
    expect(serviceRow.querySelector('button')).toBeInTheDocument()

    service.latestAccessRequest = {
      id: 'string',
      state: 'REJECTED',
      createdAt: '2020-06-30T08:30:00Z',
      updatedAt: '2020-06-30T08:35:00Z',
    }

    expect(serviceRow.querySelector('button')).toBeInTheDocument()
  })

  it('does not have request access button in other states', () => {
    service.latestAccessRequest = {
      id: 'string',
      state: 'CREATED',
      createdAt: '2020-06-30T08:30:00Z',
      updatedAt: '2020-06-30T08:30:03Z',
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

    service.latestAccessRequest = {
      id: 'string',
      state: 'APPROVED',
      createdAt: '2020-06-30T08:30:00Z',
      updatedAt: '2020-06-30T08:35:00Z',
    }
    service.latestAccessProof = {
      id: 'string',
      accessRequestId: 'string',
      createdAt: '2020-06-30T08:35:01Z',
    }

    expect(serviceRow.querySelector('button')).not.toBeInTheDocument()
  })
})
