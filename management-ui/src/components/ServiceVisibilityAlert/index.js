// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { StyledAlert, WarningMessage } from './index.styles'

export const showServiceVisibilityAlert = ({ inways, internal }) =>
  !internal && inways.length === 0

export const ServiceVisibilityMessage = (props) => {
  const { t } = useTranslation()
  return (
    <WarningMessage role="alert" {...props}>
      {t('Service not yet accessible')}
    </WarningMessage>
  )
}

export const ServiceVisibilityAlert = (props) => {
  const { t } = useTranslation()
  return (
    <StyledAlert
      variant="warning"
      title={t('Service not yet accessible')}
      {...props}
    >
      {t(
        'There are no inways connected yet. Until then other organizations cannot access this service.',
      )}
    </StyledAlert>
  )
}
