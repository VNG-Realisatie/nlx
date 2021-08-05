// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'

import { StyledEmptyContentMessage } from './index.styles'

const EmptyContentMessage = ({ children, ...props }) => (
  <StyledEmptyContentMessage {...props}>{children}</StyledEmptyContentMessage>
)
EmptyContentMessage.propTypes = {
  children: node,
}

export default EmptyContentMessage
