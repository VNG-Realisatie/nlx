// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { waitFor, fireEvent, within, screen } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../../../../../test-utils'
import AccessGrantModel from '../../../../../../stores/models/AccessGrantModel'
import CollapsibleBody from './index'

jest.mock('../../../../../../components/Modal')

test('when no access grants are available', async () => {
  const { getByText } = renderWithAllProviders(
    <MemoryRouter>
      <CollapsibleBody accessGrants={[]} />
    </MemoryRouter>,
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
      organization: {
        name: 'organization-a',
        serialNumber: '00000000000000000001',
      },
      createdAt: new Date(),
      updatedAt: new Date(),
    },
  })

  const { getByText, findByText, getByRole } = renderWithAllProviders(
    <MemoryRouter>
      <CollapsibleBody accessGrants={[accessGrant]} />
    </MemoryRouter>,
  )

  const revokeSpy = jest
    .spyOn(accessGrant, 'revoke')
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  fireEvent.click(getByText('Revoke'))

  let confirmModal = getByRole('dialog')
  const cancelButton = within(confirmModal).getByText('Cancel')
  fireEvent.click(cancelButton)

  await waitFor(() => expect(revokeSpy).not.toHaveBeenCalled())

  fireEvent.click(getByText('Revoke'))

  confirmModal = getByRole('dialog')
  let okButton = within(confirmModal).getByText('Revoke')
  fireEvent.click(okButton)

  await waitFor(() => expect(revokeSpy).toHaveBeenCalled())

  expect(screen.queryByRole('alert')).toHaveTextContent(
    "Failed to revoke access grantYou don't have the required permission.",
  )

  fireEvent.click(getByText('Revoke'))

  confirmModal = screen.getByRole('dialog')
  okButton = within(confirmModal).getByText('Revoke')
  fireEvent.click(okButton)

  // toast
  expect(await findByText('Access revoked')).toBeInTheDocument()
})
