// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { oneOf, bool } from 'prop-types'

import { IconStateUp, IconStateDown, IconStateUnknown } from '../../../icons'

import {
  StyledWrapper,
  StyledIconStateDegraded,
  StateText,
} from './index.styles'

export const SERVICE_STATE_DEGRADED = 'degraded'
const SERVICE_STATE_DOWN = 'down'
const SERVICE_STATE_UNKNOWN = 'unknown'
export const SERVICE_STATE_UP = 'up'

const GetStateIndicatorForState = (state, showText) => {
  switch (state) {
    case SERVICE_STATE_DEGRADED:
      return (
        <>
          <StyledIconStateDegraded title="Degraded" />
          {showText && <StateText>Degraded</StateText>}
        </>
      )

    case SERVICE_STATE_DOWN:
      return (
        <>
          <IconStateDown title="Down" />
          {showText && <StateText>Down</StateText>}
        </>
      )

    case SERVICE_STATE_UP:
      return (
        <>
          <IconStateUp title="Up" />
          {showText && <StateText>Up</StateText>}
        </>
      )

    case SERVICE_STATE_UNKNOWN:
    default:
      return (
        <>
          <IconStateUnknown title="Unknown" />
          {showText && <StateText>Unknown</StateText>}
        </>
      )
  }
}

const StateIndicator = ({ state, showText, ...props }) => {
  return (
    <StyledWrapper {...props}>
      {GetStateIndicatorForState(state, showText)}
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
