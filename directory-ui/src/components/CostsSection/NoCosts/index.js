// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { Icon } from '@commonground/design-system'
import { IconMoneyEuroCircleLine } from '../../../icons'
import { StyledContainer, StyledLabel } from './index.styles'

const NoCosts = () => {
  return (
    <StyledContainer>
      <Icon as={IconMoneyEuroCircleLine} />
      Kosten
      <StyledLabel>Geen</StyledLabel>
    </StyledContainer>
  )
}

export default NoCosts
