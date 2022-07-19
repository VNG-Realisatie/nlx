// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, within, waitFor, screen, act } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import { configure } from 'mobx'
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
  configure({ safeDescriptors: false })

  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData: {
      id: '1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
      state: STATES.RECEIVED,
      createdAt: '2020-10-01T12:00:00Z',
      updatedAt: '2020-10-01T12:00:01Z',
      publicKeyFingerprint: 'public-key-fingerprint',
    },
  })

  const approveSpy = jest
    .spyOn(accessRequest, 'approve')
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  const onApproveOrRejectHandler = jest.fn()

  renderWithAllProviders(
    <MemoryRouter>
      <CollapsibleBody
        accessRequests={[accessRequest]}
        onApproveOrRejectCallbackHandler={onApproveOrRejectHandler}
      />
    </MemoryRouter>,
  )

  act(() => {
    fireEvent.click(screen.getByTitle('Approve'))
  })

  let confirmModal = screen.getByRole('dialog')
  let okButton = within(confirmModal).getByText('Approve')
  fireEvent.click(okButton)

  await waitFor(() => expect(approveSpy).toHaveBeenCalled())

  expect(screen.queryByRole('alert')).toHaveTextContent(
    "Failed to approve access requestYou don't have the required permission.",
  )

  act(() => {
    fireEvent.click(screen.getByTitle('Approve'))
  })

  confirmModal = screen.getByRole('dialog')
  okButton = within(confirmModal).getByText('Approve')
  fireEvent.click(okButton)

  // toast
  expect(await screen.findByText('Access request approved')).toBeInTheDocument()

  expect(onApproveOrRejectHandler).toHaveBeenCalledTimes(1)
})

test('rejecting an incoming access request', async () => {
  configure({ safeDescriptors: false })

  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData: {
      id: '1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
      state: STATES.RECEIVED,
      createdAt: '2020-10-01T12:00:00Z',
      updatedAt: '2020-10-01T12:00:01Z',
      publicKeyFingerprint: 'public-key-fingerprint',
    },
  })

  const rejectSpy = jest
    .spyOn(accessRequest, 'reject')
    .mockRejectedValueOnce({
      response: {
        status: 403,
      },
    })
    .mockResolvedValue()

  const onApproveOrRejectHandler = jest.fn()
  const { getByTitle, getByRole, findByText } = renderWithAllProviders(
    <MemoryRouter>
      <CollapsibleBody
        accessRequests={[accessRequest]}
        onApproveOrRejectCallbackHandler={onApproveOrRejectHandler}
      />
    </MemoryRouter>,
  )

  fireEvent.click(getByTitle('Reject'))

  let confirmModal = getByRole('dialog')
  let okButton = within(confirmModal).getByText('Reject')
  fireEvent.click(okButton)

  await waitFor(() => expect(rejectSpy).toHaveBeenCalled())

  expect(screen.queryByRole('alert')).toHaveTextContent(
    "Failed to reject access requestYou don't have the required permission.",
  )

  fireEvent.click(getByTitle('Reject'))

  confirmModal = getByRole('dialog')
  okButton = within(confirmModal).getByText('Reject')
  fireEvent.click(okButton)

  // toast
  expect(await findByText('Access request rejected')).toBeInTheDocument()

  expect(onApproveOrRejectHandler).toHaveBeenCalledTimes(1)
})
