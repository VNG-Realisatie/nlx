// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { Link, useLocation } from 'react-router-dom'
import EditButton from '../../../../../../components/EditButton/index'
import OutgoingOrderModel from '../../../../../../stores/models/OutgoingOrderModel'
import { StyledContainer } from './index.styles'

interface ButtonsProps {
  revokeHandler: () => void
  order: OutgoingOrderModel
}

const EditRevoke: React.FC<ButtonsProps> = ({ revokeHandler, order }) => {
  const { t } = useTranslation()
  const location = useLocation()

  if (!order) {
    return null
  }

  const RevokeButton = () => {
    if (order.revokedAt) {
      return null
    }
    return (
      <Button onClick={revokeHandler} aria-label={t('Revoke')} variant="danger">
        {t('Revoke')}
      </Button>
    )
  }

  return (
    <StyledContainer>
      <EditButton
        as={Link}
        to={`${location.pathname}/edit`}
        title={t('Edit order')}
      />
      <RevokeButton />
    </StyledContainer>
  )
}

export default EditRevoke
