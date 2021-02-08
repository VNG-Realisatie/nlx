// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { waitFor, fireEvent, within } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithProviders } from '../../../../../../test-utils'
import AccessGrantModel from '../../../../../../stores/models/AccessGrantModel'
import CollapsibleBody from './index'

jest.mock('../../../../../../components/Modal')

test('when no access grants are available', async () => {
  const { getByText } = renderWithProviders(
    <CollapsibleBody accessGrants={[]} />,
  )
  expect(
    getByText('There are no organizations with access'),
  ).toBeInTheDocument()
})

test('revoke access grant', async () => {
  configure({ safeDescriptors: false })

  const accessGrant = new AccessGrantModel({
    accessGrantData: {
      id: '1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
      createdAt: new Date(),
      updatedAt: new Date(),
    },
  })

  const { getByText, findByText, getByRole } = renderWithProviders(
    <CollapsibleBody accessGrants={[accessGrant]} />,
  )

  const revokeSpy = jest.spyOn(accessGrant, 'revoke').mockResolvedValue()
  fireEvent.click(getByText('Revoke'))

  let confirmModal = getByRole('dialog')
  const cancelButton = within(confirmModal).getByText('Cancel')
  fireEvent.click(cancelButton)

  await waitFor(() => expect(revokeSpy).not.toHaveBeenCalled())

  fireEvent.click(getByText('Revoke'))

  confirmModal = getByRole('dialog')
  const okButton = within(confirmModal).getByText('Revoke')
  fireEvent.click(okButton)

  await waitFor(() => expect(revokeSpy).toHaveBeenCalled())

  // toast
  expect(await findByText('Access revoked')).toBeInTheDocument()
})
