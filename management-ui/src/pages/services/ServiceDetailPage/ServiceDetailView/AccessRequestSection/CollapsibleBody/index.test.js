// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, within, waitFor } from '@testing-library/react'
import { renderWithAllProviders } from '../../../../../../test-utils'
import IncomingAccessRequestModel, {
  STATES,
} from '../../../../../../stores/models/IncomingAccessRequestModel'
import CollapsibleBody from './index'

jest.mock('../../../../../../components/Modal')

test('when no access requests are available', async () => {
  const { getByText } = renderWithAllProviders(
    <CollapsibleBody accessRequests={[]} />,
  )
  expect(getByText('There are no access requests')).toBeInTheDocument()
})

test('approving an incoming access request', async () => {
  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData: {
      id: '1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
      state: STATES.RECEIVED,
      createdAt: '2020-10-01T12:00:00Z',
      updatedAt: '2020-10-01T12:00:01Z',
    },
  })

  const approveSpy = jest.fn()
  accessRequest.approve = approveSpy

  const onApproveOrRejectHandler = jest.fn()
  const { getByTitle, getByRole, findByText } = renderWithAllProviders(
    <CollapsibleBody
      accessRequests={[accessRequest]}
      onApproveOrRejectCallbackHandler={onApproveOrRejectHandler}
    />,
  )

  fireEvent.click(getByTitle('Approve'))

  const confirmModal = getByRole('dialog')
  const okButton = within(confirmModal).getByText('Approve')
  fireEvent.click(okButton)

  await waitFor(() => expect(approveSpy).toHaveBeenCalled())
  expect(onApproveOrRejectHandler).toHaveBeenCalledTimes(1)

  // toast
  expect(await findByText('Access request approved')).toBeInTheDocument()
})

test('rejecting an incoming access request', async () => {
  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData: {
      id: '1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
      state: STATES.RECEIVED,
      createdAt: '2020-10-01T12:00:00Z',
      updatedAt: '2020-10-01T12:00:01Z',
    },
  })

  const rejectSpy = jest.fn()
  accessRequest.reject = rejectSpy

  const onApproveOrRejectHandler = jest.fn()
  const { getByTitle, getByRole, findByText } = renderWithAllProviders(
    <CollapsibleBody
      accessRequests={[accessRequest]}
      onApproveOrRejectCallbackHandler={onApproveOrRejectHandler}
    />,
  )

  fireEvent.click(getByTitle('Reject'))

  const confirmModal = getByRole('dialog')
  const okButton = within(confirmModal).getByText('Reject')
  fireEvent.click(okButton)

  await waitFor(() => expect(rejectSpy).toHaveBeenCalled())
  expect(onApproveOrRejectHandler).toHaveBeenCalledTimes(1)

  // toast
  expect(await findByText('Access request rejected')).toBeInTheDocument()
})
