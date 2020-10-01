// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, bool } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { outgoingAccessRequestPropTypes } from '../../../../../models/OutgoingAccessRequestModel'
import { IconWarningCircleFill } from '../../../../../icons'
import { FailedDetail, ErrorText, WarnText } from './index.styles'

const STATE_FAILED = 'FAILED'
const STATE_CREATED = 'CREATED'
const STATE_RECEIVED = 'RECEIVED'

const getMessageForState = (state, t, inDetailView) => {
  switch (state) {
    case STATE_FAILED:
      return inDetailView ? (
        <FailedDetail>
          <span>{t('Access request')}</span>
          <ErrorText>
            <IconWarningCircleFill title={t('Error')} />
            {t('Request could not be sent')}
          </ErrorText>
        </FailedDetail>
      ) : (
        <WarnText>{t('Request could not be sent')}</WarnText>
      )

    case STATE_CREATED:
      return <span>{t('Sending request')}</span>

    case STATE_RECEIVED:
      return <span>{t('Requested')}</span>

    default:
      console.warn(`can not determine message for unknown state '${state}'`)
      return null
  }
}

const AccessRequestMessage = ({ latestAccessRequest, inDetailView }) => {
  const { t } = useTranslation()
  const state = latestAccessRequest ? latestAccessRequest.state : null
  return getMessageForState(state, t, inDetailView)
}

AccessRequestMessage.propTypes = {
  latestAccessRequest: shape(pick(outgoingAccessRequestPropTypes, 'state')),
  inDetailView: bool,
}

export default observer(AccessRequestMessage)
