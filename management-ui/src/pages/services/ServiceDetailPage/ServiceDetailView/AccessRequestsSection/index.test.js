// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders, fireEvent } from '../../../../../test-utils'
import AccessRequestsSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('should show if there are no incoming access requests', async () => {
  const { getByTestId } = renderWithProviders(
    <AccessRequestsSection accessRequests={[]} />,
  )

  const toggler = getByTestId('service-incoming-accessrequests')

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '0', // eslint-disable-line no-useless-concat
  )

  fireEvent.click(toggler)

  expect(getByTestId('service-no-incoming-accessrequests')).toBeTruthy()
})

test('should list access requests', async () => {
  const { getByTestId } = renderWithProviders(
    <AccessRequestsSection
      accessRequests={[
        {
          id: '1a2B',
          organizationName: 'Organization A',
          serviceName: 'Servicio',
          state: 'RECEIVED',
          createdAt: '2020-08-25T13:30:43.480155Z',
          updatedAt: '2020-08-25T13:30:43.480155Z',
        },
      ]}
    />,
  )

  const toggler = getByTestId('service-incoming-accessrequests')

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '1', // eslint-disable-line no-useless-concat
  )

  fireEvent.click(toggler)
  jest.runAllTimers()

  expect(getByTestId('service-incoming-accessrequests-list')).toBeTruthy()
  expect(getByTestId('service-incoming-accessrequest-1a2B')).toHaveTextContent(
    'Organization A',
  )
})
