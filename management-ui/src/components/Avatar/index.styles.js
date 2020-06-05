// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const Figure = styled.figure`
  padding: 0;
  margin: 0;
  height: ${(p) => p.theme.tokens.spacing09};

  .avatar-image {
    max-height: 100%;
    max-width: 100%;
    border-radius: 100%;
  }
`
