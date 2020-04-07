// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import React, { useState } from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'
import ServiceRepository from '../../domain/service-repository'
import AddServiceForm from './AddServiceForm'

const AddServicePage = ({ createHandler }) => {
  const { t } = useTranslation()
  const [isAdded, setIsAdded] = useState(false)
  const [error, setError] = useState(null)

  const submitService = (service) => {
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
    <PageTemplate title={t('Nieuwe service toevoegen')}>
      {error ? (
        <Alert
          title={t('Toevoegen mislukt')}
          variant="error"
          data-testid="error-message"
          role="alert"
        >
          {error}
        </Alert>
      ) : null}

      {isAdded && !error ? (
        <Alert variant="success" data-testid="error-message" role="alert">
          {t('De service is toegevoegd.')}
        </Alert>
      ) : null}

      {!isAdded ? (
        <AddServiceForm onSubmitHandler={(values) => submitService(values)} />
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
