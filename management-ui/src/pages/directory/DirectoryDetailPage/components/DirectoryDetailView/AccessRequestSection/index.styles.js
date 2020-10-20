// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const AccessSection = styled.section`
  display: flex;
  align-items: center;
  justify-content: space-between;
`

export const IconItem = styled.div`
  margin-right: ${(p) => p.theme.tokens.spacing03};
  color: ${(p) => p.theme.colorTextLabel};
`

export const StateDetail = styled.div`
  flex: 1 1 auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
`

export const ErrorText = styled.span`
  color: ${(p) => p.theme.tokens.colorError};

  svg {
    float: left;
    margin-right: ${(p) => p.theme.tokens.spacing03};
    fill: ${(p) => p.theme.tokens.colorError};
  }
`
