// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { IconStateDegraded } from '../../icons'

export const StyledWrapper = styled.span`
  display: flex;
  align-items: center;
  line-height: 0;
  vertical-align: bottom;
`

export const StyledIconStateDegraded = styled(IconStateDegraded)`
  .mainPath {
    fill: ${(p) => p.theme.tokens.colorAlertWarning};
  }
`

export const StateText = styled.span`
  margin-left: ${(p) => p.theme.tokens.spacing02};
`
