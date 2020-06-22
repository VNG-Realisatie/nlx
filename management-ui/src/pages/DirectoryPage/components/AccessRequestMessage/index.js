// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, oneOf } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { WarnText } from './index.styles'

const statusMessage = {
  FAILED: (t) => <WarnText>{t('Request could not be sent')}</WarnText>,
  CREATED: (t) => <span>{t('Sending request')}</span>,
  SENT: (t) => <span>{t('Requested')}</span>,
}

const AccessRequestMessage = ({ latestAccessRequest, fallbackStatus }) => {
  const { t } = useTranslation()

  const status = latestAccessRequest
    ? latestAccessRequest.status
    : fallbackStatus || null

  return status ? statusMessage[status](t) : null
}

AccessRequestMessage.propTypes = {
  latestAccessRequest: shape({
    id: string,
    status: oneOf(Object.keys(statusMessage)),
    createdAt: string,
    updatedAt: string,
  }),
  fallbackStatus: oneOf(Object.keys(statusMessage)),
}

export default AccessRequestMessage
