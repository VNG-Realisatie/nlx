// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

export const StyledMain = styled.main`
  display: flex;
  align-items: flex-start;
  height: 100%;
`

export const StyledContent = styled.div`
  flex: 1;
  padding: ${(p) => p.theme.tokens.spacing07} ${(p) => p.theme.tokens.spacing09};
  overflow: auto;
  height: 100%;
`
