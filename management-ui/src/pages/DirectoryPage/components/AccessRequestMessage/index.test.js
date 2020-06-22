// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../../test-utils'
import AccessRequestMessage from './index'

test('by default should render nothing', () => {
  const { container } = renderWithProviders(<AccessRequestMessage />)
  expect(container).toHaveTextContent('')
})

test('render fallback status (CREATED) when no latestAccessRequest given', () => {
  const { getByText } = renderWithProviders(
    <AccessRequestMessage fallbackStatus="CREATED" />,
  )
  expect(getByText('Sending request')).toBeInTheDocument()
})

describe('render correct message', () => {
  it('for state: FAILED', () => {
    const latestAccessRequest = {
      id: 'id',
      status: 'FAILED',
      createdAt: '2020-10-1T12:06:02Z',
      updatedAt: '2020-10-1T12:06:04Z',
    }
    const { getByText } = renderWithProviders(
      <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
    )

    expect(getByText('Request could not be sent')).toBeInTheDocument()
  })

  it('for state: SENT', () => {
    const latestAccessRequest = {
      id: 'id',
      status: 'SENT',
      createdAt: '2020-10-1T12:06:02Z',
      updatedAt: '2020-10-1T12:06:04Z',
    }
    const { getByText } = renderWithProviders(
      <AccessRequestMessage latestAccessRequest={latestAccessRequest} />,
    )

    expect(getByText('Requested')).toBeInTheDocument()
  })
})
