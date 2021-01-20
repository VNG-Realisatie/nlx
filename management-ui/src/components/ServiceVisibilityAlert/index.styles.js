// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { Alert } from '@commonground/design-system'

export const StyledAlert = styled(Alert)``

export const WarningMessage = styled.p`
  margin: ${(p) => p.theme.tokens.spacing04} 0;
  color: ${(p) => p.theme.tokens.colorWarning};
`
