// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { bool, node, shape, string } from 'prop-types'
import FocusLock from 'react-focus-lock'
import cssUnitOrEmpty from '../propTypeCssUnitOrEmpty'
import ModalHeader from './ModalHeader'
import { ModalPosition, Window, Content } from './index.styles'

const ModalContent = ({
  autoFocus,
  width,
  verticalAlignCss,
  children,
  ...headerProps
}) => {
  return (
    <ModalPosition width={width} transform={verticalAlignCss.transform}>
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
  width: cssUnitOrEmpty,
  verticalAlignCss: shape({
    transform: string.isRequired,
  }),
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
