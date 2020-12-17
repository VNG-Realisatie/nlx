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
import { IconCheck } from '../../../../../icons'
import Switch from '../../../../../components/Switch'
import { Message, WarnMessage } from './index.styles'

const AccessMessage = ({ displayState }) => {
  const { t } = useTranslation()

  return (
    <Switch test={displayState}>
      <Switch.Case value={SHOW_REQUEST_CREATED}>
        {() => <Message>{t('Sending request')}</Message>}
      </Switch.Case>
      <Switch.Case value={SHOW_REQUEST_FAILED}>
        {() => <WarnMessage>{t('Request could not be sent')}</WarnMessage>}
      </Switch.Case>
      <Switch.Case value={SHOW_REQUEST_RECEIVED}>
        {() => <Message>{t('Requested')}</Message>}
      </Switch.Case>
      <Switch.Case value={SHOW_REQUEST_REJECTED}>
        {() => <Message>{t('Rejected')}</Message>}
      </Switch.Case>
      <Switch.Case value={SHOW_HAS_ACCESS}>
        {() => (
          <Message>
            <IconCheck title={t('Approved')} inline />
          </Message>
        )}
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
