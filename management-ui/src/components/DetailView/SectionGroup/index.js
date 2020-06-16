// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { node } from 'prop-types'
import React from 'react'
import { StyledSectionGroup } from './index.styles'

const SectionGroup = ({ children, ...props }) => (
  <StyledSectionGroup {...props}>
    {Array.isArray(children) ? children.filter((x) => x) : children}
  </StyledSectionGroup>
)

SectionGroup.propTypes = {
  children: node.isRequired,
}

export default SectionGroup
