// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

export const Container = styled.div`
  display: flex;
`

export const IconContainer = styled.div`
  flex: 0;
  
  &&:before {
    content: '';
    position: absolute;
    display: block
    background: ${(p) => p.theme.colorTextLabel};
  }
`

export const IconItem = styled.div`
  margin-right: ${(p) => p.theme.tokens.spacing05};
  color: ${(p) => p.theme.colorText};
`

export const Description = styled.p`
  flex: 1;
  color: ${(p) => p.theme.colorTextLabel};

  strong {
    color: ${(p) => p.theme.colorText};
  }
`
