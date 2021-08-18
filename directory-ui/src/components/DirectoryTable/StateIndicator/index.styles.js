// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { mediaQueries } from '@commonground/design-system'
import { IconStateDegraded } from '../../../icons'

export const StyledWrapper = styled.span`
  ${mediaQueries.smDown`
    margin-right: ${(p) => p.theme.tokens.spacing02};
    margin-left: -${(p) => p.theme.tokens.spacing02};
  `}

  ${mediaQueries.mdUp`
    display: flex;
    align-items: center;
  `}
`

export const StyledIconStateDegraded = styled(IconStateDegraded)`
  fill: ${(p) => p.theme.tokens.colorWarning};
`

export const StateText = styled.span`
  margin-left: ${(p) => p.theme.tokens.spacing02};
`
