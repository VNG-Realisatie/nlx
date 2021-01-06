// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import {
  MASK_ANIMATION_SPEED_ENTER,
  MASK_ANIMATION_SPEED_EXIT,
  MASK_ANIMATION_DELAY_EXIT,
} from '../index'

export const Container = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
  overflow-y: auto;
`

export const HeightLimiter = styled.div`
  position: relative;
  display: flex;
  align-items: ${(p) => p.alignItems};
  padding: ${(p) => p.theme.tokens.spacing05};
  min-height: 100%;
`

export const Mask = styled.div`
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.6);
  cursor: ${(p) => (p.allowUserToClose ? 'pointer' : 'default')};
  opacity: 0;
  transform: translateX(-100%);

  &.mask-enter-active,
  &.mask-enter-done {
    transform: translateX(0);
    opacity: 0.75;
    transition: opacity ${() => MASK_ANIMATION_SPEED_ENTER}ms ease-in;
  }

  &.mask-exit {
    transform: translateX(0);
    opacity: 0.75;
  }

  &.mask-exit-active {
    opacity: 0;
    transition: opacity ${() => MASK_ANIMATION_SPEED_EXIT}ms ease-in
      ${() => MASK_ANIMATION_DELAY_EXIT}ms;
  }
`
