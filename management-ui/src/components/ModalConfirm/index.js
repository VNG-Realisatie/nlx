// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useRef } from 'react'
import { func, string, bool, node } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import deferredPromise from '../../utils/deferred-promise'
import Modal from '../Modal'
import { Footer } from './index.styles'

const ModalConfirm = ({
  isVisible,
  handleChoice,
  children,
  titleText,
  cancelText,
  okText,
}) => {
  const { t } = useTranslation()

  const makeChoice = (isConfirmed) => () => handleChoice(isConfirmed)

  return (
    <Modal
      isVisible={isVisible}
      handleClose={makeChoice(false)}
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
        <Button variant="secondary" onClick={makeChoice(false)} data-autofocus>
          {cancelText || t('Cancel')}
        </Button>
        <Button onClick={makeChoice(true)}>{okText || t('Ok')}</Button>
      </Footer>
    </Modal>
  )
}

ModalConfirm.propTypes = {
  isVisible: bool.isRequired,
  handleChoice: func.isRequired,
  children: node,
  titleText: string,
  cancelText: string,
  okText: string,
}

export const useModalConfirm = (props) => {
  const [showModal, setShowModal] = useState(false)
  const choicePromise = useRef(null)

  const handleChoice = (isConfirmed) => {
    choicePromise.current.resolve(isConfirmed)
    setShowModal(false)
  }

  const showModalConfirm = () => {
    choicePromise.current = deferredPromise()
    setShowModal(true)
    return choicePromise.current
  }

  const confirmProps = {
    ...props,
    isVisible: showModal,
    handleChoice,
  }

  return [
    () => React.createElement(ModalConfirm, confirmProps),
    showModalConfirm,
  ]
}
