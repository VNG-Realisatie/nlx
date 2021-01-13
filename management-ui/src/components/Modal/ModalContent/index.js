// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { bool, node, shape, string, func } from 'prop-types'
import FocusLock from 'react-focus-lock'
import cssUnitOrEmpty from '../propTypeCssUnitOrEmpty'
import ModalHeader from './ModalHeader'
import { ModalPosition, Window, Content } from './index.styles'

const ModalContent = ({
  id,
  autoFocus,
  width,
  verticalAlignCss,
  close,
  render,
  children,
  ...headerProps
}) => (
  <ModalPosition width={width} transform={verticalAlignCss.transform}>
    <FocusLock autoFocus={autoFocus} returnFocus>
      <Window
        role="dialog"
        aria-modal="true"
        aria-labelledby={headerProps.title ? `title-${id}` : `content-${id}`}
      >
        <Content id={`content-${id}`}>
          {render ? render({ closeModal: close }) : children}
        </Content>
        <ModalHeader {...headerProps} id={id} />
      </Window>
    </FocusLock>
  </ModalPosition>
)

ModalContent.propTypes = {
  id: string.isRequired,
  // Set `autoFocus` in combination with `data-autofocus` to focus on a specific element
  autoFocus: bool,
  width: cssUnitOrEmpty,
  verticalAlignCss: shape({
    transform: string.isRequired,
  }),
  close: func.isRequired,
  render: func,
  children: node,
}

ModalContent.defaultProps = {
  autoFocus: false,
  width: '',
  verticalAlignCss: shape({
    transform: '',
  }),
}

export default ModalContent
