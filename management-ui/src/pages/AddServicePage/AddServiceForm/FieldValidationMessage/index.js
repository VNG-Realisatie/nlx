// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'
import styled from 'styled-components'
import iconExclamationMark from './icon-exclamation-mark.svg'

export const StyledFieldValidationMessage = styled.p`
  background: url(${iconExclamationMark}) no-repeat left center;
  font-family: 'Source Sans Pro', sans-serif;
  font-weight: ${(p) => p.theme.tokens.fontWeightBold};
  color: ${(p) => p.theme.tokens.colorError};
  margin: ${(p) => p.theme.tokens.spacing03} 0
    ${(p) => p.theme.tokens.spacing06} 0;
  padding-left: ${(p) => p.theme.tokens.spacing06};
`

const FieldValidationMessage = ({ children, ...props }) => (
  <StyledFieldValidationMessage {...props}>
    {children}
  </StyledFieldValidationMessage>
)

FieldValidationMessage.propTypes = {
  children: node,
}

export default FieldValidationMessage
