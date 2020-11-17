// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Link, NavLink } from 'react-router-dom'
import NLXManagementLogo from '../NLXManagementLogo'

export const StyledNLXManagementLogo = styled(NLXManagementLogo)`
  width: 4rem;
`

export const Nav = styled.nav`
  flex: 0 0 6.5rem;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 100%;
  padding: ${(p) => p.theme.tokens.spacing06} 0;
  font-size: ${(p) => p.theme.tokens.spacing04};
  text-align: center;
  list-style-type: none;
  color: ${(p) => p.theme.tokens.colorPaletteGray500};
  background: #313131;
`

export const StyledHomeLink = styled(Link)`
  display: block;
  padding: ${(p) => p.theme.tokens.spacing05} 0;
`

export const StyledLink = styled(NavLink)`
  display: block;
  padding: ${(p) => p.theme.tokens.spacing05} 0;
  line-height: ${(p) => p.theme.tokens.spacing05};
  text-decoration: none;
  color: ${(p) => p.theme.tokens.colorPaletteGray500};

  svg path {
    fill: ${(p) => p.theme.tokens.colorPaletteGray500};
  }

  &:hover {
    color: ${(p) => p.theme.tokens.colorPaletteGray500};
    background: ${(p) => p.theme.tokens.colorPaletteGray800};
  }

  &.active {
    position: relative;
    color: ${(p) => p.theme.tokens.colorBrand1};

    &:before {
      content: '';
      position: absolute;
      top: ${(p) => p.theme.tokens.spacing03};
      bottom: ${(p) => p.theme.tokens.spacing03};
      left: 0;
      width: 4px;
      background: ${(p) => p.theme.tokens.colorBrand1};
    }

    svg path {
      fill: ${(p) => p.theme.tokens.colorBrand1};
    }
  }
`

export const StyledIcon = styled.svg`
  width: ${(p) => p.theme.tokens.spacing07};
  height: ${(p) => p.theme.tokens.spacing07};
  display: block;
  margin: 0 auto;
`
