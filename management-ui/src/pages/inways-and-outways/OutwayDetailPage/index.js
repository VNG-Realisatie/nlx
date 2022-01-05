// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useOutwayStore } from '../../../hooks/use-stores'
import OutwayDetailPageView from './OutwayDetailPageView'

const OutwayDetailPage = () => {
  const { name } = useParams()
  const { t } = useTranslation()
  const navigate = useNavigate()
  const outwayStore = useOutwayStore()

  const close = () => navigate('/inways-and-outways/outways')
  const outway = outwayStore.getByName(name)

  return (
    <Drawer noMask closeHandler={close} data-testid="outway-detail-page">
      <Drawer.Header
        as="header"
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="outway-name"
      />

      <Drawer.Content>
        {outway ? (
          <OutwayDetailPageView outway={outway} />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the details for this outway', { name })}
          </Alert>
        )}
      </Drawer.Content>
    </Drawer>
  )
}

export default OutwayDetailPage
