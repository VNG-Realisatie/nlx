// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, oneOf, bool } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconWarningCircleFill } from '../../../../icons'
import { FailedDetail, ErrorText, WarnText } from './index.styles'

const statusMessage = {
  FAILED: (t, inDetailView) =>
    inDetailView ? (
      <FailedDetail>
        <span>{t('Access request')}</span>
        <ErrorText>
          <IconWarningCircleFill title={t('Error')} />
          {t('Request could not be sent')}
        </ErrorText>
      </FailedDetail>
    ) : (
      <WarnText>{t('Request could not be sent')}</WarnText>
    ),
  CREATED: (t) => <span>{t('Sending request')}</span>,
  RECEIVED: (t) => <span>{t('Requested')}</span>,
}

const AccessRequestMessage = ({ latestAccessRequest, inDetailView }) => {
  const { t } = useTranslation()
  const status = latestAccessRequest ? latestAccessRequest.status : null

  return status ? statusMessage[status](t, inDetailView) : null
}

AccessRequestMessage.propTypes = {
  latestAccessRequest: shape({
    id: string,
    status: oneOf(Object.keys(statusMessage)),
    createdAt: string,
    updatedAt: string,
  }),
  inDetailView: bool,
}

export default AccessRequestMessage
