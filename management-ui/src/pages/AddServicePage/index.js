// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useState } from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { useHistory } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'
import ServiceRepository from '../../domain/service-repository'
import ServiceForm from '../../components/ServiceForm'

const AddServicePage = ({ createHandler }) => {
  const { t } = useTranslation()
  const [isAdded, setIsAdded] = useState(false)
  const [error, setError] = useState(null)
  const history = useHistory()

  const submitService = (service) => {
    // placeholder until we've implemented adding authorizations in the form
    service.authorizationSettings = service.authorizationSettings || {}
    service.authorizationSettings.authorizations = []

    createHandler(service)
      .then((service) => {
        history.push(`/services/${service.name}?new=true`)
      })
      .catch((err) => {
        setIsAdded(false)
        setError(err.message)
      })
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

      {!isAdded ? (
        <ServiceForm
          onSubmitHandler={(values) => submitService(values)}
          submitButtonText={t('Add service')}
        />
      ) : null}
    </PageTemplate>
  )
}

AddServicePage.propTypes = {
  createHandler: func,
}

AddServicePage.defaultProps = {
  createHandler: ServiceRepository.create,
}

export default AddServicePage
