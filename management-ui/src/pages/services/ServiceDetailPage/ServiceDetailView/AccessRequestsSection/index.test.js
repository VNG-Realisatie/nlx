// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders, fireEvent } from '../../../../../test-utils'
import AccessRequestsSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('listing the access requests', async () => {
  const onApproveOrRejectCallbackHandler = jest.fn().mockResolvedValue(null)

  const { getByTestId, rerender, getByText } = renderWithProviders(
    <AccessRequestsSection
      accessRequests={[]}
      onApproveOrRejectCallbackHandler={onApproveOrRejectCallbackHandler}
    />,
  )

  const toggler = getByTestId('service-incoming-accessrequests')
  fireEvent.click(toggler)

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '0', // eslint-disable-line no-useless-concat
  )

  expect(
    getByTestId('service-incoming-accessrequests-amount'),
  ).toBeInTheDocument()
  expect(getByText('There are no access requests')).toBeInTheDocument()

  const accessRequest = {
    id: '1a2B',
    organizationName: 'Organization A',
    serviceName: 'Servicio',
    state: 'RECEIVED',
    createdAt: new Date('2020-10-01T12:00:00Z'),
    updatedAt: new Date('2020-10-01T12:00:01Z'),
    approve: jest.fn().mockResolvedValue(null),
    reject: jest.fn().mockResolvedValue(null),
  }

  rerender(
    <AccessRequestsSection
      accessRequests={[accessRequest]}
      onApproveOrRejectCallbackHandler={onApproveOrRejectCallbackHandler}
    />,
  )

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '1', // eslint-disable-line no-useless-concat
  )

  expect(
    getByTestId('service-incoming-accessrequests-amount-accented'),
  ).toBeTruthy()
  expect(
    getByTestId('service-incoming-accessrequests-list'),
  ).toBeInTheDocument()
  expect(getByText('Organization A')).toBeInTheDocument()
})
