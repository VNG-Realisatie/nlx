// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { oneOf, bool } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconStateUp, IconStateDown, IconStateUnknown } from '../../icons'

import {
  StyledWrapper,
  StyledIconStateDegraded,
  StateText,
} from './index.styles'

export const DIRECTORY_SERVICE_STATES = ['degraded', 'down', 'unknown', 'up']

// Generic component that will handle different kinds of state codes (not only directory service)
const StateIndicator = ({ state, showText }) => {
  const { t } = useTranslation()

  if (!DIRECTORY_SERVICE_STATES.includes(state)) {
    console.warn(`Invalid state '${state}'`)
    return null
  }

  // Make this smarter when refactoring for more states:
  return (
    <StyledWrapper>
      {
        {
          degraded: (
            <>
              <StyledIconStateDegraded title={t('Degraded')} />
              {showText && <StateText>{t('Degraded')}</StateText>}
            </>
          ),
          down: (
            <>
              <IconStateDown title={t('Down')} />
              {showText && <StateText>{t('Down')}</StateText>}
            </>
          ),
          up: (
            <>
              <IconStateUp title={t('Up')} />
              {showText && <StateText>{t('Up')}</StateText>}
            </>
          ),
          unknown: (
            <>
              <IconStateUnknown title={t('Unknown')} />
              {showText && <StateText>{t('Unknown')}</StateText>}
            </>
          ),
        }[state]
      }
    </StyledWrapper>
  )
}

StateIndicator.propTypes = {
  state: oneOf(DIRECTORY_SERVICE_STATES),
  showText: bool,
}

StateIndicator.defaultProps = {
  showText: false,
}

export default StateIndicator
