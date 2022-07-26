// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useOutwayStore } from '../../../hooks/use-stores'
import OutwayDetailPageView from './OutwayDetailPageView'

const OutwayDetailPage = () => {
  const { name } = useParams()
  const { t } = useTranslation()
  const navigate = useNavigate()
  const outwayStore = useOutwayStore()
  const { removeOutway } = useOutwayStore()
  const { showToast } = useContext(ToasterContext)

  const close = () => navigate('/inways-and-outways/outways')
  const outway = outwayStore.getByName(name)

  const handleRemoveOutway = async () => {
    try {
      await removeOutway(outway.name)

      close()

      showToast({
        title: outway.name,
        body: t('The outway has been removed'),
        variant: 'success',
      })
    } catch (error) {
      let message = error.message

      if (error.response && error.response.status === 403) {
        message = t(`You don't have the required permission.`)
      }

      showToast({
        title: t('Failed to remove the outway'),
        body: message,
        variant: 'error',
      })
    }
  }

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
          <OutwayDetailPageView
            outway={outway}
            removeHandler={handleRemoveOutway}
          />
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
