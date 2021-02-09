// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import {
  AUDIT_LOG_ACTION_LOGIN_SUCCESS,
  AUDIT_LOG_ACTION_LOGOUT_SUCCESS,
} from '../../../../stores/models/AuditLogModel'
import AuditLogRecord from './index'

test.concurrent.each([
  [AUDIT_LOG_ACTION_LOGIN_SUCCESS, 'John Doe has logged in'],
  [AUDIT_LOG_ACTION_LOGOUT_SUCCESS, 'John Doe has logged out'],
  ['unknown action', "John Doe has performed unknown action 'unknown action'"],
])('AuditLogRecord message for action %s', (action, expectedMessage) => {
  const { getByTestId } = renderWithProviders(
    <AuditLogRecord action={action} user="John Doe" />,
  )
  expect(getByTestId('message')).toHaveTextContent(expectedMessage)
})
