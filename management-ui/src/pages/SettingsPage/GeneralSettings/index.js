// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, ToasterContext } from '@commonground/design-system'
import { func } from 'prop-types'
import usePromise from '../../../hooks/use-promise'
import LoadingMessage from '../../../components/LoadingMessage'
import { StyledUpdatedError } from '../../services/EditServicePage/index.styles'
import SettingsRepository from '../../../domain/settings-repository'
import Form from './Form'

const GeneralSettings = ({ getSettings, updateHandler }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const [updateError, setUpdatedError] = useState(null)
  const { isReady, error, result: settings } = usePromise(getSettings)

  const updateSettings = async (values) => {
    try {
      await updateHandler(values)

      showToast({
        body: t('Successfully updated the settings.'),
        variant: 'success',
      })

      setUpdatedError(false)
    } catch (err) {
      setUpdatedError(err.message)
    }
  }

  return (
    <>
      {!isReady || (!error && !settings) ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the settings.')}
        </Alert>
      ) : settings ? (
        <>
          {updateError ? (
            <StyledUpdatedError
              title={t('Failed to update the settings.')}
              variant="error"
              data-testid="error-message"
            >
              {t(`${updateError}`)}
            </StyledUpdatedError>
          ) : null}

          <Form
            initialValues={settings}
            onSubmitHandler={(values) => updateSettings(values)}
          />
        </>
      ) : null}
    </>
  )
}

GeneralSettings.propTypes = {
  updateHandler: func,
  getSettings: func,
}

GeneralSettings.defaultProps = {
  updateHandler: SettingsRepository.update,
  getSettings: SettingsRepository.get,
}

export default GeneralSettings
