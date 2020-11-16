// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { act } from '@testing-library/react'
import { renderWithProviders, fireEvent } from '../../../../../test-utils'
import AccessRequestsSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('listing the access requests', async () => {
  global.confirm = jest.fn(() => true)
  const fetchServiceHandler = jest.fn().mockResolvedValue(null)

  const { getByTestId, rerender, getByText } = renderWithProviders(
    <AccessRequestsSection
      accessRequests={[]}
      fetchServiceHandler={fetchServiceHandler}
    />,
  )

  const toggler = getByTestId('service-incoming-accessrequests')
  fireEvent.click(toggler)

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '0', // eslint-disable-line no-useless-concat
  )

  expect(getByTestId('service-incoming-accessrequests-amount')).toBeTruthy()
  expect(getByTestId('service-no-incoming-accessrequests')).toBeTruthy()

  rerender(
    <AccessRequestsSection
      accessRequests={[
        {
          id: '1a2B',
          organizationName: 'Organization A',
          serviceName: 'Servicio',
          state: 'RECEIVED',
          createdAt: new Date('2020-10-01T12:00:00Z'),
          updatedAt: new Date('2020-10-01T12:00:01Z'),
          approve: () => Promise.resolve(),
          reject: () => Promise.resolve(),
        },
      ]}
      fetchServiceHandler={fetchServiceHandler}
    />,
  )

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '1', // eslint-disable-line no-useless-concat
  )

  expect(
    getByTestId('service-incoming-accessrequests-amount-accented'),
  ).toBeTruthy()
  expect(getByTestId('service-incoming-accessrequests-list')).toBeTruthy()
  expect(getByText('Organization A')).toBeInTheDocument()

  await act(async () => {
    await fireEvent.click(getByTestId('button-approve'))
  })

  expect(fetchServiceHandler).toHaveBeenCalledTimes(1)

  await act(async () => {
    await fireEvent.click(getByTestId('button-reject'))
  })

  expect(fetchServiceHandler).toHaveBeenCalledTimes(2)
})
