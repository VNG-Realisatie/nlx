// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import {
  fireEvent,
  waitForElementToBeRemoved,
  within,
} from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithProviders } from '../../../../../../test-utils'
import AccessGrantModel from '../../../../../../stores/models/AccessGrantModel'
import CollapsibleBody from './index'

test('when no access grants are available', async () => {
  const { getByText } = renderWithProviders(
    <CollapsibleBody accessGrants={[]} />,
  )
  expect(
    getByText('There are no organizations with access'),
  ).toBeInTheDocument()
})

test('listing the access grants', async () => {
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

  accessGrant.revoke = jest.fn().mockResolvedValue()

  const { getByTestId, getByText, getByTitle, getByRole } = renderWithProviders(
    <CollapsibleBody accessGrants={[accessGrant]} />,
  )

  expect(getByTestId('service-accessgrant-list')).toBeInTheDocument()
  expect(getByText('organization-a')).toBeInTheDocument()

  fireEvent.click(getByTitle('Revoke'))

  // confirm revoke
  const confirmButton = await within(getByRole('dialog')).findByText('Revoke')
  fireEvent.click(confirmButton)
  await waitForElementToBeRemoved(confirmButton)

  expect(accessGrant.revoke).toHaveBeenCalledTimes(1)
})
