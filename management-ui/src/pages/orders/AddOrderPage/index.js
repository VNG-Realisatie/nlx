// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { useHistory } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../../components/PageTemplate'
import useStores, { useOrderStore } from '../../../hooks/use-stores'
import OrderForm from './components/OrderForm'

const AddOrderPage = () => {
  const { t } = useTranslation()
  const { create } = useOrderStore()
  const history = useHistory()
  const [error, setError] = useState(null)
  const [warning, setWarning] = useState(null)
  const { directoryServicesStore } = useStores()

  useEffect(() => {
    directoryServicesStore.fetchAll()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const serviceNames = directoryServicesStore.servicesWithAccess.map(
    (service) => ({
      service: service.serviceName,
      organization: service.organizationName,
    }),
  )

  if (!warning && !serviceNames.length) {
    setWarning('Check the directory to see which services are accessible')
  }

  const submitOrder = async (formData) => {
    try {
      await create(formData)
      history.push(`/orders?lastAction=added`)
    } catch (err) {
      window.scrollTo(0, 0)
      setError(err.message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo="/orders"
        title={t('New order')}
      />

      {warning && (
        <Alert
          title={t('No services were found')}
          variant="warning"
          data-testid="warning-message"
        >
          {warning}
        </Alert>
      )}

      {error && (
        <Alert
          title={t('Failed to add order')}
          variant="error"
          data-testid="error-message"
        >
          {error}
        </Alert>
      )}

      <OrderForm services={serviceNames} onSubmitHandler={submitOrder} />
    </PageTemplate>
  )
}

export default AddOrderPage
