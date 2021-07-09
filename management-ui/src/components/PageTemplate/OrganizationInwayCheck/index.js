// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { useStores } from '../../../hooks/use-stores'
import GlobalAlert from '../../GlobalAlert'
import { StyledLink } from './index.styles'

const OrganizationInwayCheck = () => {
  const { t } = useTranslation()
  const { applicationStore, servicesStore, orderStore } = useStores()

  const displayMessage =
    applicationStore.isOrganizationInwaySet === false &&
    (servicesStore.services.length || orderStore.outgoingOrders.length)

  return displayMessage ? (
    <GlobalAlert>
      {t(
        'Please select an organization inway. At the moment access requests can not be received and outgoing orders can not be retrieved by other organizations.',
      )}
      <StyledLink to="/settings/general">{t('Go to settings')}</StyledLink>
    </GlobalAlert>
  ) : null
}

export default observer(OrganizationInwayCheck)
