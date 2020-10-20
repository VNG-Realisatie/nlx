// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import {
  outgoingAccessRequestPropTypes,
  ACCESS_REQUEST_STATES,
} from '../../../../../models/OutgoingAccessRequestModel'
import { InlineIcon, IconCheck } from '../../../../../icons'
import { WarnText } from './index.styles'

const { FAILED, CREATED, RECEIVED, APPROVED, REJECTED } = ACCESS_REQUEST_STATES

const getMessageForState = (state, t) => {
  switch (state) {
    case FAILED:
      return <WarnText>{t('Request could not be sent')}</WarnText>

    case CREATED:
      return <span>{t('Sending request')}</span>

    case RECEIVED:
      return <span>{t('Requested')}</span>

    case APPROVED:
      return <InlineIcon as={IconCheck} title={t('Approved')} />

    case REJECTED:
      return <span>{t('Rejected')}</span>

    default:
      throw new Error(`Can not determine message for unknown state '${state}'`)
  }
}

const AccessRequestMessage = ({ latestAccessRequest }) => {
  const { t } = useTranslation()
  const state = latestAccessRequest ? latestAccessRequest.state : null
  return getMessageForState(state, t)
}

AccessRequestMessage.propTypes = {
  latestAccessRequest: shape(pick(outgoingAccessRequestPropTypes, 'state')),
}

export default observer(AccessRequestMessage)
