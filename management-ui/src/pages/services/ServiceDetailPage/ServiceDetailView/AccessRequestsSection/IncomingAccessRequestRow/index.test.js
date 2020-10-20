// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders, fireEvent } from '../../../../../../test-utils'
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
    createdAt: '2020-08-25T13:30:43.480155Z',
    updatedAt: '2020-08-25T13:30:43.480155Z',
  }
})

test('requesting access will fire approve handler', async () => {
  const { getByTestId } = renderWithProviders(
    <table>
      <tbody>
        <IncomingAccessRequestRow
          accessRequest={accessRequest}
          approveHandler={mockHandler}
        />
      </tbody>
    </table>,
  )

  global.confirm = jest.fn(() => true)

  const approveButton = getByTestId('button-approve')
  fireEvent.click(approveButton)

  expect(mockHandler).toHaveBeenCalledWith(accessRequest)
})

test('rejecting access will fire reject handler', async () => {
  const { getByTestId } = renderWithProviders(
    <table>
      <tbody>
        <IncomingAccessRequestRow
          accessRequest={accessRequest}
          rejectHandler={mockHandler}
        />
      </tbody>
    </table>,
  )

  global.confirm = jest.fn(() => true)

  const rejectButton = getByTestId('button-reject')
  fireEvent.click(rejectButton)

  expect(mockHandler).toHaveBeenCalledWith(accessRequest)
})

test('clicking cancel will not fire handler', async () => {
  const { getByTestId } = renderWithProviders(
    <table>
      <tbody>
        <IncomingAccessRequestRow
          accessRequest={accessRequest}
          approveHandler={mockHandler}
        />
      </tbody>
    </table>,
  )

  global.confirm = jest.fn(() => false)

  const approveButton = getByTestId('button-approve')
  fireEvent.click(approveButton)

  expect(mockHandler).not.toHaveBeenCalled()
})
