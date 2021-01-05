// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { bool, node } from 'prop-types'
import FocusLock from 'react-focus-lock'

import ModalHeader from './ModalHeader'
import { ModalPosition, Window, Content } from './index.styles'

const ModalContent = ({
  autoFocus,
  maxWidth,
  offsetY,
  children,
  ...headerProps
}) => {
  return (
    <ModalPosition maxWidth={maxWidth} offsetY={offsetY}>
      <FocusLock autoFocus={autoFocus} returnFocus>
        <Window>
          <Content>{children}</Content>
          <ModalHeader {...headerProps} />
        </Window>
      </FocusLock>
    </ModalPosition>
  )
}

ModalContent.propTypes = {
  // Set `autoFocus` in combination with `data-autofocus` to focus on a specific element
  autoFocus: bool,
  maxWidth: cssUnitOrEmpty,
  // Vertical offset from center of screen (uses translateX, so % is relative to modal height)
  offsetY: cssUnitOrEmpty,
  children: node,
}

ModalContent.defaultProps = {
  autoFocus: false,
  maxWidth: '',
  offsetY: '',
}

function cssUnitOrEmpty(props, propName) {
  // eslint-disable-next-line security/detect-object-injection
  if (!/(^|px|rem|%)$/.test(props[propName])) {
    return new Error(
      `Invalid prop \`${propName}\` supplied to \`Modal\`. Use a css value in px, rem or %.`,
    )
  }
}

export default ModalContent
