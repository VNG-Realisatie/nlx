// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import { useHistory, useParams } from 'react-router-dom'

import { observer } from 'mobx-react'
import serviceActions from '../ServicesPage/serviceActions'
import ServiceForm from '../../../components/ServiceForm'
import PageTemplate from '../../../components/PageTemplate'
import LoadingMessage from '../../../components/LoadingMessage'
import useService from '../use-service'
import { StyledUpdatedError } from './index.styles'

const EditServicePage = () => {
  const { name } = useParams()
  const { t } = useTranslation()
  const [service, error, isReady] = useService(name)
  const [updateError, setUpdatedError] = useState(null)
  const history = useHistory()

  const submitService = async (values) => {
    try {
      setUpdatedError(null)
      const updatedService = await service.update(values)
      history.push(
        `/services/${updatedService.name}?lastAction=${serviceActions.EDITED}`,
      )
    } catch (err) {
      setUpdatedError(err.message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo={`/services/${name}`}
        title={t('Edit service')}
      />

      {!isReady || (!error && !service) ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the service.', { name })}
        </Alert>
      ) : service ? (
        <>
          {updateError ? (
            <StyledUpdatedError
              title={t('Failed to update the service.')}
              variant="error"
              data-testid="error-message"
            >
              {t(`${updateError}`)}
            </StyledUpdatedError>
          ) : null}

          <ServiceForm
            initialValues={service}
            onSubmitHandler={(values) => submitService(values)}
            disableName
            submitButtonText={t('Update service')}
          />
        </>
      ) : null}
    </PageTemplate>
  )
}

export default observer(EditServicePage)
