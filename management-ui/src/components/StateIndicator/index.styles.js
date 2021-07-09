// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { IconStateDegraded } from '../../icons'

export const StyledWrapper = styled.span`
  display: flex;
  align-items: center;
`

export const StyledIconStateDegraded = styled(IconStateDegraded)`
  fill: ${(p) => p.theme.tokens.colorWarning};
`

export const StateText = styled.span`
  margin-left: ${(p) => p.theme.tokens.spacing02};
`
