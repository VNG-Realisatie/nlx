// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const FailedDetail = styled.div`
  display: flex;
  flex-direction: column;
`

export const ErrorText = styled.span`
  color: ${(p) => p.theme.tokens.colorError};

  svg {
    float: left;
    margin-right: ${(p) => p.theme.tokens.spacing03};
    fill: ${(p) => p.theme.tokens.colorError};
  }
`

export const WarnText = styled.span`
  color: ${(p) => p.theme.tokens.colorWarning};
`
