// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { IconMoneyEuroCircleLine } from '../../../icons'
import { StyledLabel, StyledDetailHeading } from './index.styles'

const CollapsibleHeader = ({ label }) => {
  return (
    <StyledDetailHeading>
      <IconMoneyEuroCircleLine />
      Costs
      <StyledLabel>{label}</StyledLabel>
    </StyledDetailHeading>
  )
}

CollapsibleHeader.propTypes = {
  label: string,
}

export default CollapsibleHeader
