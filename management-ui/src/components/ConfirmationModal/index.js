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

const ConfirmationModal = ({
  isVisible,
  handleChoice,
  children,
  title,
  cancelText,
  okText,
}) => {
  const [choice, setChoice] = useState(null)
  const { t } = useTranslation()

  const makeChoice = (isConfirmed, closeModal) => () => {
    setChoice(isConfirmed)
    closeModal()
  }

  const afterModalClose = () => {
    handleChoice(choice)
  }

  return (
    <Modal
      isVisible={isVisible}
      handleClose={afterModalClose}
      title={title || t('Are you sure?')}
      width="480px"
      verticalAlign={{
        from: 'top',
        offset: '8rem',
      }}
      autoFocus
      render={({ closeModal }) => (
        <>
          {children}
          <Footer>
            <Button
              type="button"
              variant="secondary"
              onClick={makeChoice(false, closeModal)}
              data-autofocus
            >
              {cancelText || t('Cancel')}
            </Button>
            <Button type="button" onClick={makeChoice(true, closeModal)}>
              {okText || t('Ok')}
            </Button>
          </Footer>
        </>
      )}
    />
  )
}

ConfirmationModal.propTypes = {
  isVisible: bool.isRequired,
  handleChoice: func.isRequired,
  children: node,
  title: string,
  cancelText: string,
  okText: string,
}

export const useConfirmationModal = (props) => {
  const [showModal, setShowModal] = useState(false)
  const choicePromise = useRef(null)

  const show = () => {
    choicePromise.current = deferredPromise()
    setShowModal(true)
    return choicePromise.current
  }

  const handleChoice = (isConfirmed) => {
    if (choicePromise === null) {
      return new Error(
        "Can't handle choice when ConfirmationModal is not shown",
      )
    }

    choicePromise.current.resolve(isConfirmed)
    setShowModal(false)
  }

  const confirmProps = {
    ...props,
    isVisible: showModal,
    handleChoice,
  }

  return [() => React.createElement(ConfirmationModal, confirmProps), show]
}
