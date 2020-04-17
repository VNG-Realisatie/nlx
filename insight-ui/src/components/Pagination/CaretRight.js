// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string } from 'prop-types'

const CaretRight = ({ color, ...props }) => (
  <svg viewBox="0 0 7 10" xmlns="http://www.w3.org/2000/svg" {...props}>
    <path
      d="M.59 8.825L4.407 5 .59 1.175 1.765 0l5 5-5 5z"
      fill={color}
      fillRule="nonzero"
    />
  </svg>
)

CaretRight.propTypes = {
  color: string,
  height: string,
}

CaretRight.defaultProps = {
  color: '#CAD0E0',
  height: '10px',
}

export default CaretRight
