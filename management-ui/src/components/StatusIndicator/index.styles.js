// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { IconStatusDegraded } from '../../icons'

export const StyledWrapper = styled.span`
  line-height: 0;
  vertical-align: bottom;
`

export const StyledIconStatusDegraded = styled(IconStatusDegraded)`
  .mainPath {
    fill: ${(p) => p.theme.tokens.colorAlertWarning};
  }
`
