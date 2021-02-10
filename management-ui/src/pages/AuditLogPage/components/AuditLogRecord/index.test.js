// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import {
  ACTION_ACCESS_GRANT_REVOKE,
  ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
  ACTION_INCOMING_ACCESS_REQUEST_REJECT,
  ACTION_LOGIN_FAIL,
  ACTION_LOGIN_SUCCESS,
  ACTION_LOGOUT_SUCCESS,
  ACTION_ORGANIZATION_SETTINGS_UPDATE,
  ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
  ACTION_SERVICE_CREATE,
  ACTION_SERVICE_DELETE,
  ACTION_SERVICE_UPDATE,
} from '../../../../stores/models/AuditLogModel'
import AuditLogRecord from './index'

test.concurrent.each([
  [{ action: ACTION_LOGIN_SUCCESS }, 'John Doe has logged in'],
  [{ action: ACTION_LOGIN_FAIL }, 'Failed login attempt'],
  [{ action: ACTION_LOGOUT_SUCCESS }, 'John Doe has logged out'],
  [
    {
      action: ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'John Doe has accepted the access request from Gemeente Haarlem for Kadaster',
  ],
  [
    {
      action: ACTION_INCOMING_ACCESS_REQUEST_REJECT,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'John Doe has rejected the access request from Gemeente Haarlem for Kadaster',
  ],
  [
    {
      action: ACTION_ACCESS_GRANT_REVOKE,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'John Doe has revoked access for Kadaster from Gemeente Haarlem',
  ],
  [
    {
      action: ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'John Doe has requested access for Kadaster from Gemeente Haarlem',
  ],
  [
    {
      action: ACTION_SERVICE_CREATE,
      service: 'Kadaster',
    },
    'John Doe has created the service Kadaster',
  ],
  [
    {
      action: ACTION_SERVICE_UPDATE,
      service: 'Kadaster',
    },
    'John Doe has updated the service Kadaster',
  ],
  [
    {
      action: ACTION_SERVICE_DELETE,
      service: 'Kadaster',
    },
    'John Doe has removed the service Kadaster',
  ],
  [
    {
      action: ACTION_ORGANIZATION_SETTINGS_UPDATE,
    },
    'John Doe updated the organization settings',
  ],
  [
    { action: 'unknown action' },
    "John Doe has performed unknown action 'unknown action'",
  ],
])('AuditLogRecord message for audit log %s', (auditLog, expectedMessage) => {
  const { getByTestId } = renderWithProviders(
    <AuditLogRecord {...auditLog} user="John Doe" />,
  )
  expect(getByTestId('message')).toHaveTextContent(expectedMessage)
})
