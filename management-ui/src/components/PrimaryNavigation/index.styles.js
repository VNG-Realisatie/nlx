import styled from 'styled-components'
import { Link, NavLink } from 'react-router-dom'
import NLXManagementLogo from '../NLXManagementLogo'

export const StyledNLXManagementLogo = styled(NLXManagementLogo)`
  width: 4rem;
  display: block;
  margin: 0 auto;
`

export const StyledNav = styled.nav`
  background: #313131;
  flex: 0 0 6.5rem;
  text-align: center;
  padding-top: ${(p) => p.theme.tokens.spacing06};
  height: 100%;
  color: ${(p) => p.theme.tokens.colorPaletteGray500};
  list-style-type: none;
  font-size: ${(p) => p.theme.tokens.spacing04};
  text-align: center;
`

export const StyledHomeLink = styled(Link)`
  display: block;
  padding: ${(p) => p.theme.tokens.spacing05} 0;
`

export const StyledLink = styled(NavLink)`
  padding: ${(p) => p.theme.tokens.spacing05} 0;
  text-decoration: none;
  color: ${(p) => p.theme.tokens.colorPaletteGray500};
  display: block;
  line-height: ${(p) => p.theme.tokens.spacing05};

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
  display: block;
  margin: 0 auto;
`
