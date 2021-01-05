// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect } from 'react'
import { bool, func, string, node } from 'prop-types'
import { createPortal } from 'react-dom'

import ModalFrame from './ModalFrame'

export const MASK_ANIMATION_SPEED_ENTER = 200
export const MASK_ANIMATION_SPEED_EXIT = 125
export const MASK_ANIMATION_DELAY_EXIT = 50
export const MODAL_ANIMATION_DELAY_ENTER = 50
export const MODAL_ANIMATION_SPEED_ENTER = 200
export const MODAL_ANIMATION_SPEED_EXIT = 125

const createRandomId = () => Math.random().toString(36).slice(2)
const MODAL_ROOT_ID = `modal-${createRandomId()}`

const Modal = (props) => {
  const [modalRoot] = useState(document.createElement('div'))

  useEffect(() => {
    modalRoot.setAttribute('id', MODAL_ROOT_ID)
    const reactRoot = document.querySelector('#root')
    reactRoot.parentNode.insertBefore(modalRoot, reactRoot.nextSibling)

    return () => {
      modalRoot.parentNode.removeChild(modalRoot)
    }
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  return createPortal(<ModalFrame {...props} />, modalRoot)
}

// For convenience, here's a list of all props used by ModalFrame and ModalContent
Modal.propTypes = {
  isVisible: bool.isRequired,
  handleClose: func.isRequired,
  title: string,
  showCloseButton: bool,
  allowUserToClose: bool,
  autoFocus: bool,
  maxWidth: string,
  offsetY: string,
  children: node,
}

export default Modal
