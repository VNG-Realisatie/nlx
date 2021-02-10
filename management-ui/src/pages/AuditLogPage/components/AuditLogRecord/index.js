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
} from '../../../../stores/models/AuditLogModel'
import { IconWarningCircle, IconShutdown, IconCheck } from '../../../../icons'
import { Container, IconContainer, Description, IconItem } from './index.styles'

const actionToIcon = (action) => {
  switch (action) {
    case ACTION_LOGIN_SUCCESS:
    case ACTION_LOGOUT_SUCCESS:
      return IconShutdown

    case ACTION_INCOMING_ACCESS_REQUEST_ACCEPT:
      return IconCheck

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
