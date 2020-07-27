// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders, fireEvent } from '../../../../../test-utils'
import { AccessRequestContext } from '../../index'
import DirectoryServiceRow from './index'

describe('show data from a service we do not have access to', () => {
  let service
  let requestSentTo
  let handleRequestAccess

  beforeEach(() => {
    service = {
      organizationName: 'Test Organization',
      serviceName: 'Test Service',
      state: 'degraded',
      apiSpecificationType: 'API',
    }
    requestSentTo = {
      organizationName: '',
      serviceName: '',
    }
    handleRequestAccess = jest.fn()
  })

  afterEach(() => {
    handleRequestAccess.mockClear()
  })

  it('shows the data', () => {
    const { getByTestId, getByText } = renderWithProviders(
      <AccessRequestContext.Provider
        value={{ requestSentTo, handleRequestAccess }}
      >
        <table>
          <tbody>
            <DirectoryServiceRow service={service} />
          </tbody>
        </table>
      </AccessRequestContext.Provider>,
    )

    const serviceRow = getByTestId('directory-service-row')
    expect(serviceRow).toHaveTextContent('Test Organization')
    expect(serviceRow).toHaveTextContent('Test Service')
    expect(serviceRow).toHaveTextContent('state-degraded.svg')
    expect(serviceRow).toHaveTextContent('API')

    const button = getByText('Request')
    expect(button)
    expect(button).not.toBeVisible()
  })

  it('should be possible to request access', () => {
    const { getByText } = renderWithProviders(
      <AccessRequestContext.Provider
        value={{ requestSentTo, handleRequestAccess }}
      >
        <table>
          <tbody>
            <DirectoryServiceRow service={service} />
          </tbody>
        </table>
      </AccessRequestContext.Provider>,
    )

    const button = getByText('Request')
    fireEvent.click(button)

    expect(handleRequestAccess).toHaveBeenCalledWith({
      organizationName: 'Test Organization',
      serviceName: 'Test Service',
    })
  })
})

test('show the state of the latest access request', () => {
  const service = {
    organizationName: 'Test Organization',
    serviceName: 'Test Service',
    state: 'degraded',
    apiSpecificationType: 'API',
    latestAccessRequest: {
      id: 'string',
      state: 'FAILED',
      createdAt: '2020-06-30T08:31:41.106Z',
      updatedAt: '2020-06-30T08:31:41.106Z',
    },
  }
  const requestSentTo = {
    organizationName: 'Test Organization',
    serviceName: 'Test Service',
  }
  const handleRequestAccess = jest.fn()

  const { getByTestId } = renderWithProviders(
    <AccessRequestContext.Provider
      value={{ requestSentTo, handleRequestAccess }}
    >
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>
    </AccessRequestContext.Provider>,
  )

  const serviceRow = getByTestId('directory-service-row')
  expect(serviceRow.querySelector('button')).not.toBeInTheDocument()
})
