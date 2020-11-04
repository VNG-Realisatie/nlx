// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders, fireEvent } from '../../../../../../test-utils'
import AccessGrantRow from './index'

let mockHandler
let accessGrant

beforeEach(() => {
  mockHandler = jest.fn()
  accessGrant = {
    id: 'abc',
    organizationName: 'Organization A',
    serviceName: 'Servicio',
    publicKeyFingerprint: 'abc',
    createdAt: new Date('2020-10-01T12:00:00Z'),
    updatedAt: new Date('2020-10-01T12:00:01Z'),
  }
})

test('clicking confirm', async () => {
  const { getByText } = renderWithProviders(
    <table>
      <tbody>
        <AccessGrantRow accessGrant={accessGrant} revokeHandler={mockHandler} />
      </tbody>
    </table>,
  )

  global.confirm = jest.fn(() => true)

  const revokeButton = getByText('Revoke')
  fireEvent.click(revokeButton)

  expect(mockHandler).toHaveBeenCalledWith(accessGrant)
})

test('clicking confirm cancel', async () => {
  const { getByText } = renderWithProviders(
    <table>
      <tbody>
        <AccessGrantRow accessGrant={accessGrant} revokeHandler={mockHandler} />
      </tbody>
    </table>,
  )

  global.confirm = jest.fn(() => false)

  const revokeButton = getByText('Revoke')
  fireEvent.click(revokeButton)

  expect(mockHandler).not.toHaveBeenCalled()
})
