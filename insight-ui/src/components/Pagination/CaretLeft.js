import React from 'react'
import { string } from 'prop-types'

const CaretLeft = ({ color, ...props }) => (
  <svg viewBox="0 0 7 10" xmlns="http://www.w3.org/2000/svg" {...props}>
    <path
      d="M7 9.707L2.673 5.5 7 1.293 5.668 0 0 5.5 5.668 11z"
      fill={color}
      fillRule="nonzero"
    />
  </svg>
)

CaretLeft.propTypes = {
  color: string,
  height: string,
}

CaretLeft.defaultProps = {
  color: '#CAD0E0',
  height: '10px',
}

export default CaretLeft
