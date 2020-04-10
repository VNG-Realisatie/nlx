// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useState } from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'
import ServiceRepository from '../../domain/service-repository'
import AddServiceForm from './AddServiceForm'
import {
  StyledBackButton,
  StyledIconChevron,
  StyledTitle,
} from './index.styles'

const AddServicePage = ({ createHandler }) => {
  const { t } = useTranslation()
  const [isAdded, setIsAdded] = useState(false)
  const [error, setError] = useState(null)

  const submitService = (service) => {
    // placeholder until we've implemented adding authorizations in the form
    service.authorizationSettings = service.authorizationSettings || {}
    service.authorizationSettings.authorizations = []

    createHandler(service)
      .then(() => {
        setIsAdded(true)
        setError(null)
      })
      .catch((err) => {
        setIsAdded(false)
        setError(err.message)
      })
  }

  return (
    <PageTemplate>
      <p>
        <StyledBackButton to="/services" aria-label={t('Back')}>
          <StyledIconChevron />
          {t('Back')}
        </StyledBackButton>
      </p>

      <StyledTitle>{t('Add new service')}</StyledTitle>
      {error ? (
        <Alert
          title={t('Failed adding service')}
          variant="error"
          data-testid="error-message"
          role="alert"
        >
          {error}
        </Alert>
      ) : null}

      {isAdded && !error ? (
        <Alert variant="success" data-testid="error-message" role="alert">
          {t('The service has been added.')}
        </Alert>
      ) : null}

      {!isAdded ? (
        <AddServiceForm
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
