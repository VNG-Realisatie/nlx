// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders, fireEvent, act } from '../../../../../test-utils'
import AccessGrantSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('should show if there are no access grants', async () => {
  const { getByTestId, getByText } = renderWithProviders(
    <AccessGrantSection accessGrants={[]} />,
  )

  const toggler = getByTestId('service-accessgrants')
  fireEvent.click(toggler)
  jest.runAllTimers()

  expect(toggler).toHaveTextContent(
    'checkbox-multiple.svg' + 'Organizations with access' + '0', // eslint-disable-line no-useless-concat
  )
  expect(
    getByText('There are no organizations with access'),
  ).toBeInTheDocument()
})

test('should list access grants', async () => {
  const { getByTestId } = renderWithProviders(
    <AccessGrantSection
      accessGrants={[
        {
          id: '1234abcd',
          serviceName: 'service',
          organizationName: 'Organization-B',
          publicKeyFingerprint: 'printFinger=',
          createdAt: '2020-10-07T13:01:11.288349Z',
        },
      ]}
    />,
  )

  const toggler = getByTestId('service-accessgrants')
  fireEvent.click(toggler)
  jest.runAllTimers()

  expect(toggler).toHaveTextContent(
    'checkbox-multiple.svg' + 'Organizations with access' + '1', // eslint-disable-line no-useless-concat
  )
  expect(getByTestId('service-accessgrant-list')).toBeInTheDocument()
})

test('revoking access', async () => {
  const mockRevoke = jest.fn().mockResolvedValue()
  const { getByTestId, getByText, queryByRole } = renderWithProviders(
    <AccessGrantSection
      accessGrants={[
        {
          id: '1234abcd',
          serviceName: 'service',
          organizationName: 'Organization-B',
          publicKeyFingerprint: 'printFinger=',
          createdAt: new Date('2020-10-01T12:00:00Z'),
          revokedAt: null,
          revoke: mockRevoke,
          error: '',
        },
      ]}
    />,
  )

  const toggler = getByTestId('service-accessgrants')
  fireEvent.click(toggler)
  jest.runAllTimers()

  global.confirm = jest.fn(() => true)

  const revokeButton = getByText('Revoke')

  await act(async () => {
    await fireEvent.click(revokeButton)
  })

  expect(queryByRole('alert')).toHaveTextContent('Access revoked')
})
