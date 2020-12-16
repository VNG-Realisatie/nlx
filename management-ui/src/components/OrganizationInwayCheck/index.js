// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect } from 'react'
import { observer } from 'mobx-react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import SettingsRepository from '../../domain/settings-repository'
import { useApplicationStore, useServicesStore } from '../../hooks/use-stores'
import GlobalAlert from '../GlobalAlert'
import { StyledLink } from './index.styles'

const OrganizationInwayCheck = ({ getSettings }) => {
  // Destructuring `applicationStore` breaks mobx reactivity
  const applicationStore = useApplicationStore()
  const { services } = useServicesStore()
  const { t } = useTranslation()

  useEffect(() => {
    const fetch = async () => {
      if (applicationStore.isOrganizationInwaySet === null) {
        try {
          const settings = await getSettings()
          applicationStore.update({
            isOrganizationInwaySet: settings.organizationInway,
          })
        } catch (e) {
          console.error(e)
        }
      }
    }

    fetch()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const render =
    applicationStore.isOrganizationInwaySet === false && services.length

  return render ? (
    <GlobalAlert>
      {t(
        'Access requests can not be received. Set which inway handles access requests.',
      )}
      <StyledLink to="/settings/general">{t('Go to settings')}</StyledLink>
    </GlobalAlert>
  ) : null
}

OrganizationInwayCheck.propTypes = {
  getSettings: func,
}

OrganizationInwayCheck.defaultProps = {
  getSettings: SettingsRepository.getGeneralSettings,
}

export default observer(OrganizationInwayCheck)
