// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number, bool } from 'prop-types'
import { StyledAmount } from './index.styles'

const Amount = ({ value, ...props }) => (
  <StyledAmount {...props}>{value}</StyledAmount>
)

Amount.propTypes = {
  value: number.isRequired,
  isAccented: bool,
}

Amount.defaultProps = {
  isAccented: false,
}

export default Amount
