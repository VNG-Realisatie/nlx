// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const List = styled.dl`
  display: table;
  width: 100%;
  margin: 0;

  dd {
    ${({ alignValuesRight }) => (alignValuesRight ? 'text-align: right;' : '')}
  }
`

export const Item = styled.div`
  display: table-row;

  dt,
  dd {
    display: table-cell;
    padding-bottom: 0.5rem;
  }

  dt {
    font-size: ${(p) => p.theme.tokens.fontSizeSmall};
    color: ${(p) => p.theme.tokens.colorPaletteGray500};
  }

  dd {
    ${({ alignValue }) => (alignValue ? `text-align: ${alignValue}` : '')}
  }
`
