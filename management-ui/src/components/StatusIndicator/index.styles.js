// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { ReactComponent as IconStatusDegraded } from './status-degraded.svg'

export const StyledWrapper = styled.span`
  line-height: 0;
  vertical-align: bottom;
`

export const StyledIconStatusDegraded = styled(IconStatusDegraded)`
  .mainPath {
    fill: ${(p) => p.theme.tokens.colorAlertWarning};
  }
`
