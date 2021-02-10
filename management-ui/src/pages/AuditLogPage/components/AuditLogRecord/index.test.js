// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import {
  AUDIT_LOG_ACTION_LOGIN_FAIL,
  AUDIT_LOG_ACTION_LOGIN_SUCCESS,
  AUDIT_LOG_ACTION_LOGOUT_SUCCESS,
  AUDIT_LOG_ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
} from '../../../../stores/models/AuditLogModel'
import AuditLogRecord from './index'

test.concurrent.each([
  [AUDIT_LOG_ACTION_LOGIN_SUCCESS, 'John Doe has logged in'],
  [AUDIT_LOG_ACTION_LOGIN_FAIL, 'failed login attempt'],
  [AUDIT_LOG_ACTION_LOGOUT_SUCCESS, 'John Doe has logged out'],
  [
    AUDIT_LOG_ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
    'John Doe has accepted the access request from Gemeente Haarlem for Kadaster',
  ],
  ['unknown action', "John Doe has performed unknown action 'unknown action'"],
])('AuditLogRecord message for action %s', (action, expectedMessage) => {
  const { getByTestId } = renderWithProviders(
    <AuditLogRecord action={action} user="John Doe" />,
  )
  expect(getByTestId('message')).toHaveTextContent(expectedMessage)
})
