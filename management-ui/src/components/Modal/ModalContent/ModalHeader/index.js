// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string, func, bool } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Header, Title, CloseButton, StyledIconClose } from './index.styles'

const ModalHeader = ({ onClose, title, showCloseButton }) => {
  const { t } = useTranslation()
  return (
    <Header hasTitle={!!title} hasCloseButton={showCloseButton}>
      {title && <Title>{title}</Title>}
      {showCloseButton && (
        <CloseButton onClick={onClose}>
          <StyledIconClose title={t('Close')} />
        </CloseButton>
      )}
    </Header>
  )
}

ModalHeader.propTypes = {
  onClose: func.isRequired,
  title: string,
  showCloseButton: bool,
}

ModalHeader.defaultProps = {
  showCloseButton: true,
}

export default ModalHeader
