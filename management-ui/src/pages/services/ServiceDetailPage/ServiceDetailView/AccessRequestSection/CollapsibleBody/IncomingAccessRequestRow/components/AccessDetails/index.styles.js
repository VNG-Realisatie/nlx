// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Icon } from '@commonground/design-system'

export const SectionHeader = styled.p`
  margin-bottom: ${(p) => p.theme.tokens.spacing02};
  color: ${(p) => p.theme.tokens.colorPaletteGray400};
  padding-left: ${(p) => (p.withoutIcon ? p.theme.tokens.spacing07 : '0')};
`

export const SectionContent = styled.p`
  margin-left: ${(p) => p.theme.tokens.spacing07};
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`

export const StyledIcon = styled(Icon)`
  fill: ${(p) => p.theme.tokens.colorPaletteGray500};
`
