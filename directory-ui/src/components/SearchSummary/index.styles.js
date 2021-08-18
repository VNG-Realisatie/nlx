// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'

export const Text = styled.p`
  margin: ${(p) => p.theme.tokens.spacing06} 0
    ${(p) => p.theme.tokens.spacing05};
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
  color: ${(p) => p.theme.tokens.colorPaletteGray600};
  line-height: ${(p) => p.theme.tokens.lineHeightText};
`
