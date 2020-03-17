// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import React from 'react'
import { StyledSvg } from './index.styles'

const Spinner = ({ ...props }) => (
  <StyledSvg viewBox="0 0 24 24" {...props}>
    <path fill="none" d="M0 0h24v24H0z" />
    <path d="M18.364 5.636L16.95 7.05A7 7 0 1019 12h2a9 9 0 11-2.636-6.364z" />
  </StyledSvg>
)

export default Spinner
