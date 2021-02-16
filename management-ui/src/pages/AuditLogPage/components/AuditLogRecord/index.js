// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { arrayOf, instanceOf, node, string } from 'prop-types'
import { Trans, useTranslation } from 'react-i18next'
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
  ACTION_OUTGOING_ACCESS_REQUEST_FAIL,
  ACTION_SERVICE_CREATE,
  ACTION_SERVICE_DELETE,
  ACTION_SERVICE_UPDATE,
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

const AuditLogRecord = ({
  action,
  user,
  createdAt,
  organization,
  service,
  operatingSystem,
  browser,
  client,
  ...props
}) => {
  const { t } = useTranslation()
  const dateTimeString = t('Audit log created at', { date: createdAt })

  const meta = [dateTimeString]

  if (operatingSystem) {
    meta.push(operatingSystem)
  }

  if (browser) {
    meta.push(browser)
  }

  if (client) {
    meta.push(client)
  }

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
        <Trans values={{ user, organization, service }}>
          <strong>{{ user }}</strong> has approved the access request from{' '}
          <strong>{{ organization }}</strong> for <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_INCOMING_ACCESS_REQUEST_REJECT ? (
        <Trans values={{ user, organization, service }}>
          <strong>{{ user }}</strong> has rejected the access request from{' '}
          <strong>{{ organization }}</strong> for <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_ACCESS_GRANT_REVOKE ? (
        <Trans values={{ user, organization, service }}>
          <strong>{{ user }}</strong> has revoked access for{' '}
          <strong>{{ service }}</strong> from{' '}
          <strong>{{ organization }}</strong>
        </Trans>
      ) : action === ACTION_OUTGOING_ACCESS_REQUEST_CREATE ? (
        <Trans values={{ user, organization, service }}>
          <strong>{{ user }}</strong> has requested access to{' '}
          <strong>{{ service }}</strong> from{' '}
          <strong>{{ organization }}</strong>
        </Trans>
      ) : action === ACTION_OUTGOING_ACCESS_REQUEST_FAIL ? (
        <Trans values={{ user, organization, service }}>
          <strong>{{ user }}</strong> failed to request access to{' '}
          <strong>{{ service }}</strong> from{' '}
          <strong>{{ organization }}</strong>
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
      ) : action === ACTION_INSIGHT_CONFIGURATION_UPDATE ? (
        <Trans values={{ user, action }}>
          <strong>{{ user }}</strong> updated the insight configuration settings
        </Trans>
      ) : (
        <Trans values={{ user, action }}>
          <strong>{{ user }}</strong> has performed unknown action{' '}
          <strong>'{{ action }}'</strong>
        </Trans>
      )}
    </Template>
  )
}

AuditLogRecord.propTypes = {
  action: string,
  user: string,
  createdAt: instanceOf(Date),
  organization: string,
  service: string,
  operatingSystem: string,
  browser: string,
  client: string,
}

export default AuditLogRecord
