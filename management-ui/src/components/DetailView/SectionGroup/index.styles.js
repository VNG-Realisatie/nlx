// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const StyledSectionGroup = styled.div`
  border-bottom: 1px solid ${(p) => p.theme.tokens.colorPaletteGray700};

  & > * {
    border-top: 1px solid ${(p) => p.theme.tokens.colorPaletteGray700};
    padding: ${(p) => p.theme.tokens.spacing05} 0;
  }
`
