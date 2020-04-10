// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import UserNavigation from '../../UserNavigation'

export const StyledPageTitle = styled.h1`
  margin-bottom: ${(p) => p.theme.tokens.spacing01};
`

export const StyledUserNavigation = styled(UserNavigation)`
  margin-left: auto; /* Aligns it right when no title present */
`

export const StyledDescription = styled.p`
  margin-bottom: ${(p) => p.theme.tokens.spacing07};
`

export const StyledHeader = styled.div`
  display: flex;
  justify-content: space-between;
`
