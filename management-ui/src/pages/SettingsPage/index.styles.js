// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

export const StyledMain = styled.div`
  display: flex;
  align-items: flex-start;
  height: calc(100% - 6rem);
`

export const StyledSidebar = styled.div`
  flex: 0 0 16rem;
  height: 100%;
  padding: 0 ${(p) => p.theme.tokens.spacing09} 0 0;
  border-right: 1px solid ${(p) => p.theme.tokens.colorPaletteGray700};
`

export const StyledContent = styled.div`
  flex: 1;
  padding: 0 0 0 ${(p) => p.theme.tokens.spacing08};
`
