// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import {
  useApplicationStore,
  useServicesStore,
} from '../../../hooks/use-stores'
import GlobalAlert from '../../GlobalAlert'
import { StyledLink } from './index.styles'

const OrganizationInwayCheck = () => {
  // Destructuring `applicationStore` breaks mobx reactivity
  const applicationStore = useApplicationStore()
  const { services } = useServicesStore()
  const { t } = useTranslation()

  const displayMessage =
    applicationStore.isOrganizationInwaySet === false && services.length

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
