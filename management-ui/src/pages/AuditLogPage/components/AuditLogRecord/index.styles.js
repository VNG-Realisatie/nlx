// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

export const IconContainer = styled.div`
  flex: 0;
  position: relative;

  &&:before {
    content: '';
    position: absolute;
    display: block;
    background: ${(p) => p.theme.colorTextLabel};
    width: 1px;
    top: 32px;
    left: 9.5px;
    bottom: 8px;
  }
`

export const Container = styled.div`
  display: flex;

  &&:last-child {
    ${IconContainer}:before {
      content: none;
    }
  }
`

export const IconItem = styled.div`
  margin-right: ${(p) => p.theme.tokens.spacing05};
  color: ${(p) => p.theme.colorTextLabel};
`

export const Description = styled.p`
  flex: 1;
  color: ${(p) => p.theme.colorPaletteGray400};
  padding-bottom: ${(p) => p.theme.tokens.spacing04};

  strong {
    color: ${(p) => p.theme.colorText};
    font-weight: 600;
  }
`
