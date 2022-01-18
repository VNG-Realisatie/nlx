// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import useStores, { useOrderStore } from '../../../hooks/use-stores'
import PageTemplate from '../../../components/PageTemplate'
import LoadingMessage from '../../../components/LoadingMessage'
import { StyledUpdatedError } from '../../services/EditServicePage/index.styles'
import EditOrderForm from './components/EditOrderForm'

const EditOrderPage = () => {
  const { t } = useTranslation()
  const { delegatee, reference } = useParams()
  const { directoryServicesStore } = useStores()
  const orderStore = useOrderStore()
  const [loadingInitial, setLoadingInitial] = useState(true)
  const [error, setError] = useState(null)
  const [updateError, setUpdatedError] = useState(null)
  const [serviceNames, setServiceNames] = useState(null)
  const [order, setOrder] = useState(null)
  const navigate = useNavigate()

  useEffect(() => {
    const fetchInitialData = async () => {
      try {
        await Promise.all([
          orderStore.fetchOutgoing(),
          orderStore.fetchIncoming(),
          directoryServicesStore.fetchAll(),
        ])

        const orderModel = orderStore.outgoingOrders.getOutgoing(
          delegatee,
          reference,
        )

        if (!orderModel) {
          throw new Error('could not find order in outgoingOrders')
        }

        setOrder(orderModel)

        const fetchedServiceNames =
          directoryServicesStore.servicesWithAccess.map((service) => ({
            service: service.serviceName,
            organization: service.organization,
          }))
        setServiceNames(fetchedServiceNames)

        setLoadingInitial(false)
      } catch (err) {
        setError(err.message)
        setLoadingInitial(false)
      }
    }
    fetchInitialData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const submitOrder = async (formData) => {
    try {
      setUpdatedError(null)
      await orderStore.updateOutgoing({
        ...formData,
        delegatee,
        reference,
      })
      navigate(`/orders/outgoing/${delegatee}/${reference}`)
    } catch (err) {
      window.scrollTo(0, 0)
      setUpdatedError(err.message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo={`/orders/outgoing/${delegatee}/${reference}`}
        title={t('Edit order')}
      />
      {error ? (
        <StyledUpdatedError
          title={t('Failed to update the order')}
          variant="error"
          data-testid="error-message"
        >
          {error}
        </StyledUpdatedError>
      ) : orderStore.isLoading || loadingInitial || !order || !serviceNames ? (
        <LoadingMessage />
      ) : (
        <>
          {updateError ? (
            <StyledUpdatedError
              title={t('Failed to update the order')}
              variant="error"
              data-testid="error-message"
            >
              {t(`${updateError || ''}`)}
            </StyledUpdatedError>
          ) : null}

          <EditOrderForm
            order={order}
            services={serviceNames}
            onSubmitHandler={submitOrder}
          />
        </>
      )}
    </PageTemplate>
  )
}

export default observer(EditOrderPage)
