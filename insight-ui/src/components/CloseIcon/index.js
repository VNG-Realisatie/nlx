// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { string } from 'prop-types'

const CloseIcon = ({ ...props }) =>
  <svg viewBox="0 0 14 14" xmlns="http://www.w3.org/2000/svg" {...props}>
    <path d="M14 1.41L12.59 0 7 5.59 1.41 0 0 1.41 5.59 7 0 12.59 1.41 14 7 8.41 12.59 14 14 12.59 8.41 7z" fill="#A3AABF" fillRule="nonzero"/>
  </svg>

CloseIcon.propTypes = {
  width: string,
  height: string
}

CloseIcon.defaultProps = {
  width: '14px',
  height: '14px'
}

export default CloseIcon

