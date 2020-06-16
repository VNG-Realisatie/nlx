// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { oneOf } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconStatusUp, IconStatusDown, IconStatusUnknown } from '../../icons'

import { StyledWrapper, StyledIconStatusDegraded } from './index.styles'

export const DIRECTORY_SERVICE_STATUS = ['degraded', 'down', 'unknown', 'up']

// Generic component that will handle different kinds of status codes (not only directory service)
const StatusIndicator = ({ status }) => {
  const { t } = useTranslation()

  if (!DIRECTORY_SERVICE_STATUS.includes(status)) {
    console.warn(`Invalid status '${status}'`)
    return null
  }

  return (
    <StyledWrapper>
      {
        {
          degraded: <StyledIconStatusDegraded title={t('Degraded')} />,
          down: <IconStatusDown title={t('Down')} />,
          up: <IconStatusUp title={t('Up')} />,
          unknown: <IconStatusUnknown title={t('Unknown')} />,
        }[status]
      }
    </StyledWrapper>
  )
}

StatusIndicator.propTypes = {
  status: oneOf(DIRECTORY_SERVICE_STATUS),
}

export default StatusIndicator
