// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Link } from 'react-router-dom'
import { ReactComponent as Chevron } from './chevron-right.svg'

export const StyledTable = styled.table`
  border-spacing: unset;
  width: 100%;
`

export const Th = styled.th`
  text-transform: uppercase;
  text-align: left;
  padding: ${(p) => p.theme.tokens.spacing04} 0;
  border-bottom: 1px solid ${(p) => p.theme.colorBorderTable};
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`

export const Td = styled.td`
  padding: ${(p) => p.theme.tokens.spacing04} 0;
  border-bottom: 1px solid ${(p) => p.theme.colorBorderTable};
`

export const TrAsLink = styled.tr`
  cursor: pointer;

  &:hover {
    background-color: ${(p) => p.theme.colorBackgroundTableHover};
  }

  &:focus {
    outline: 2px solid ${(p) => p.theme.colorBorderTableFocus};
  }
`

export const TrAsLinkTd = styled(Td)`
  width: 1px;
`

export const TrAsLinkLink = styled(Link)`
  display: block;
  border: none;
  line-height: 100%;
`

export const StyledChevron = styled(Chevron)`
  vertical-align: middle;
  fill: ${(p) => p.theme.colorTextLabel};
`
