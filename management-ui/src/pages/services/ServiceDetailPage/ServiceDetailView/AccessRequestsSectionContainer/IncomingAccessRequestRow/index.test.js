// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import { clickConfirmButtonAndAssert } from '../../../../../../components/ConfirmationModal/testUtils'
import IncomingAccessRequestRow from './index'

let mockHandler
let accessRequest

beforeEach(() => {
  mockHandler = jest.fn()
  accessRequest = {
    id: '1a2B',
    organizationName: 'Organization A',
    serviceName: 'Servicio',
    state: 'RECEIVED',
    createdAt: new Date('2020-10-01T12:00:00Z'),
    updatedAt: new Date('2020-10-01T12:00:01Z'),
  }
})

test('requesting access will fire approve handler', async () => {
  const { getByTitle, getByText } = renderWithProviders(
    <table>
      <tbody>
        <IncomingAccessRequestRow
          accessRequest={accessRequest}
          approveHandler={mockHandler}
          rejectHandler={jest.fn()}
        />
      </tbody>
    </table>,
  )

  fireEvent.click(getByTitle('Approve'))

  await clickConfirmButtonAndAssert(getByText('Approve'), () =>
    expect(mockHandler).toHaveBeenCalled(),
  )
})

test('rejecting access will fire reject handler', async () => {
  const { getByTitle, getByText } = renderWithProviders(
    <table>
      <tbody>
        <IncomingAccessRequestRow
          accessRequest={accessRequest}
          approveHandler={jest.fn()}
          rejectHandler={mockHandler}
        />
      </tbody>
    </table>,
  )

  fireEvent.click(getByTitle('Reject'))

  await clickConfirmButtonAndAssert(getByText('Reject'), () =>
    expect(mockHandler).toHaveBeenCalled(),
  )
})

test('clicking cancel will not fire handler', async () => {
  const { getByTitle } = renderWithProviders(
    <table>
      <tbody>
        <IncomingAccessRequestRow
          accessRequest={accessRequest}
          approveHandler={mockHandler}
          rejectHandler={jest.fn()}
        />
      </tbody>
    </table>,
  )

  global.confirm = jest.fn(() => false)

  const approveButton = getByTitle('Approve')
  fireEvent.click(approveButton)

  expect(mockHandler).not.toHaveBeenCalled()
})
