// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Link, NavLink } from 'react-router-dom'
import { Icon } from '@commonground/design-system'

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

export const StyledIcon = styled(Icon)`
  display: block;
  margin: 0 auto 2px;
`

export const OrdersLink = styled(StyledLink)`
  &::after {
    position: absolute;
    top: 18px;
    left: 40px;
    width: 20px;
    height: 20px;
    content: '😃';
    font-size: 15px;
    opacity: 0;
  }
  @keyframes smile {
    0% {
      opacity: 0;
    }
    80% {
      opacity: 0;
    }
    81% {
      opacity: 1;
    }
    99% {
      opacity: 1;
    }
  }

  &.active::after {
    animation: 6s infinite smile;
  }
`

export const InwaysIcon = styled(StyledIcon)`
  @keyframes turn {
    0% {
      transform: scaleX(1);
    }
    100% {
      transform: scaleX(-1);
    }
  }

  .active & {
    animation: 0.25s ease-in-out forwards turn;
  }
`

export const ServicesIcon = styled(StyledIcon)`
  & > path {
    opacity: 1;
  }

  & #elec {
    opacity: 0;
  }

  @keyframes elec-off {
    12% {
      opacity: 1;
    }
    13% {
      opacity: 0;
    }
    15% {
      opacity: 0;
    }
    16% {
      opacity: 1;
    }
    19% {
      opacity: 1;
    }
    20% {
      opacity: 0;
    }
    22% {
      opacity: 0;
    }
    23% {
      opacity: 1;
    }
  }
  @keyframes elec-on {
    12% {
      opacity: 0;
    }
    13% {
      opacity: 1;
    }
    15% {
      opacity: 1;
    }
    16% {
      opacity: 0;
    }
    19% {
      opacity: 0;
    }
    20% {
      opacity: 1;
    }
    22% {
      opacity: 1;
    }
    23% {
      opacity: 0;
    }
  }

  .active & > path {
    animation: 5s infinite elec-off;
  }
  .active & > #elec {
    animation: 5s infinite elec-on;
  }
`

export const DirectoryIcon = styled(StyledIcon)`
  @keyframes appear {
    20% {
      transform: scaleY(0.3) translateY(30px);
    }
    60% {
      transform: scaleY(1.2) translateY(-15px);
    }
    100% {
      transform: scaleY(1) translateY(0);
    }
  }

  .active & {
    animation: 1.2s ease-out 3s appear;
  }
`

export const BarChartIcon = styled.div`
  position: relative;
  width: 24px;
  height: 24px;
  margin: 0 auto 2px;

  @keyframes juke {
    0% {
      transform: scaleY(2.8);
    }
    20% {
      transform: scaleY(2);
    }
    40% {
      transform: scaleY(0.8);
    }
    60% {
      transform: scaleY(1.2);
    }
    80% {
      transform: scaleY(3.4);
    }
    100% {
      transform: scaleY(1.5);
    }
  }

  & > div {
    position: absolute;
    bottom: 0;
    background-color: ${(p) => p.theme.tokens.colorPaletteGray500};
    transform-origin: bottom;
  }

  & :nth-child(1) {
    left: 0;
    width: 3px;
    height: 13px;
  }
  & :nth-child(2) {
    left: 11px;
    width: 3px;
    height: 28px;
  }
  & :nth-child(3) {
    left: 22px;
    width: 3px;
    height: 20px;
  }

  .active & > div {
    background-color: ${(p) => p.theme.tokens.colorBrand1};
  }
  .active:hover & > :nth-child(1) {
    height: 10px;
    animation: 1.4s linear infinite juke;
  }
  .active:hover & > :nth-child(2) {
    height: 10px;
    animation: 1.5s linear -2.5s infinite juke;
  }
  .active:hover & > :nth-child(3) {
    height: 10px;
    animation: 1s linear -1.3s infinite juke;
  }
`

export const TimeIcon = styled(StyledIcon)`
  @keyframes clock {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }

  .active & {
    animation: 60s steps(60) infinite clock;
  }
`

export const SettingsIcon = styled(StyledIcon)`
  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    80% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }

  .active & {
    animation: 5s ease-in-out infinite spin;
  }
`
