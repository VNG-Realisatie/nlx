// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled, { css } from 'styled-components'

import {
  MODAL_ANIMATION_DELAY_ENTER,
  MODAL_ANIMATION_SPEED_ENTER,
  MODAL_ANIMATION_SPEED_EXIT,
} from '../index'

export const ModalPosition = styled.div`
  width: max-content;
  max-width: 100%;
  margin: 0 auto;
  z-index: 2;

  ${(p) =>
    p.maxWidth.match(/(px|rem|%)$/) &&
    css`
      width: 100%;
      max-width: ${p.maxWidth};
    `}

  ${(p) =>
    p.offsetY.match(/(px|rem|%)$/) &&
    css`
      transform: translateY(${p.offsetY});
    `}
`

export const Window = styled.div`
  display: flex;
  flex-direction: column-reverse;
  overflow-x: auto;
  background-color: ${(p) => p.theme.tokens.colorBackgroundAlt};
  transform: scale(0);

  .modal-content-enter-active & {
    transform: scale(1);
    transition: transform ${() => MODAL_ANIMATION_SPEED_ENTER}ms ease-in-out
      ${() => MODAL_ANIMATION_DELAY_ENTER}ms;
  }

  .modal-content-enter-done &,
  .modal-content-exit & {
    transform: scale(1);
  }

  .modal-content-exit-active & {
    transform: scale(0);
    transition: transform ${() => MODAL_ANIMATION_SPEED_EXIT}ms ease-in;
  }
`

export const Content = styled.div`
  padding: ${(p) => {
    const { spacing07 } = p.theme.tokens
    return `0 ${spacing07} ${spacing07}`
  }};
`
