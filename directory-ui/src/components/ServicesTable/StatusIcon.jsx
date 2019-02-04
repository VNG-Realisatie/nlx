import React from 'react'
import { oneOf } from 'prop-types'
import StyledSvg from './StatusIcon.styles'

const StatusIcon = ({ status }) =>
  <StyledSvg viewBox="0 0 10 10" status={status}>
    <circle
      cx="5"
      cy="14"
      r="5"
      transform="translate(0 -9)"
      fill="currentColor"
    />
  </StyledSvg>

StatusIcon.propTypes = {
  status: oneOf(['online', 'offline']).isRequired
}

export default StatusIcon
