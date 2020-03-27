// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

export const StyledAmount = styled.span`
  color: ${(p) => p.theme.colorTextLabel};
  font-weight: ${(p) => p.theme.tokens.fontWeightRegular};

  &:before {
    content: '·';
    padding: 0 ${(p) => p.theme.tokens.spacing03};
  }
`
