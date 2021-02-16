// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import {
  ACTION_ACCESS_GRANT_REVOKE,
  ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
  ACTION_INCOMING_ACCESS_REQUEST_REJECT,
  ACTION_INSIGHT_CONFIGURATION_UPDATE,
  ACTION_LOGIN_FAIL,
  ACTION_LOGIN_SUCCESS,
  ACTION_LOGOUT,
  ACTION_ORGANIZATION_SETTINGS_UPDATE,
  ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
  ACTION_SERVICE_CREATE,
  ACTION_SERVICE_DELETE,
  ACTION_SERVICE_UPDATE,
} from '../../../../stores/models/AuditLogModel'
import AuditLogRecord from './index'

test.concurrent.each([
  [{ action: ACTION_LOGIN_SUCCESS }, 'shut-down.svg', 'John Doe has logged in'],
  [{ action: ACTION_LOGIN_FAIL }, 'shut-down.svg', 'Failed login attempt'],
  [
    { action: ACTION_LOGOUT },
    'shut-down.svg',
    'John Doe has logged out',
  ],
  [
    {
      action: ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'check.svg',
    'John Doe has approved the access request from Gemeente Haarlem for Kadaster',
  ],
  [
    {
      action: ACTION_INCOMING_ACCESS_REQUEST_REJECT,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'close.svg',
    'John Doe has rejected the access request from Gemeente Haarlem for Kadaster',
  ],
  [
    {
      action: ACTION_ACCESS_GRANT_REVOKE,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'revoke.svg',
    'John Doe has revoked access for Kadaster from Gemeente Haarlem',
  ],
  [
    {
      action: ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
      organization: 'Gemeente Haarlem',
      service: 'Kadaster',
    },
    'key.svg',
    'John Doe has requested access to Kadaster from Gemeente Haarlem',
  ],
  [
    {
      action: ACTION_SERVICE_CREATE,
      service: 'Kadaster',
    },
    'services.svg',
    'John Doe has created the service Kadaster',
  ],
  [
    {
      action: ACTION_SERVICE_UPDATE,
      service: 'Kadaster',
    },
    'services.svg',
    'John Doe has updated the service Kadaster',
  ],
  [
    {
      action: ACTION_SERVICE_DELETE,
      service: 'Kadaster',
    },
    'services.svg',
    'John Doe has removed the service Kadaster',
  ],
  [
    {
      action: ACTION_ORGANIZATION_SETTINGS_UPDATE,
    },
    'cog.svg',
    'John Doe updated the organization settings',
  ],
  [
    {
      action: ACTION_INSIGHT_CONFIGURATION_UPDATE,
    },
    'cog.svg',
    'John Doe updated the insight configuration settings',
  ],
  [
    { action: 'unknown action' },
    'error-warning.svg',
    "John Doe has performed unknown action 'unknown action'",
  ],
])(
  'AuditLogRecord message for %j',
  (auditLog, expectedIcon, expectedMessage) => {
    const { getByTestId, getByRole } = renderWithProviders(
      <AuditLogRecord {...auditLog} user="John Doe" />,
    )
    expect(getByRole('img')).toHaveTextContent(expectedIcon)
    expect(getByTestId('message')).toHaveTextContent(expectedMessage)
  },
)

test('meta information', () => {
  const { getByTestId } = renderWithProviders(
    <AuditLogRecord
      createdAt={new Date('2021-02-15T12:59:02.898590Z')}
      operatingSystem="Mac OS X"
      browser="Safari"
      client="NLX Management"
    />,
  )
  expect(getByTestId('meta')).toHaveTextContent(
    'Audit log created at • Mac OS X • Safari • NLX Management',
  )
})
