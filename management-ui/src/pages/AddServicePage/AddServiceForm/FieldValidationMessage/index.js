// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'
import styled from 'styled-components'

const getSvg = (fillColor) =>
  `<svg viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg"><path d="M16 16H0V0h16v16zm-9-5v2h2v-2H7zm0-8v6h2V3H7z" fill="${fillColor}"/></svg>`

export const StyledFieldValidationMessage = styled.p`
  background: url(data:image/svg+xml;base64,${(p) =>
    btoa(getSvg(p.theme.tokens.colorError))}) no-repeat left center;
  background-size: ${(p) => p.theme.tokens.spacing05};
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
