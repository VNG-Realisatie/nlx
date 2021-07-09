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
  const { applicationStore, servicesStore } = useStores()

  const displayMessage =
    applicationStore.isOrganizationInwaySet === false &&
    servicesStore.services.length

  return displayMessage ? (
    <GlobalAlert>
      {t(
        'Access requests can not be received. Please specify which inway should handle access requests.',
      )}
      <StyledLink to="/settings/general">{t('Go to settings')}</StyledLink>
    </GlobalAlert>
  ) : null
}

export default observer(OrganizationInwayCheck)
