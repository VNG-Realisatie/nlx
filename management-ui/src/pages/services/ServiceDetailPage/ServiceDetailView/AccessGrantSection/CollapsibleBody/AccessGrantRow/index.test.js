// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, within, waitFor } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../../test-utils'
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
  const { getByText, getByRole } = renderWithProviders(
    <table>
      <tbody>
        <AccessGrantRow accessGrant={accessGrant} revokeHandler={mockHandler} />
      </tbody>
    </table>,
  )

  fireEvent.click(getByText('Revoke'))

  const confirmModal = getByRole('dialog')
  const okButton = within(confirmModal).getByText('Revoke')

  fireEvent.click(okButton)
  await waitFor(() => expect(mockHandler).toHaveBeenCalledWith(accessGrant))
})

test('clicking confirm cancel', async () => {
  const { getByText } = renderWithProviders(
    <table>
      <tbody>
        <AccessGrantRow accessGrant={accessGrant} revokeHandler={mockHandler} />
      </tbody>
    </table>,
  )

  fireEvent.click(getByText('Revoke'))
  fireEvent.click(getByText('Cancel'))
  await waitFor(() => expect(mockHandler).not.toHaveBeenCalled())
})
