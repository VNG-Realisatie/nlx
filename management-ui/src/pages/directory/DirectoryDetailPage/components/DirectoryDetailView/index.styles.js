// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Alert } from '@commonground/design-system'

export const StyledAlert = styled(Alert)`
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`

export const AccessMessage = styled.p`
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`
