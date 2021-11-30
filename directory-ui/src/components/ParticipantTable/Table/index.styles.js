// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Link } from 'react-router-dom'
import { mediaQueries } from '@commonground/design-system'

export const StyledTable = styled.table`
  border-collapse: collapse;
  width: 100%;
`

export const Thead = styled.thead`
  ${mediaQueries.smDown`
    display: none;
  `}
`

export const Th = styled.th`
  text-transform: uppercase;
  text-align: left;
  padding: ${(p) => p.theme.tokens.spacing04} ${(p) => p.theme.tokens.spacing03};
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`

export const Td = styled.td`
  display: none;

  ${mediaQueries.mdUp`
    display: table-cell;
    padding: ${(p) => p.theme.tokens.spacing04}
      ${(p) => p.theme.tokens.spacing03};
  `}
`

export const TrAsLinkTd = styled.td`
  padding: ${(p) => p.theme.tokens.spacing04} ${(p) => p.theme.tokens.spacing03};
  width: 1px;
`

export const MobileTd = styled.td`
  display: flex;
  padding: ${(p) => p.theme.tokens.spacing04} ${(p) => p.theme.tokens.spacing03}
    ${(p) => p.theme.tokens.spacing04} 0;

  ${mediaQueries.mdUp`
    display: none;
  `}
`

export const MobileTdContent = styled.div`
  span {
    margin-right: ${(p) => p.theme.tokens.spacing02};
  }
  p {
    margin: 0;
    display: block;
  }
  p:first-of-type {
    font-weight: bold;
  }
  p:nth-child(3) {
    margin-top: ${(p) => p.theme.tokens.spacing02};
    color: ${(p) => p.theme.tokens.colorPaletteGray600};
  }
`

export const StyledTr = styled.tr`
  background-color: ${(p) => p.theme.colorBackground};
  border-bottom: 1px solid ${(p) => p.theme.colorBorderTable};

  :first-child {
    border-top: 1px solid ${(p) => p.theme.colorBorderTable};
  }
`

export const TrAsLink = styled(StyledTr)`
  cursor: pointer;
  background-color: ${(p) =>
    p.selected && p.theme.colorBackgroundTableSelected};

  :hover {
    background-color: ${(p) =>
      !p.selected && p.theme.colorBackgroundTableHover};
  }
`

export const TrAsLinkLink = styled(Link)`
  display: block;
  border: none;
  line-height: 100%;
  color: ${(p) => p.theme.tokens.colorPaletteGray500};
`
