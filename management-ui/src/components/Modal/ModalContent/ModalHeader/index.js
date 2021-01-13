// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string, func, bool } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Header, Title, CloseButton, StyledIconClose } from './index.styles'

const ModalHeader = ({ id, userClose, title, showCloseButton }) => {
  const { t } = useTranslation()
  return (
    <Header hasTitle={!!title} hasCloseButton={showCloseButton}>
      {title && <Title id={`title-${id}`}>{title}</Title>}
      {showCloseButton && (
        <CloseButton onClick={userClose}>
          <StyledIconClose title={t('Close')} />
        </CloseButton>
      )}
    </Header>
  )
}

ModalHeader.propTypes = {
  id: string.isRequired,
  userClose: func.isRequired,
  title: string,
  showCloseButton: bool,
}

ModalHeader.defaultProps = {
  showCloseButton: true,
}

export default ModalHeader
