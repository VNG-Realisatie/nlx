// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled, { css } from 'styled-components'
import { Link } from 'react-router-dom'
import { IconChevronRight } from '../../icons'

export const StyledTable = styled.table`
  border-spacing: unset;
  width: 100%;
`

export const Th = styled.th`
  text-transform: uppercase;
  text-align: left;
  padding: ${(p) => p.theme.tokens.spacing04} ${(p) => p.theme.tokens.spacing03};
  border-bottom: 1px solid ${(p) => p.theme.colorBorderTable};
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`

const selectedTrCss = css`
  background-color: ${(p) => p.theme.colorBackgroundTableSelected};
`

export const StyledTr = styled.tr`
  ${(p) => (p.selected ? selectedTrCss : '')}
`

export const Td = styled.td`
  padding: ${(p) => p.theme.tokens.spacing04} ${(p) => p.theme.tokens.spacing03};
  border-bottom: 1px solid ${(p) => p.theme.colorBorderTable};
`

export const TrAsLink = styled(StyledTr)`
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
  color: ${(p) => p.theme.tokens.colorPaletteGray500};
`

export const StyledChevron = styled(IconChevronRight)`
  vertical-align: middle;
`
