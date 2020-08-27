// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useHistory } from 'react-router-dom'
import { Alert } from '@commonground/design-system'

import serviceActions from '../ServicesPage/serviceActions'
import PageTemplate from '../../../components/PageTemplate'
import ServiceForm from '../../../components/ServiceForm'
import { useServicesStore } from '../../../hooks/use-stores'

const AddServicePage = () => {
  const { t } = useTranslation()
  const [error, setError] = useState(null)
  const history = useHistory()
  const { addService } = useServicesStore()

  const submitService = async (service) => {
    try {
      const addedService = await addService(service)
      history.push(
        `/services/${addedService.name}?lastAction=${serviceActions.ADDED}`,
      )
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo="/services"
        title={t('Add new service')}
      />

      {error ? (
        <Alert
          title={t('Failed adding service')}
          variant="error"
          data-testid="error-message"
        >
          {error}
        </Alert>
      ) : null}

      <ServiceForm
        onSubmitHandler={submitService}
        submitButtonText={t('Add service')}
      />
    </PageTemplate>
  )
}

export default AddServicePage
