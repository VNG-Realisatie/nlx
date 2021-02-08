// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, within, waitFor } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithProviders } from '../../../../../../test-utils'
import IncomingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../../../../../stores/models/IncomingAccessRequestModel'
import CollapsibleBody from './index'

jest.mock('../../../../../../components/Modal')

test('when no access requests are available', async () => {
  const { getByText } = renderWithProviders(
    <CollapsibleBody accessRequests={[]} />,
  )
  expect(getByText('There are no access requests')).toBeInTheDocument()
})

test('approving an incoming access request', async () => {
  configure({ safeDescriptors: false })

  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData: {
      id: '1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
      state: ACCESS_REQUEST_STATES.RECEIVED,
      createdAt: new Date(),
      updatedAt: new Date(),
    },
  })

  const onApproveOrRejectHandler = jest.fn().mockResolvedValue()
  const { getByTitle, getByRole, findByText } = renderWithProviders(
    <CollapsibleBody
      accessRequests={[accessRequest]}
      onApproveOrRejectCallbackHandler={onApproveOrRejectHandler}
    />,
  )

  const approveSpy = jest.spyOn(accessRequest, 'approve').mockResolvedValue()
  fireEvent.click(getByTitle('Approve'))

  let confirmModal = getByRole('dialog')
  const cancelButton = within(confirmModal).getByText('Cancel')
  fireEvent.click(cancelButton)

  await waitFor(() => expect(approveSpy).not.toHaveBeenCalled())

  fireEvent.click(getByTitle('Approve'))

  confirmModal = getByRole('dialog')
  const okButton = within(confirmModal).getByText('Approve')
  fireEvent.click(okButton)

  await waitFor(() => expect(approveSpy).toHaveBeenCalled())
  expect(onApproveOrRejectHandler).toHaveBeenCalledTimes(1)

  // toast
  expect(await findByText('Access request approved')).toBeInTheDocument()
})

test('rejecting an incoming access request', async () => {
  configure({ safeDescriptors: false })

  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData: {
      id: '1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
      state: ACCESS_REQUEST_STATES.RECEIVED,
      createdAt: new Date(),
      updatedAt: new Date(),
    },
  })

  const onApproveOrRejectHandler = jest.fn().mockResolvedValue()
  const { getByTitle, getByRole, findByText } = renderWithProviders(
    <CollapsibleBody
      accessRequests={[accessRequest]}
      onApproveOrRejectCallbackHandler={onApproveOrRejectHandler}
    />,
  )

  const rejectSpy = jest.spyOn(accessRequest, 'reject').mockResolvedValue()
  fireEvent.click(getByTitle('Reject'))

  let confirmModal = getByRole('dialog')
  const cancelButton = within(confirmModal).getByText('Cancel')
  fireEvent.click(cancelButton)

  await waitFor(() => expect(rejectSpy).not.toHaveBeenCalled())

  fireEvent.click(getByTitle('Reject'))

  confirmModal = getByRole('dialog')
  const okButton = within(confirmModal).getByText('Reject')
  fireEvent.click(okButton)

  await waitFor(() => expect(rejectSpy).toHaveBeenCalled())
  expect(onApproveOrRejectHandler).toHaveBeenCalledTimes(1)

  // toast
  expect(await findByText('Access request rejected')).toBeInTheDocument()
})
