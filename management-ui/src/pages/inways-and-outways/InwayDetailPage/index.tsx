// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import Inway from '../../../types/Inway'
import { useInwayStore, useApplicationStore } from '../../../hooks/use-stores'
import InwayDetailPageView from './InwayDetailPageView'

interface InwayDetailPageProps {
  parentUrl: string
  inway: Inway
}

const InwayDetailPage: React.FC<InwayDetailPageProps> = ({
  parentUrl = '/inways-and-outways',
  inway,
}) => {
  const { name } = useParams<{ name: string }>()
  const { t } = useTranslation()
  const history = useHistory()
  const { removeInway } = useInwayStore()
  const { applicationStore } = useApplicationStore()
  const { showToast } = useContext(ToasterContext)
  const close = () => history.push(parentUrl)

  const handleRemoveInway = async () => {
    try {
      await removeInway(inway.name)

      // Update isOrganizationInwaySet if needed, to trigger the warning banner
      const settings = await applicationStore.getGeneralSettings()
      applicationStore.updateOrganizationInway({
        isOrganizationInwaySet: !!settings.organizationInway,
      })
    } catch (err) {
      showToast({
        title: t('Failed to remove the inway'),
        // body: err.message,
        body: err,
        variant: 'error',
      })
    }
  }

  const Content = () => {
    if (inway) {
      return (
        <InwayDetailPageView inway={inway} removeHandler={handleRemoveInway} />
      )
    }
    return (
      <Alert variant="error" data-testid="error-message">
        {t('Failed to load the details for this inway', { name })}
      </Alert>
    )
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
        <Content />
      </Drawer.Content>
    </Drawer>
  )
}

export default InwayDetailPage
