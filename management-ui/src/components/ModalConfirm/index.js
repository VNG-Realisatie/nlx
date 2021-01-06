// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
// TODO: an API like this would be awesome ↓
// const askConfirmation = useConfirmationWindow('Text or Component')
// if (await askConfirmation()) {}
//
import React from 'react'
import { func, string, bool, node } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import Modal from '../Modal'
import { Footer } from './index.styles'

const ModalConfirm = ({
  isVisible,
  onChoice,
  children,
  titleText,
  cancelText,
  okText,
}) => {
  const { t } = useTranslation()

  const handleChoice = (isConfirmed) => () => {
    onChoice(isConfirmed)
  }

  return (
    <Modal
      isVisible={isVisible}
      handleClose={handleChoice(false)}
      title={titleText || t('Are you sure?')}
      width="480px"
      verticalAlign={{
        from: 'top',
        offset: '8rem',
      }}
      autoFocus
    >
      {children}
      <Footer>
        <Button
          variant="secondary"
          onClick={handleChoice(false)}
          data-autofocus
        >
          {cancelText || t('Cancel')}
        </Button>
        <Button onClick={handleChoice(true)}>{okText || t('Ok')}</Button>
      </Footer>
    </Modal>
  )
}

ModalConfirm.propTypes = {
  isVisible: bool.isRequired,
  onChoice: func.isRequired,
  children: node,
  titleText: string,
  cancelText: string,
  okText: string,
}

ModalConfirm.defaultProps = {}

export default ModalConfirm
