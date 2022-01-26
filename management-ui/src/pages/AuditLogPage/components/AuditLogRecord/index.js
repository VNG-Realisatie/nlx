// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { arrayOf, instanceOf, node, string } from 'prop-types'
import { Trans, useTranslation } from 'react-i18next'
import AuditLogModel, {
  ACTION_ACCESS_GRANT_REVOKE,
  ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
  ACTION_INCOMING_ACCESS_REQUEST_REJECT,
  ACTION_LOGIN_FAIL,
  ACTION_LOGIN_SUCCESS,
  ACTION_LOGOUT,
  ACTION_ORDER_CREATE,
  ACTION_ORDER_OUTGOING_REVOKE,
  ACTION_ORDER_OUTGOING_UPDATE,
  ACTION_ORGANIZATION_SETTINGS_UPDATE,
  ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
  ACTION_OUTGOING_ACCESS_REQUEST_FAIL,
  ACTION_SERVICE_CREATE,
  ACTION_SERVICE_DELETE,
  ACTION_SERVICE_UPDATE,
  ACTION_INWAY_DELETE,
} from '../../../../stores/models/AuditLogModel'
import iconForActionType from './icon-for-action-type'
import {
  Container,
  Description,
  IconContainer,
  IconItem,
  Meta,
} from './index.styles'

const Template = ({ action, meta, children, ...props }) => (
  <Container {...props}>
    <IconContainer data-testid="icon">
      <IconItem as={iconForActionType(action)} role="img" />
    </IconContainer>
    <Description>
      <span data-testid="message">{children}</span>
      <br />
      <Meta data-testid="meta">{meta.join('   •   ')}</Meta>
    </Description>
  </Container>
)

Template.propTypes = {
  action: string,
  dateTime: string,
  children: node,
  meta: arrayOf(string),
}

const AuditLogRecord = ({ model, ...props }) => {
  const { t } = useTranslation()
  const dateTimeString = t('Audit log created at', { date: model.createdAt })

  const meta = [dateTimeString]

  if (model.operatingSystem) {
    meta.push(model.operatingSystem)
  }

  if (model.browser) {
    meta.push(model.browser)
  }

  if (model.client) {
    meta.push(model.client)
  }

  let organization = {}
  let service = ''
  let servicesList = ''

  if (model.services && model.services.length) {
    organization = model.services[0].organization
    service = model.services[0].service
    servicesList = model.services
      .map(
        (service) =>
          `${service.service} (${service.organization.name} (${service.organization.serialNumber}))`,
      )
      .join(', ')
  }

  const dataDelegatee = model.data.delegatee
  const dataReference = model.data.reference

  const dataInwayName = model.data.inwayName

  const organizationSerialNumber = organization.serialNumber
  const organizationName = organization.name

  const user = model.user
  const action = model.action

  return (
    <Template action={action} dateTime={dateTimeString} meta={meta} {...props}>
      {action === ACTION_LOGIN_SUCCESS ? (
        <Trans values={{ user }}>
          <strong>{{ user }}</strong> has logged in
        </Trans>
      ) : action === ACTION_LOGOUT ? (
        <Trans values={{ user }}>
          <strong>{{ user }}</strong> has logged out
        </Trans>
      ) : action === ACTION_LOGIN_FAIL ? (
        <Trans>Failed login attempt</Trans>
      ) : action === ACTION_INCOMING_ACCESS_REQUEST_ACCEPT ? (
        <Trans
          values={{ user, organizationSerialNumber, organizationName, service }}
        >
          <strong>{{ user }}</strong> has approved the access request from{' '}
          <strong>
            {{ organizationName }} ({{ organizationSerialNumber }})
          </strong>{' '}
          for <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_INCOMING_ACCESS_REQUEST_REJECT ? (
        <Trans values={{ user, organization, service }}>
          <strong>{{ user }}</strong> has rejected the access request from{' '}
          <strong>{{ organizationSerialNumber }}</strong> for{' '}
          <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_ACCESS_GRANT_REVOKE ? (
        <Trans
          values={{ user, organizationSerialNumber, organizationName, service }}
        >
          <strong>{{ user }}</strong> has revoked access for{' '}
          <strong>{{ service }}</strong> from{' '}
          <strong>
            {{ organizationName }} ({{ organizationSerialNumber }})
          </strong>
        </Trans>
      ) : action === ACTION_OUTGOING_ACCESS_REQUEST_CREATE ? (
        <Trans
          values={{ user, organizationSerialNumber, organizationName, service }}
        >
          <strong>{{ user }}</strong> has requested access to{' '}
          <strong>{{ service }}</strong> from{' '}
          <strong>{{ organizationSerialNumber }}</strong>
        </Trans>
      ) : action === ACTION_OUTGOING_ACCESS_REQUEST_FAIL ? (
        <Trans
          values={{ user, organizationSerialNumber, organizationName, service }}
        >
          <strong>{{ user }}</strong> failed to request access to{' '}
          <strong>{{ service }}</strong> from{' '}
          <strong>
            {{ organizationName }} ({{ organizationSerialNumber }})
          </strong>
        </Trans>
      ) : action === ACTION_SERVICE_CREATE ? (
        <Trans values={{ user, service }}>
          <strong>{{ user }}</strong> has created the service{' '}
          <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_SERVICE_UPDATE ? (
        <Trans values={{ user, service }}>
          <strong>{{ user }}</strong> has updated the service{' '}
          <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_SERVICE_UPDATE ? (
        <Trans values={{ user, service }}>
          <strong>{{ user }}</strong> has updated the service{' '}
          <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_SERVICE_DELETE ? (
        <Trans values={{ user, service }}>
          <strong>{{ user }}</strong> has removed the service{' '}
          <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_ORGANIZATION_SETTINGS_UPDATE ? (
        <Trans values={{ user, action }}>
          <strong>{{ user }}</strong> updated the organization settings
        </Trans>
      ) : action === ACTION_ORDER_CREATE ? (
        <Trans values={{ user, servicesList, dataDelegatee, action }}>
          <strong>{{ user }}</strong> gave {{ dataDelegatee }} the order to
          consume the services {{ servicesList }}
        </Trans>
      ) : action === ACTION_ORDER_OUTGOING_REVOKE ? (
        <Trans values={{ user, dataDelegatee, dataReference }}>
          <strong>{{ user }}</strong> has revoked the outgoing order for{' '}
          {{ dataDelegatee }} with reference {{ dataReference }}
        </Trans>
      ) : action === ACTION_INWAY_DELETE ? (
        <Trans values={{ user, dataInwayName }}>
          <strong>{{ user }}</strong> has removed the inway{' '}
          <strong>{{ dataInwayName }}</strong>
        </Trans>
      ) : action === ACTION_ORDER_OUTGOING_UPDATE ? (
        <Trans values={{ user, servicesList, dataDelegatee, action }}>
          <strong>{{ user }}</strong> updated the order for {{ dataDelegatee }}{' '}
          the services {{ servicesList }}
        </Trans>
      ) : (
        <Trans values={{ user, action }}>
          <strong>{{ user }}</strong> has performed unknown action{' '}
          <strong>&apos;{{ action }}&apos;</strong>
        </Trans>
      )}
    </Template>
  )
}

AuditLogRecord.propTypes = {
  model: instanceOf(AuditLogModel).isRequired,
}

export default AuditLogRecord
