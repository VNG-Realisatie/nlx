// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { Icon } from '@commonground/design-system'
import { IconMoneyEuroCircleLine } from '../../../icons'
import { StyledLabel, StyledDetailHeading } from './index.styles'

const CollapsibleHeader = ({ label }) => {
  return (
    <StyledDetailHeading>
      <Icon as={IconMoneyEuroCircleLine} />
      Kosten
      <StyledLabel>{label}</StyledLabel>
    </StyledDetailHeading>
  )
}

CollapsibleHeader.propTypes = {
  label: string,
}

export default CollapsibleHeader
