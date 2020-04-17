// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { node, string } from 'prop-types'
import { StyledCard } from './index.styles'

const ErrorMessage = ({ title, children, ...props }) => (
  <StyledCard {...props}>
    <h1>{title}</h1>
    {children}
  </StyledCard>
)

ErrorMessage.propTypes = {
  title: string.isRequired,
  children: node,
}

export default ErrorMessage
