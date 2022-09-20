// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect, useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router-dom'
import { Alert, ToasterContext } from '@commonground/design-system'
import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import useStores, { useOrderStore } from '../../../hooks/use-stores'
import OrderForm from '../components/OrderForm'
import LoadingMessage from '../../../components/LoadingMessage'

const AddOrderPage = () => {
  const { t } = useTranslation()
  const orderStore = useOrderStore()
  const navigate = useNavigate()
  const [error, setError] = useState(null)
  const { directoryServicesStore } = useStores()
  const { showToast } = useContext(ToasterContext)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    async function fetchData() {
      try {
        await directoryServicesStore.syncAllOutgoingAccessRequests()
      } catch (error) {
        showToast({
          title: t('Failed to synchronize access states'),
          variant: 'error',
        })
      }

      await directoryServicesStore.fetchAll()

      setIsLoading(false)
    }

    fetchData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const servicesWithAccess = directoryServicesStore.servicesWithAccess

  const noAccessibleServices =
    directoryServicesStore.isInitiallyFetched &&
    !servicesWithAccess.length &&
    !error

  const submitOrder = async (formData) => {
    try {
      await orderStore.create(formData)
      navigate(`/orders?lastAction=added`)
    } catch (err) {
      let message = err.message

      if (err.response && err.response.status === 403) {
        message = t(`You don't have the required permission.`)
      }

      window.scrollTo(0, 0)
      setError(message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo="/orders"
        title={t('New order')}
      />

      {isLoading ? (
        <LoadingMessage />
      ) : (
        <>
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
            submitButtonText={t('Add order')}
            onSubmitHandler={submitOrder}
          />
        </>
      )}
    </PageTemplate>
  )
}

export default observer(AddOrderPage)
