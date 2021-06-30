// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, ToasterContext } from '@commonground/design-system'

import { useApplicationStore } from '../../../hooks/use-stores'
import usePromise from '../../../hooks/use-promise'
import LoadingMessage from '../../../components/LoadingMessage'
import Form from './Form'

const GeneralSettings = () => {
  const applicationStore = useApplicationStore()
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

  const {
    isReady,
    error,
    result: settings,
  } = usePromise(applicationStore.getGeneralSettings)

  const updateSettings = async (values) => {
    try {
      await applicationStore.updateGeneralSettings(values)

      applicationStore.updateOrganizationInway({
        isOrganizationInwaySet: !!values.organizationInway,
      })

      showToast({
        body: t('Successfully updated the settings'),
        variant: 'success',
      })
    } catch (err) {
      showToast({
        body: t('Failed to update the settings'),
        variant: 'error',
      })
    }
  }

  return (
    <>
      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the settings')}
        </Alert>
      ) : settings ? (
        <Form initialValues={settings} onSubmitHandler={updateSettings} />
      ) : null}
    </>
  )
}

export default GeneralSettings
