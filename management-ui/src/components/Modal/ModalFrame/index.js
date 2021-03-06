// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect, useCallback } from 'react'
import { bool, func, shape, oneOf } from 'prop-types'
import { CSSTransition } from 'react-transition-group'
import ModalContent from '../ModalContent'
import {
  MASK_ANIMATION_SPEED_ENTER,
  MASK_ANIMATION_DELAY_EXIT,
  MASK_ANIMATION_SPEED_EXIT,
  MODAL_ANIMATION_DELAY_ENTER,
  MODAL_ANIMATION_SPEED_ENTER,
  MODAL_ANIMATION_SPEED_EXIT,
} from '../index'
import { Container, HeightLimiter, Mask } from './index.styles'

const ESCAPE_EVENT_KEY = 'Escape'

const ModalFrame = ({
  isVisible,
  handleClose,
  allowUserToClose,
  ...passProps
}) => {
  const [inProp, setInProp] = useState(false)

  // In case modal is placed within clickable element
  const stopPropagation = (evt) => evt.stopPropagation()

  const startClose = useCallback(() => {
    setInProp(false)
  }, [setInProp])

  const handleUserClose = useCallback(
    (evt) => {
      if (!allowUserToClose) return
      startClose(evt)
    },
    [allowUserToClose, startClose],
  )

  useEffect(() => {
    setInProp(isVisible)
  }, [isVisible])

  useEffect(() => {
    const keydownHandler = (evt) => {
      if (!isVisible) return
      if (evt.key === ESCAPE_EVENT_KEY) {
        evt.stopPropagation()
        handleUserClose(evt)
      }
    }

    document.addEventListener('keydown', keydownHandler, true)
    return () => document.removeEventListener('keydown', keydownHandler, true)
  }, [isVisible, handleUserClose])

  return (
    <>
      {isVisible && (
        <Container onClick={stopPropagation}>
          <HeightLimiter alignItems={passProps.verticalAlignCss.alignItems}>
            <CSSTransition
              in={inProp}
              classNames="mask"
              timeout={{
                appear: 0,
                enter: MASK_ANIMATION_SPEED_ENTER,
                exit: MASK_ANIMATION_DELAY_EXIT + MASK_ANIMATION_SPEED_EXIT,
              }}
              onExited={() => {
                handleClose()
              }}
            >
              <Mask
                allowUserToClose={allowUserToClose}
                onClick={handleUserClose}
              />
            </CSSTransition>

            <CSSTransition
              in={inProp}
              classNames="modal-content"
              timeout={{
                appear: 0,
                enter:
                  MODAL_ANIMATION_DELAY_ENTER + MODAL_ANIMATION_SPEED_ENTER,
                exit: MODAL_ANIMATION_SPEED_EXIT,
              }}
            >
              <ModalContent
                {...passProps}
                close={startClose}
                userClose={handleUserClose}
                showCloseButton={allowUserToClose && passProps.showCloseButton}
              />
            </CSSTransition>
          </HeightLimiter>
        </Container>
      )}
    </>
  )
}

ModalFrame.propTypes = {
  isVisible: bool.isRequired,
  handleClose: func.isRequired,
  allowUserToClose: bool,
  verticalAlignCss: shape({
    alignItems: oneOf(['center', 'flex-start', 'flex-end']).isRequired,
  }),
}

ModalFrame.defaultProps = {
  allowUserToClose: true,
  verticalAlignCss: {
    alignItems: 'center',
  },
}

export default ModalFrame
