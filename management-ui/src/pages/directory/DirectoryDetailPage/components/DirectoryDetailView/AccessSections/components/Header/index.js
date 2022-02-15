// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { IconKey } from '../../../../../../../../icons'
import { StyledLabel, StyledDetailHeading } from './index.styles'

const Header = ({ title, label }) => (
  <StyledDetailHeading>
    <IconKey />
    {title}
    <StyledLabel>{label}</StyledLabel>
  </StyledDetailHeading>
)

Header.propTypes = {
  title: string,
  label: string,
}

export default Header
