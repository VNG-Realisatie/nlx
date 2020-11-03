// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { makeAutoObservable } from 'mobx'

import { fireEvent, renderWithProviders } from '../../../../../test-utils'
import DirectoryDetailView from './index'

describe('detail view of directory service we do not have access to', () => {
  let service

  global.confirm = jest.fn(() => true)

  beforeEach(() => {
    service = makeAutoObservable({
      organizationName: 'Organization',
      latestAccessRequest: null,
      requestAccess: jest.fn(),
    })
  })

  it('should show a failed message', () => {
    service.latestAccessRequest = {
      id: 'string',
      organizationName: 'organization',
      serviceName: 'service',
      state: 'FAILED',
      createdAt: new Date('2020-06-30T08:31:41.106Z'),
      updatedAt: new Date('2020-06-30T08:31:41.106Z'),
      errorDetails: {
        cause: 'Something went wrong',
        stackTrace: ['Go main panic'],
      },
    }

    const { getAllByText } = renderWithProviders(
      <DirectoryDetailView service={service} />,
    )

    const failedMessages = getAllByText('Request could not be sent')

    expect(failedMessages).toHaveLength(2)
  })

  it('should show show-trace button', () => {
    service.latestAccessRequest = {
      id: 'string',
      organizationName: 'organization',
      serviceName: 'service',
      state: 'FAILED',
      createdAt: new Date('2020-06-30T08:31:41.106Z'),
      updatedAt: new Date('2020-06-30T08:31:41.106Z'),
      errorDetails: {
        cause: 'Something went wrong',
        stackTrace: ['Go main panic', 'main.go:10'],
      },
    }

    const { getByText, getByTestId } = renderWithProviders(
      <DirectoryDetailView service={service} />,
    )

    const button = getByText('Show stacktrace')
    fireEvent.click(button)

    const drawer = getByTestId('stacktrace')
    const pre = getByTestId('stacktrace-content')

    expect(drawer).toBeVisible()
    expect(pre.innerHTML).toContain('Go main panic<br>main.go:10')
  })
})
