// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'

import {
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
} from '../../../directoryServiceAccessState'
import { InlineIcon, IconCheck } from '../../../../../icons'
import Switch from '../../../../../components/Switch'
import { WarnText } from './index.styles'

const AccessMessage = ({ displayState }) => {
  const { t } = useTranslation()

  return (
    <Switch test={displayState}>
      <Switch.Case value={SHOW_REQUEST_CREATED}>
        {() => <span>{t('Sending request')}</span>}
      </Switch.Case>
      <Switch.Case value={SHOW_REQUEST_FAILED}>
        {() => <WarnText>{t('Request could not be sent')}</WarnText>}
      </Switch.Case>
      <Switch.Case value={SHOW_REQUEST_RECEIVED}>
        {() => <span>{t('Requested')}</span>}
      </Switch.Case>
      <Switch.Case value={SHOW_REQUEST_REJECTED}>
        {() => <span>{t('Rejected')}</span>}
      </Switch.Case>
      <Switch.Case value={SHOW_HAS_ACCESS}>
        {() => <InlineIcon as={IconCheck} title={t('Approved')} />}
      </Switch.Case>
      {/* Purposely displaying nothing in any other case */}
      <Switch.Default>{() => null}</Switch.Default>
    </Switch>
  )
}

AccessMessage.propTypes = {
  displayState: number,
}

export default observer(AccessMessage)
