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
  ACTION_LOGOUT,
  ACTION_ORDER_CREATE,
  ACTION_ORGANIZATION_SETTINGS_UPDATE,
  ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
  ACTION_OUTGOING_ACCESS_REQUEST_FAIL,
  ACTION_SERVICE_CREATE,
  ACTION_SERVICE_DELETE,
  ACTION_SERVICE_UPDATE,
  ACTION_ORDER_OUTGOING_REVOKE,
} from '../../../../stores/models/AuditLogModel'
import AuditLogRecord from './index'

test.concurrent.each([
  [{ action: ACTION_LOGIN_SUCCESS }, 'shut-down.svg', 'John Doe has logged in'],
  [{ action: ACTION_LOGIN_FAIL }, 'shut-down.svg', 'Failed login attempt'],
  [{ action: ACTION_LOGOUT }, 'shut-down.svg', 'John Doe has logged out'],
  [
    {
      action: ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '00000000000000000001',
            name: 'Gemeente Stijns',
          },
        },
      ],
    },
    'check.svg',
    'John Doe has approved the access request from Gemeente Stijns (00000000000000000001) for Kadaster',
  ],
  [
    {
      action: ACTION_INCOMING_ACCESS_REQUEST_REJECT,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '00000000000000000001',
            name: 'Gemeente Stijns',
          },
        },
      ],
    },
    'close.svg',
    'John Doe has rejected the access request from Gemeente Stijns (00000000000000000001) for Kadaster',
  ],
  [
    {
      action: ACTION_ACCESS_GRANT_REVOKE,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '00000000000000000001',
            name: 'Gemeente Stijns',
          },
        },
      ],
    },
    'revoke.svg',
    'John Doe has revoked access for Kadaster from Gemeente Stijns',
  ],
  [
    {
      action: ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '00000000000000000001',
            name: 'Gemeente Stijns',
          },
        },
      ],
    },
    'key.svg',
    'John Doe has requested access to Kadaster from Gemeente Stijns (00000000000000000001)',
  ],
  [
    {
      action: ACTION_OUTGOING_ACCESS_REQUEST_FAIL,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '00000000000000000001',
            name: 'Gemeente Stijns',
          },
        },
      ],
    },
    'key.svg',
    'John Doe failed to request access to Kadaster from Gemeente Stijns (00000000000000000001)',
  ],
  [
    {
      action: ACTION_SERVICE_CREATE,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '',
            name: '',
          },
        },
      ],
    },
    'services.svg',
    'John Doe has created the service Kadaster',
  ],
  [
    {
      action: ACTION_SERVICE_UPDATE,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '',
            name: '',
          },
        },
      ],
    },
    'services.svg',
    'John Doe has updated the service Kadaster',
  ],
  [
    {
      action: ACTION_SERVICE_DELETE,
      services: [
        {
          service: 'Kadaster',
          organization: {
            serialNumber: '',
            name: '',
          },
        },
      ],
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
      action: ACTION_ORDER_CREATE,
      delegatee: 'Vergunningsoftware BV',
      services: [
        {
          service: 'fictieve-kentekens',
          organization: {
            serialNumber: '00000000000000000002',
            name: 'RvRD',
          },
        },
        {
          service: 'vakantieverhuur',
          organization: {
            serialNumber: '00000000000000000003',
            name: 'Gemeente Amsterdam',
          },
        },
      ],
    },
    'cog.svg',
    'John Doe gave Vergunningsoftware BV the order to consume the services fictieve-kentekens (RvRD (00000000000000000002)), vakantieverhuur (Gemeente Amsterdam (00000000000000000003))',
  ],
  [
    {
      action: ACTION_ORDER_OUTGOING_REVOKE,
      data: {
        delegatee: '00000000000000000001',
        reference: '0123456AB',
      },
    },
    'revoke.svg',
    'John Doe has revoked the outgoing order for 00000000000000000001 with reference 0123456AB',
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
