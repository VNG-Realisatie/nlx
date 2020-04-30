// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const StyledEmptyContentMessage = styled.p`
  color: ${(p) => p.theme.colorTextLabel};
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
  margin-bottom: 0;
  text-align: center;
  height: calc(100vh - 20rem);
  line-height: calc(100vh - 20rem);
`
