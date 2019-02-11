import React from 'react'
import { oneOf } from 'prop-types'
import StyledLinkIcon from './LinkIcon.styles'

const colors = {
  blue: '#517FFF',
  grey: '#CAD0E0'
}

const LinkIcon = ({ color }) =>
  <StyledLinkIcon viewBox="0 0 24 24">
    <g fill="none" fillRule="evenodd">
      <circle cx="12" cy="12" r="12"/>
      <path
        d="M6.845 6.841a2.792 2.792 0 0 1 3.946 0l2.545 2.546 1.21-1.21L12 5.633a4.502 4.502 0 0 0-6.364 0 4.502 4.502 0 0 0 0 6.364l2.546 2.546 1.209-1.21-2.546-2.545a2.792 2.792 0 0 1 0-3.946zm1.973 3.246l5.091 5.09 1.273-1.272-5.091-5.091-1.273 1.273zm9.546 1.909L15.818 9.45l-1.209 1.21 2.546 2.545a2.792 2.792 0 0 1 0 3.946 2.792 2.792 0 0 1-3.946 0l-2.545-2.546-1.21 1.21L12 18.36a4.502 4.502 0 0 0 6.364 0 4.502 4.502 0 0 0 0-6.364z"
        fill={colors[color]} fillRule="nonzero"/>
    </g>
  </StyledLinkIcon>

LinkIcon.propTypes = {
  color: oneOf(['blue', 'grey']).isRequired
}

export default LinkIcon
