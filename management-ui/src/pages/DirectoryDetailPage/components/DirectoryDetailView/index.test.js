// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent } from '@testing-library/react'

import { renderWithProviders } from '../../../../test-utils'
import { AccessRequestContext } from '../../../DirectoryPage'
import DirectoryDetailView from './index'

describe('detail view of directory service we do not have access to', () => {
  let service
  let requestSentTo
  let handleRequestAccess

  beforeEach(() => {
    service = {
      organizationName: 'Test Organization',
      serviceName: 'Test Service',
      state: 'degraded',
      apiSpecificationType: 'API',
      latestAccessRequest: null,
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

  it('should have a button to request access', () => {
    const { getByText } = renderWithProviders(
      <AccessRequestContext.Provider
        value={{ requestSentTo, handleRequestAccess }}
      >
        <DirectoryDetailView {...service} />
      </AccessRequestContext.Provider>,
    )

    const button = getByText('Request Access')
    expect(button).toBeInTheDocument()

    fireEvent.click(button)
    expect(handleRequestAccess).toHaveBeenCalled()
  })

  it('should show a loading message', () => {
    const serviceWithOutstandingRequest = {
      ...service,
      latestAccessRequest: {
        id: 'string',
        state: 'CREATED',
        createdAt: '2020-06-30T08:31:41.106Z',
        updatedAt: '2020-06-30T08:31:41.106Z',
      },
    }

    const { getByText } = renderWithProviders(
      <AccessRequestContext.Provider
        value={{ requestSentTo, handleRequestAccess }}
      >
        <DirectoryDetailView {...serviceWithOutstandingRequest} />
      </AccessRequestContext.Provider>,
    )

    expect(getByText('Sending request')).toBeInTheDocument()
  })

  it('should show a failed message and retry button', () => {
    const serviceWithFailedRequest = {
      ...service,
      latestAccessRequest: {
        id: 'string',
        state: 'FAILED',
        createdAt: '2020-06-30T08:31:41.106Z',
        updatedAt: '2020-06-30T08:31:41.106Z',
      },
    }

    const { getAllByText } = renderWithProviders(
      <AccessRequestContext.Provider
        value={{ requestSentTo, handleRequestAccess }}
      >
        <DirectoryDetailView {...serviceWithFailedRequest} />
      </AccessRequestContext.Provider>,
    )

    const failedMessages = getAllByText('Request could not be sent')
    // const retryButton = getByText('Try again')

    expect(failedMessages).toHaveLength(2)
    // expect(retryButton).toBeInTheDocument()

    // fireEvent.click(retryButton)
    // expect(handleRequestAccess).toHaveBeenCalled()
  })
})
