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
} from '../../../../stores/models/AuditLogModel'
import {
  IconWarningCircle,
  IconShutdown,
  IconCheck,
  IconRevoke,
  IconClose,
  IconKey,
  IconServices,
} from '../../../../icons'
import { Container, IconContainer, Description, IconItem } from './index.styles'

const actionToIcon = (action) => {
  switch (action) {
    case ACTION_LOGIN_SUCCESS:
    case ACTION_LOGOUT_SUCCESS:
      return IconShutdown

    case ACTION_INCOMING_ACCESS_REQUEST_ACCEPT:
      return IconCheck

    case ACTION_INCOMING_ACCESS_REQUEST_REJECT:
      return IconClose

    case ACTION_ACCESS_GRANT_REVOKE:
      return IconRevoke

    case ACTION_OUTGOING_ACCESS_REQUEST_CREATE:
      return IconKey

    case ACTION_SERVICE_CREATE:
      return IconServices

    default:
      return IconWarningCircle
  }
}

const Template = ({ action, dateTime, children, ...props }) => (
  <Container {...props}>
    <IconContainer>
      <IconItem as={actionToIcon(action)} />
    </IconContainer>
    <Description>
      <span data-testid="message">{children}</span>
      <br />
      <small>{dateTime}</small>
    </Description>
  </Container>
)

Template.propTypes = {
  action: string,
  dateTime: string,
  children: node,
}

const AuditLogRecord = ({ action, user, createdAt, organization, service }) => {
  const { t } = useTranslation()
  const dateTimeString = t('Audit log created at', { date: createdAt })

  return (
    <Template action={action} dateTime={dateTimeString}>
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
        <Trans values={{ user, action, organization, service }}>
          <strong>{{ user }}</strong> has approved the access request from{' '}
          <strong>{{ organization }}</strong> for <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_INCOMING_ACCESS_REQUEST_REJECT ? (
        <Trans values={{ user, action, organization, service }}>
          <strong>{{ user }}</strong> has rejected the access request from{' '}
          <strong>{{ organization }}</strong> for <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_ACCESS_GRANT_REVOKE ? (
        <Trans values={{ user, action, organization, service }}>
          <strong>{{ user }}</strong> has revoked acces for{' '}
          <strong>{{ service }}</strong> from <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_OUTGOING_ACCESS_REQUEST_CREATE ? (
        <Trans values={{ user, action, organization, service }}>
          <strong>{{ user }}</strong> requested access to{' '}
          <strong>{{ service }}</strong> from{' '}
          <strong>{{ organization }}</strong>
        </Trans>
      ) : action === ACTION_SERVICE_CREATE ? (
        <Trans values={{ user, action, service }}>
          <strong>{{ user }}</strong> has created the service{' '}
          <strong>{{ service }}</strong>
        </Trans>
      ) : action === ACTION_SERVICE_UPDATE ? (
        <Trans values={{ user, action, service }}>
          <strong>{{ user }}</strong> has updated the service{' '}
          <strong>{{ service }}</strong>
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
}

export default AuditLogRecord
