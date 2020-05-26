// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

export const StyledPageTitle = styled.h1`
  margin-bottom: ${(p) => p.theme.tokens.spacing01};
`

export const StyledHeaderItems = styled.div`
  margin-left: auto; /* Aligns it right when no title present */
  display: flex;
  align-items: center;

  & > * + *:before {
    content: '';
    height: ${(p) => p.theme.tokens.spacing07};
    width: 1px;
    background-color: ${(p) => p.theme.tokens.colorPaletteGray600};
    margin-right: ${(p) => p.theme.tokens.spacing07};
    margin-left: ${(p) => p.theme.tokens.spacing07};
  }
`

export const StyledDescription = styled.p`
  margin-bottom: ${(p) => p.theme.tokens.spacing07};
`

export const StyledHeader = styled.div`
  display: flex;
  justify-content: space-between;
`
