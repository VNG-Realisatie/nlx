// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { fireEvent, waitForElementToBeRemoved } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import IncomingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../../../../../stores/models/IncomingAccessRequestModel'
import CollapsibleBody from './index'

test('when no access requests are available', async () => {
  const { getByText } = renderWithProviders(
    <CollapsibleBody accessRequests={[]} />,
  )
  expect(getByText('There are no access requests')).toBeInTheDocument()
})

test('listing the access requests', async () => {
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

  accessRequest.approve = jest.fn().mockResolvedValue()
  accessRequest.reject = jest.fn().mockResolvedValue()

  const onApproveOrRejectHandler = jest.fn().mockResolvedValue()

  const {
    getByTestId,
    getByText,
    getByTitle,
    findByText,
  } = renderWithProviders(
    <CollapsibleBody
      accessRequests={[accessRequest]}
      onApproveOrRejectCallbackHandler={onApproveOrRejectHandler}
    />,
  )

  expect(
    getByTestId('service-incoming-accessrequests-list'),
  ).toBeInTheDocument()
  expect(getByText('organization-a')).toBeInTheDocument()

  fireEvent.click(getByTitle('Reject'))

  // confirm rejection
  const confirmButton = await findByText('Reject')
  fireEvent.click(confirmButton)
  await waitForElementToBeRemoved(confirmButton)

  expect(onApproveOrRejectHandler).toHaveBeenCalledTimes(1)
})
