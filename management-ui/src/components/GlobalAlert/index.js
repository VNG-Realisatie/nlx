// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'
import { IconWarning } from '../../icons'
import { Alert, StyledIcon } from './index.styles'

const GlobalAlert = ({ children }) => (
  <Alert role="alert">
    <StyledIcon as={IconWarning} inline />
    {children}
  </Alert>
)

GlobalAlert.propTypes = {
  children: node.isRequired,
}

GlobalAlert.defaultProps = {}

export default GlobalAlert
