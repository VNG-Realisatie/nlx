// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const Footer = styled.div`
  display: flex;
  justify-content: flex-end;
  margin-top: ${(p) => p.theme.tokens.spacing07};

  & > button {
    margin-left: ${(p) => p.theme.tokens.spacing05};
  }
`
