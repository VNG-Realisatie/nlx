// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect } from 'react'
import { bool, func, string, node, shape, oneOf } from 'prop-types'
import { createPortal } from 'react-dom'
import cssUnitOrEmpty from './propTypeCssUnitOrEmpty'
import ModalFrame from './ModalFrame'

export const MASK_ANIMATION_SPEED_ENTER = 200
export const MASK_ANIMATION_SPEED_EXIT = 125
export const MASK_ANIMATION_DELAY_EXIT = 50
export const MODAL_ANIMATION_DELAY_ENTER = 50
export const MODAL_ANIMATION_SPEED_ENTER = 200
export const MODAL_ANIMATION_SPEED_EXIT = 125

const createRandomId = () => Math.random().toString(36).slice(8)

export const verticalAlignToCssValues = ({ from, offset }) => {
  // Results in array, eg: ['100px', '100', 'px']
  const offsetMatch = offset && offset.match(/(-?\d+)([px|%|rem]+)/)
  const offsetBottom = offsetMatch
    ? {
        amount: offsetMatch[1],
        unit: offsetMatch[2],
      }
    : {}

  return {
    alignItems:
      from === 'top' ? 'flex-start' : from === 'bottom' ? 'flex-end' : 'center',
    transform: offsetMatch
      ? `translateY(${
          from === 'bottom'
            ? offsetBottom.amount * -1 + offsetBottom.unit
            : offset
        })`
      : '',
  }
}

const Modal = ({ verticalAlign = {}, ...props }) => {
  const [modalRoot] = useState(document.createElement('div'))

  useEffect(() => {
    modalRoot.setAttribute('id', `modal-${createRandomId()}`)
    const reactRoot = document.querySelector('#root')
    reactRoot.parentNode.insertBefore(modalRoot, reactRoot.nextSibling)

    return () => {
      modalRoot.parentNode.removeChild(modalRoot)
    }
  }, [modalRoot])

  const verticalAlignCss = verticalAlignToCssValues(verticalAlign)

  return createPortal(
    <ModalFrame {...props} verticalAlignCss={verticalAlignCss} />,
    modalRoot,
  )
}

// For convenience, here's a list of all props used by ModalFrame and ModalContent
Modal.propTypes = {
  isVisible: bool.isRequired,
  handleClose: func.isRequired,
  title: string,
  showCloseButton: bool,
  allowUserToClose: bool,
  autoFocus: bool,
  width: string,
  verticalAlign: shape({
    from: oneOf(['center', 'top', 'bottom']),
    offset: cssUnitOrEmpty,
  }),
  children: node,
}

export default Modal
