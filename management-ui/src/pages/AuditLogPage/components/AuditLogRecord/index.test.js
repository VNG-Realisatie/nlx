// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import AuditLogModel, {
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
  ACTION_INWAY_DELETE,
  ACTION_ACCEPT_TERMS_OF_SERVICE,
  ACTION_OUTWAY_DELETE,
} from '../../../../stores/models/AuditLogModel'
import AuditLogRecord from './index'

const createModel = (data) => {
  const result = new AuditLogModel({ auditLogData: data })
  result.user = 'John Doe'
  return result
}

test.concurrent.each([
  [
    createModel({ action: ACTION_LOGIN_SUCCESS }),
    'shut-down.svg',
    'John Doe has logged in',
  ],
  [
    createModel({ action: ACTION_LOGIN_FAIL }),
    'shut-down.svg',
    'Failed login attempt',
  ],
  [
    createModel({ action: ACTION_LOGOUT }),
    'shut-down.svg',
    'John Doe has logged out',
  ],
  [
    createModel({
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
    }),
    'check.svg',
    'John Doe has approved the access request from Gemeente Stijns for Kadaster',
  ],
  [
    createModel({
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
    }),
    'close.svg',
    'John Doe has rejected the access request from Gemeente Stijns for Kadaster',
  ],
  [
    createModel({
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
    }),
    'revoke.svg',
    'John Doe has revoked access for Kadaster from Gemeente Stijns',
  ],
  [
    createModel({
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
    }),
    'key.svg',
    'John Doe has requested access to Kadaster from Gemeente Stijns',
  ],
  [
    createModel({
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
    }),
    'key.svg',
    'John Doe failed to request access to Kadaster from Gemeente Stijns',
  ],
  [
    createModel({
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
    }),
    'services.svg',
    'John Doe has created the service Kadaster',
  ],
  [
    createModel({
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
    }),
    'services.svg',
    'John Doe has updated the service Kadaster',
  ],
  [
    createModel({
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
    }),
    'services.svg',
    'John Doe has removed the service Kadaster',
  ],
  [
    createModel({
      action: ACTION_ORGANIZATION_SETTINGS_UPDATE,
    }),
    'cog.svg',
    'John Doe updated the organization settings',
  ],
  [
    createModel({
      action: ACTION_ORDER_CREATE,
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
      data: {
        delegatee: {
          serialNumber: '00000000000000000001',
          name: 'Vergunningsoftware BV',
        },
      },
    }),
    'cog.svg',
    'John Doe gave Vergunningsoftware BV the order to consume the services fictieve-kentekens (RvRD), vakantieverhuur (Gemeente Amsterdam)',
  ],
  [
    createModel({
      action: ACTION_ORDER_OUTGOING_REVOKE,
      data: {
        delegatee: {
          serialNumber: '00000000000000000001',
          name: 'Vergunningsoftware BV',
        },
        reference: '0123456AB',
      },
    }),
    'revoke.svg',
    'John Doe has revoked the outgoing order for Vergunningsoftware BV with reference 0123456AB',
  ],
  [
    createModel({
      action: ACTION_INWAY_DELETE,
      data: {
        inwayName: 'my-inway',
      },
    }),
    'cog.svg',
    'John Doe has removed the Inway my-inway',
  ],
  [
    createModel({
      action: ACTION_ACCEPT_TERMS_OF_SERVICE,
    }),
    'cog.svg',
    'John Doe has accepted the Terms of Service',
  ],
  [
    createModel({
      action: ACTION_OUTWAY_DELETE,
      data: {
        outwayName: 'my-outway',
      },
    }),
    'cog.svg',
    'John Doe has removed the Outway my-outway',
  ],
  [
    createModel({ action: 'unknown action' }),
    'error-warning.svg',
    "John Doe has performed unknown action 'unknown action'",
  ],
])('AuditLogRecord message for %j', (model, expectedIcon, expectedMessage) => {
  renderWithProviders(<AuditLogRecord model={model} />)
  expect(screen.getByRole('img')).toHaveTextContent(expectedIcon)
  expect(screen.getByTestId('message')).toHaveTextContent(expectedMessage)
})

test('meta information', () => {
  const model = new AuditLogModel({
    auditLogData: {
      createdAt: new Date('2021-02-15T12:59:02.898590Z'),
      operatingSystem: 'Mac OS X',
      browser: 'Safari',
      client: 'NLX Management',
    },
  })
  renderWithProviders(<AuditLogRecord model={model} />)
  expect(screen.getByTestId('meta')).toHaveTextContent(
    'Audit log created at • Mac OS X • Safari • NLX Management',
  )
})
