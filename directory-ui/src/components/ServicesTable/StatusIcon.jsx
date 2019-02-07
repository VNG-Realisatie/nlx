import React from 'react'
import { oneOf } from 'prop-types'
import { StyledSvg, StyledCircle } from './StatusIcon.styles'

const StatusIcon = ({ status }) =>
  <StyledSvg viewBox="0 0 20 20">
    <StyledCircle status={status} strokeWidth="2" cx="10" cy="10" r="4" fill="none" fillRule="evenodd"/>
  </StyledSvg>

StatusIcon.propTypes = {
  status: oneOf(['online', 'offline']).isRequired
}

export default StatusIcon
