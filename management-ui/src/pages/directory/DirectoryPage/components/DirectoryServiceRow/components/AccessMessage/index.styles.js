// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const Message = styled.span`
  flex: 1;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
`

export const WarnMessage = styled(Message)`
  color: ${(p) => p.theme.tokens.colorWarning};
`
