// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useInwayStore, useApplicationStore } from '../../../hooks/use-stores'
import InwayDetailPageView from './InwayDetailPageView'

const InwayDetailPage: React.FC = () => {
  const { name } = useParams<{ name: string }>()
  const { t } = useTranslation()
  const navigate = useNavigate()
  const { removeInway } = useInwayStore()
  const applicationStore = useApplicationStore()
  const inwayStore = useInwayStore()
  const { showToast } = useContext(ToasterContext)

  const close = () => navigate('/inways-and-outways/inways')
  const inway = inwayStore.getByName(name)

  const handleRemoveInway = async () => {
    try {
      await removeInway(inway.name)

      close()

      // Update isOrganizationInwaySet if needed, to trigger the warning banner
      const settings = await applicationStore.getGeneralSettings()
      applicationStore.updateOrganizationInway({
        isOrganizationInwaySet: !!settings.organizationInway,
      })
    } catch (err) {
      const e = err as { message: string }
      console.warn(e)
      showToast({
        title: t('Failed to remove the inway'),
        body: e.message,
        variant: 'error',
      })
    }
  }

  return (
    <Drawer noMask closeHandler={close}>
      <Drawer.Header
        as="header"
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="gateway-name"
      />

      <Drawer.Content>
        {inway ? (
          <InwayDetailPageView
            inway={inway}
            removeHandler={handleRemoveInway}
          />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the details for this inway', { name })}
          </Alert>
        )}
      </Drawer.Content>
    </Drawer>
  )
}

export default InwayDetailPage
