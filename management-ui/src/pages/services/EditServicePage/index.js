// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useEffect, useState, useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, ToasterContext } from '@commonground/design-system'
import { useNavigate, useParams } from 'react-router-dom'
import { observer } from 'mobx-react'
import serviceActions from '../ServicesPage/serviceActions'
import ServiceForm from '../../../components/ServiceForm'
import PageTemplate from '../../../components/PageTemplate'
import LoadingMessage from '../../../components/LoadingMessage'
import { useServiceStore } from '../../../hooks/use-stores'
import { StyledUpdatedError } from './index.styles'

const EditServicePage = () => {
  const { name } = useParams()
  const { t } = useTranslation()
  const serviceStore = useServiceStore()
  const [updateError, setUpdatedError] = useState(null)
  const navigate = useNavigate()
  const [service, setService] = useState(null)
  const { showToast } = useContext(ToasterContext)

  useEffect(() => {
    if (serviceStore.isInitiallyFetched) {
      setService(serviceStore.getByName(name))
    }
  }, [serviceStore.isInitiallyFetched]) // eslint-disable-line react-hooks/exhaustive-deps

  const submitService = async (formData) => {
    try {
      setUpdatedError(null)
      await serviceStore.update(formData)

      navigate(`/services/${service.name}?lastAction=${serviceActions.EDITED}`)
    } catch (err) {
      let message = ''

      if (err.response) {
        const res = await err.response.json()
        message = res.message

        if (err.response.status === 403) {
          message = t(`You don't have the required permission.`)
        } else {
          window.scrollTo(0, 0)
          setUpdatedError(message)
        }
      }

      showToast({
        title: t('Failed to update the service'),
        body: message,
        variant: 'error',
      })
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo={`/services/${name}`}
        title={t('Edit service')}
      />

      {!serviceStore.isInitiallyFetched ? (
        <LoadingMessage />
      ) : serviceStore.error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the service', { name })}
        </Alert>
      ) : service ? (
        <>
          {updateError ? (
            <StyledUpdatedError
              title={t('Failed to update the service')}
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
