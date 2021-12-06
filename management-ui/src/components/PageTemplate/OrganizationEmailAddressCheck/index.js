// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { useStores } from '../../../hooks/use-stores'
import GlobalAlert from '../../GlobalAlert'
import { StyledLink } from './index.styles'

const OrganizationEmailAddressCheck = () => {
  const { t } = useTranslation()
  const { applicationStore } = useStores()

  const displayMessage =
    applicationStore.isOrganizationEmailAddressSet === false

  if (displayMessage) {
    return (
      <GlobalAlert>
        {t('Please set an organization email address.')}
        <StyledLink to="/settings/general">{t('Go to settings')}</StyledLink>
      </GlobalAlert>
    )
  }

  return null
}

export default observer(OrganizationEmailAddressCheck)
