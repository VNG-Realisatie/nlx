// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useRef } from 'react'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import deferredPromise from '../../utils/deferred-promise'
import Modal from '../Modal'
import { Footer } from './index.styles'

interface UseConfirmationModalProps {
  children?: React.ReactNode
  title?: string
  cancelText?: string
  okText?: string
}

interface ConfirmationModalProps extends UseConfirmationModalProps {
  isVisible: boolean
  handleChoice: (choice: boolean | null) => void
}

const ConfirmationModal: React.FC<ConfirmationModalProps> = ({
  isVisible,
  handleChoice,
  children,
  title,
  cancelText,
  okText,
}) => {
  const [choice, setChoice] = useState<boolean | null>(null)
  const { t } = useTranslation()

  const makeChoice = (isConfirmed: boolean, closeModal: () => void) => () => {
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
      render={({ closeModal }: { closeModal: () => void }) => (
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

type UseModelConfirmation = (
  props: UseConfirmationModalProps,
) => [React.FC, () => Promise<boolean | null>]
export const useConfirmationModal: UseModelConfirmation = (props) => {
  const [showModal, setShowModal] = useState(false)
  const choicePromise = useRef<Promise<boolean | null>>(null)

  const show = async () => {
    if (!!choicePromise) {
      // This has been confirmed to work correctly
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      choicePromise.current = deferredPromise<boolean | null>()
    }
    setShowModal(true)
    return choicePromise.current
  }

  const handleChoice = (isConfirmed: boolean | null) => {
    if (choicePromise === null) {
      return new Error(
        "Can't handle choice when ConfirmationModal is not shown",
      )
    }

    // This has been confirmed to work correctly
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    choicePromise?.current?.resolve(isConfirmed)
    setShowModal(false)
  }

  const confirmProps = {
    ...props,
    isVisible: showModal,
    handleChoice,
  }

  return [() => React.createElement(ConfirmationModal, confirmProps), show]
}
