// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components/macro'
import { IconFlippingChevron } from '@commonground/design-system'
import Avatar from '../Avatar'

export const UserNavigationChevron = styled(IconFlippingChevron)`
  margin-left: ${(p) => p.theme.tokens.spacing04};
`

export const StyledUserNavigation = styled.div`
  position: relative;
  display: flex;
  align-items: center;

  @keyframes menuToggle {
    0% {
      transform: scaleY(0);
    }
    80% {
      transform: scaleY(1.1);
    }
    100% {
      transform: scaleY(1);
    }
  }
`

export const StyledUserMenu = styled.ul`
  position: absolute;
  margin: 0;
  padding: 0;
  top: calc(
    ${(p) => p.theme.tokens.spacing09} + ${(p) => p.theme.tokens.spacing02}
  );
  right: 0;
  z-index: 2;
  display: block;
  min-width: 12.5rem;
  list-style-type: none;
  background: ${(p) => p.theme.tokens.colorPaletteGray800};
  box-shadow: 0 5px 20px 0 rgba(0, 0, 0, 0.25);

  transform: scaleY(0);
  transform-origin: top center;
  transition: transform ${(p) => p.animationDuration}ms ease-in-out;

  &.user-menu-slide-enter-active,
  &.user-menu-slide-enter-done,
  &.user-menu-slide-exit {
    transform: scaleY(1);
  }

  &.user-menu-slide-exit-active {
    transform: scaleY(0);
  }

  button {
    width: 100%;
    padding: 0;
    border: none;
    font: inherit;
    color: inherit;
    background-color: transparent;
    cursor: pointer;
  }
`

export const StyledUserMenuItem = styled.li`
  &:not(:last-child) {
    border-bottom: 1px solid #e6e6e6;
  }

  a,
  button {
    display: block;
    text-align: left;
    text-decoration: none;
    color: ${(p) => p.theme.tokens.colorText};
    font-size: ${(p) => p.theme.tokens.fontSizeMedium};
    padding: ${(p) => `
      ${p.theme.tokens.spacing04} 
      ${p.theme.tokens.spacing06} 
      ${p.theme.tokens.spacing04} 
      ${p.theme.tokens.spacing05}
    `};
    border: 2px solid transparent;
  }

  a:hover,
  button:hover,
  li:hover {
    background-color: ${(p) => p.theme.colorBackgroundDropdownHover};
  }

  a:active,
  button:active {
    background-color: ${(p) => p.theme.colorBackgroundDropdownActive};
  }

  a:focus,
  button:focus {
    border-color: ${(p) => p.theme.colorBorderDropdownFocus};
  }
`

export const StyledToggleButton = styled.button`
  padding: 0;
  font: inherit;
  color: inherit;
  background-color: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
`

export const StyledAvatar = styled(Avatar)`
  padding: ${(p) => p.theme.tokens.spacing02};
  padding-right: 0;
  margin-right: ${(p) => p.theme.tokens.spacing04};
`

export const StyledUsername = styled.span`
  max-width: 15rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`
