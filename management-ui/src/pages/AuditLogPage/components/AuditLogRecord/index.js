// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string, node, instanceOf } from 'prop-types'
import { useTranslation, Trans } from 'react-i18next'
import {
  ACTION_LOGIN_FAIL,
  ACTION_LOGIN_SUCCESS,
  ACTION_LOGOUT_SUCCESS,
  ACTION_INCOMING_ACCESS_REQUEST_ACCEPT,
  ACTION_INCOMING_ACCESS_REQUEST_REJECT,
  ACTION_ACCESS_GRANT_REVOKE,
  ACTION_OUTGOING_ACCESS_REQUEST_CREATE,
  ACTION_SERVICE_CREATE,
  ACTION_SERVICE_UPDATE,
  ACTION_SERVICE_DELETE,
  ACTION_ORGANIZATION_SETTINGS_UPDATE,
  ACTION_INSIGHT_CONFIGURATION_UPDATE,
} from '../../../../stores/models/AuditLogModel'
import iconForActionType from './icon-for-action-type'
import { Container, IconContainer, Description, IconItem } from './index.styles'

const Template = ({ action, dateTime, userAgent, children, ...props }) => (
  <Container {...props}>
    <IconContainer data-testid="icon">
      <IconItem as={iconForActionType(action)} role="img" />
    </IconContainer>
    <Description>
      <span data-testid="message">{children}</span>
      <br />
      <small>
        {dateTime}
        {'  '}●{'  '}
        {userAgent}
      </small>
    </Description>
  </Container>
)

Template.propTypes = {
  action: string,
  dateTime: string,
  children: node,
  userAgent: string,
}

const AuditLogRecord = ({
  action,
  user,
  createdAt,
  organization,
  service,
  userAgent,
  ...props
}) => {
  const { t } = useTranslation()
  const dateTimeString = t('Audit log created at', { date: createdAt })

  return (
    <Template
      action={action}
      dateTime={dateTimeString}
      userAgent={userAgent}
      {...props}
    >
      {action === ACTION_LOGIN_SUCCESS ? (
        <Trans values={{ user }}>
          <strong>{{ user }}</strong> has logged in
        </Trans>
      ) : action === ACTION_LOGOUT_SUCCESS ? (
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
  userAgent: string,
}

export default AuditLogRecord
