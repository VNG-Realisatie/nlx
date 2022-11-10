// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { useTranslation } from 'react-i18next'
import OutgoingAccessRequestSyncErrorModel from '../../../../../../stores/models/OutgoingAccessRequestSyncErrorModel'
import { StyledAlert } from './index.styles'

const SyncErrorSection = ({ syncError }) => {
  const { t } = useTranslation()

  return (
    <StyledAlert variant="warning" title={t('Synchronization error')}>
      {t(syncError.message)}
    </StyledAlert>
  )
}

SyncErrorSection.propTypes = {
  syncError: instanceOf(OutgoingAccessRequestSyncErrorModel),
}

export default SyncErrorSection
