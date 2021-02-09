// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string, node, instanceOf } from 'prop-types'
import { useTranslation, Trans } from 'react-i18next'
import {
  AUDIT_LOG_ACTION_LOGIN,
  AUDIT_LOG_ACTION_LOGOUT,
} from '../../../../stores/models/AuditLogModel'
import { IconWarningCircle, IconShutdown } from '../../../../icons'
import { Container, IconContainer, Description, IconItem } from './index.styles'

const actionToIcon = (action) => {
  switch (action) {
    case AUDIT_LOG_ACTION_LOGIN:
    case AUDIT_LOG_ACTION_LOGOUT:
      return IconShutdown

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

const AuditLogRecord = ({ action, user, createdAt }) => {
  const { t } = useTranslation()
  const dateTimeString = t('Audit log created at', { date: createdAt })

  return (
    <Template action={action} dateTime={dateTimeString}>
      {action === AUDIT_LOG_ACTION_LOGIN ? (
        <Trans values={{ user }}>
          <strong>{{ user }}</strong> has logged in
        </Trans>
      ) : action === AUDIT_LOG_ACTION_LOGOUT ? (
        <Trans values={{ user }}>
          <strong>{{ user }}</strong> has logged out
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
}

export default AuditLogRecord
