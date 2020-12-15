// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Icon } from '@commonground/design-system'

export const Alert = styled.div`
  padding: ${(p) => p.theme.tokens.spacing05} ${(p) => p.theme.tokens.spacing05}
    ${(p) => p.theme.tokens.spacing05} ${(p) => p.theme.tokens.spacing09};
  background-color: ${(p) => p.theme.colorAlertWarningBackground};
`

export const StyledIcon = styled(Icon)`
  margin-right: ${(p) => p.theme.tokens.spacing05};
`
