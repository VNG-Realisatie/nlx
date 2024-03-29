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

export const SERVICE_STATE_DEGRADED = 'STATE_DEGRADED'
const SERVICE_STATE_DOWN = 'STATE_DOWN'
const SERVICE_STATE_UNKNOWN = 'STATE_UNSPECIFIED'
export const SERVICE_STATE_UP = 'STATE_UP'

const GetStateIndicatorForState = (state, showText, t) => {
  switch (state) {
    case SERVICE_STATE_DEGRADED:
      return (
        <>
          <StyledIconStateDegraded title={t('Degraded')} />
          {showText && <StateText>{t('Degraded')}</StateText>}
        </>
      )

    case SERVICE_STATE_DOWN:
      return (
        <>
          <IconStateDown title={t('Down')} />
          {showText && <StateText>{t('Down')}</StateText>}
        </>
      )

    case SERVICE_STATE_UP:
      return (
        <>
          <IconStateUp title={t('Up')} />
          {showText && <StateText>{t('Up')}</StateText>}
        </>
      )

    case SERVICE_STATE_UNKNOWN:
    default:
      return (
        <>
          <IconStateUnknown title={t('Unknown')} />
          {showText && <StateText>{t('Unknown')}</StateText>}
        </>
      )
  }
}

const StateIndicator = ({ state, showText, ...props }) => {
  const { t } = useTranslation()

  return (
    <StyledWrapper {...props}>
      {GetStateIndicatorForState(state, showText, t)}
    </StyledWrapper>
  )
}

StateIndicator.propTypes = {
  state: oneOf([
    SERVICE_STATE_DOWN,
    SERVICE_STATE_DEGRADED,
    SERVICE_STATE_UNKNOWN,
    SERVICE_STATE_UP,
  ]),
  showText: bool,
}

StateIndicator.defaultProps = {
  showText: false,
}

export default StateIndicator
