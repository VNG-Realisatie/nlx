// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { oneOf, bool } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconStatusUp, IconStatusDown, IconStatusUnknown } from '../../icons'

import {
  StyledWrapper,
  StyledIconStatusDegraded,
  StatusText,
} from './index.styles'

export const DIRECTORY_SERVICE_STATUS = ['degraded', 'down', 'unknown', 'up']

// Generic component that will handle different kinds of status codes (not only directory service)
const StatusIndicator = ({ status, showText }) => {
  const { t } = useTranslation()

  if (!DIRECTORY_SERVICE_STATUS.includes(status)) {
    console.warn(`Invalid status '${status}'`)
    return null
  }

  // Make this smarter when refactoring for more statuses:
  return (
    <StyledWrapper>
      {
        {
          degraded: (
            <>
              <StyledIconStatusDegraded title={t('Degraded')} />
              {showText && <StatusText>{t('Degraded')}</StatusText>}
            </>
          ),
          down: (
            <>
              <IconStatusDown title={t('Down')} />
              {showText && <StatusText>{t('Down')}</StatusText>}
            </>
          ),
          up: (
            <>
              <IconStatusUp title={t('Up')} />
              {showText && <StatusText>{t('Up')}</StatusText>}
            </>
          ),
          unknown: (
            <>
              <IconStatusUnknown title={t('Unknown')} />
              {showText && <StatusText>{t('Unknown')}</StatusText>}
            </>
          ),
        }[status]
      }
    </StyledWrapper>
  )
}

StatusIndicator.propTypes = {
  status: oneOf(DIRECTORY_SERVICE_STATUS),
  showText: bool,
}

StatusIndicator.defaultProps = {
  showText: false,
}

export default StatusIndicator
