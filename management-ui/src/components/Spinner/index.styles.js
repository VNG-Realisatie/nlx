// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled, { keyframes } from 'styled-components'

const rotate = keyframes`
  100% {
    transform:rotate(360deg);
  }
`

export const StyledSvg = styled.svg`
  width: ${(p) => p.theme.tokens.spacing06};
  height: ${(p) => p.theme.tokens.spacing06};
  margin-right: ${(p) => p.theme.tokens.spacing04};
  animation: ${rotate} 0.75s linear infinite;
  fill: ${(p) => p.theme.colorTextLink};
`
