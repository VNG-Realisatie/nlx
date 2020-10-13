// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { act } from '@testing-library/react'
import { renderWithProviders, fireEvent } from '../../../../../test-utils'
import AccessGrantSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('should show if there are no access grants', async () => {
  const { getByTestId } = renderWithProviders(
    <AccessGrantSection accessGrants={[]} />,
  )

  const toggler = getByTestId('service-accessgrants')
  fireEvent.click(toggler)
  jest.runAllTimers()

  expect(toggler).toHaveTextContent(
    'checkbox-multiple.svg' + 'Organizations with access' + '0', // eslint-disable-line no-useless-concat
  )
  expect(getByTestId('service-no-accessgrants')).toBeTruthy()
})

test('should list access grants', async () => {
  const { getByTestId, getByText } = renderWithProviders(
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
  expect(getByTestId('service-accessgrant-list')).toBeTruthy()
  expect(getByText('Organization-B')).toBeInTheDocument()
})

test('revoking access grants', async () => {
  global.confirm = jest.fn(() => true)

  const revokeAccessGrantHandler = jest.fn()

  const { getByTestId, getByText, getByRole } = renderWithProviders(
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
      revokeAccessGrantHandler={revokeAccessGrantHandler}
    />,
  )

  const toggler = getByTestId('service-accessgrants')
  fireEvent.click(toggler)
  jest.runAllTimers()

  const revokeAccessButton = getByText('Revoke')
  expect(revokeAccessButton).toBeInTheDocument()

  await act(async () => fireEvent.click(revokeAccessButton))

  expect(revokeAccessGrantHandler).toHaveBeenCalledWith({
    organizationName: 'Organization-B',
    serviceName: 'service',
    accessGrantId: '1234abcd',
  })

  expect(getByRole('alert')).toBeTruthy()
  expect(getByRole('alert').textContent).toBe('Access revoked')

  global.confirm.mockRestore()
})
