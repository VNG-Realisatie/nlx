// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { NavLink } from 'react-router-dom'

export const StyledNav = styled.nav`
  flex: 0 0 6.5rem;
  height: 100%;
  color: ${(p) => p.theme.tokens.colorText};
  list-style-type: none;
`

export const StyledLink = styled(NavLink)`
  padding: ${(p) => p.theme.tokens.spacing03} ${(p) => p.theme.tokens.spacing05};
  text-decoration: none;
  color: ${(p) => p.theme.colorText};
  display: block;

  &:hover {
    color: ${(p) => p.theme.colorText};
    background: rgba(255, 255, 255, 0.1);
  }

  &.active {
    background: rgba(255, 255, 255, 0.2);
    position: relative;
    font-weight: ${(p) => p.theme.tokens.fontWeightBold};

    &:before {
      content: '';
      position: absolute;
      top: -1px;
      bottom: -1px;
      left: -1px;
      width: 4px;
      background: ${(p) => p.theme.tokens.colorBrand1};
    }

    &:hover {
      background: rgba(255, 255, 255, 0.3);
    }
  }
`
