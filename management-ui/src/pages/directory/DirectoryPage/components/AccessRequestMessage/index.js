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

const stateMessage = {
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
  const state = latestAccessRequest ? latestAccessRequest.state : null

  return state ? stateMessage[state](t, inDetailView) : null
}

AccessRequestMessage.propTypes = {
  latestAccessRequest: shape(pick(outgoingAccessRequestPropTypes, 'state')),
  inDetailView: bool,
}

export default observer(AccessRequestMessage)
