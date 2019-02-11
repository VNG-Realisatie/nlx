import React from 'react'
import { oneOf } from 'prop-types'
import StyledDocsIcon from './DocsIcon.styles'

const DocsIcon = ({ color }) =>
  <StyledDocsIcon color={ color } viewBox="0 0 24 24">
    <g fillRule="nonzero">
      <path d="M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H17c1.1 0 2-.9 2-2V7l-5-5zM6 20V4h7l4 4v12H6z"/>
      <path d="M13 2h1l5 5v1h-6zM8 12h7v1H8v-1zm0 2h7v1H8v-1zm0 2h7v1H8v-1zm0-6h5v1H8v-1z"/>
    </g>
  </StyledDocsIcon>

DocsIcon.propTypes = {
  color: oneOf(['blue', 'grey']).isRequired
}

export default DocsIcon
