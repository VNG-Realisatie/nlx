// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import useStores, { useOrderStore } from '../../../hooks/use-stores'
import OrderForm from '../components/OrderForm'

const AddOrderPage = () => {
  const { t } = useTranslation()
  const { create } = useOrderStore()
  const navigate = useNavigate()
  const [error, setError] = useState(null)
  const { directoryServicesStore } = useStores()

  useEffect(() => {
    directoryServicesStore.fetchAll()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const servicesWithAccess = directoryServicesStore.servicesWithAccess

  const noAccessibleServices =
    directoryServicesStore.isInitiallyFetched &&
    !servicesWithAccess.length &&
    !error

  const submitOrder = async (formData) => {
    try {
      await create(formData)
      navigate(`/orders?lastAction=added`)
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

      {noAccessibleServices && (
        <Alert
          title={t('No services available')}
          variant="warning"
          data-testid="warning-message"
        />
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

      <OrderForm
        services={servicesWithAccess}
        submitButtonText={t('Add order')}
        onSubmitHandler={submitOrder}
      />
    </PageTemplate>
  )
}

export default observer(AddOrderPage)
