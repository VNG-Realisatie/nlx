// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import useStores, { useOrderStore } from '../../../hooks/use-stores'
import Order from '../../../types/Order'
import PageTemplate from '../../../components/PageTemplate'
import LoadingMessage from '../../../components/LoadingMessage'
import { StyledUpdatedError } from '../../services/EditServicePage/index.styles'
import OutgoingOrderModel from '../../../stores/models/OutgoingOrderModel'
import Service from '../../../types/Service'
import EditOrderForm from './components/EditOrderForm'

interface OrderWithStringDates {
  description: string
  delegatee: string
  reference: string
  services: Service[]
  validFrom: string
  validUntil: string
}

const EditOrderPage: React.FC = () => {
  const { t } = useTranslation()
  const { delegatee, reference } =
    useParams<{ delegatee: string; reference: string }>()
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  const { directoryServicesStore } = useStores()
  const orderStore = useOrderStore()
  const [loadingInitial, setLoadingInitial] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [updateError, setUpdatedError] = useState<string | null>(null)
  const [serviceNames, setServiceNames] = useState(null)
  const [order, setOrder] = useState<OrderWithStringDates | null>(null)
  const navigate = useNavigate()

  useEffect(() => {
    const fetchInitialData = async () => {
      try {
        await Promise.all([
          orderStore.fetchOutgoing(),
          orderStore.fetchIncoming(),
          directoryServicesStore.fetchAll(),
        ])

        const orderData = orderStore.outgoingOrders.find(
          (o: OutgoingOrderModel) =>
            o.delegatee === delegatee && o.reference === reference,
        )
        if (!orderData) {
          throw new Error('could not find order in outgoingOrders')
        }
        const orderObj = {
          delegatee: orderData.delegatee,
          reference: orderData.reference,
          description: orderData.description,
          publicKeyPem: orderData.publicKeyPem,
          validFrom: orderData.validFrom.toString(),
          validUntil: orderData.validUntil.toString(),
          services: orderData.services,
        }
        setOrder(orderObj)

        const fetchedServiceNames =
          directoryServicesStore.servicesWithAccess.map((service: Service) => ({
            service: service.serviceName,
            organization: service.organization,
          }))
        setServiceNames(fetchedServiceNames)

        setLoadingInitial(false)
      } catch (err: any) { // eslint-disable-line
        setError(err.message)
        setLoadingInitial(false)
      }
    }
    fetchInitialData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const redirectUrl = `/orders/outgoing/${delegatee}/${reference}`

  const submitOrder = async (formData: Order) => {
    try {
      setUpdatedError(null)
      await orderStore.updateOutgoing({
        ...formData,
        delegatee,
        reference,
      })
      navigate(redirectUrl)
      // eslint-disable-next-line
    } catch (err: any) {
      window.scrollTo(0, 0)
      setUpdatedError(err.message)
    }
  }

  const UpdateError = () => {
    if (!updateError) {
      return null
    }
    return (
      <StyledUpdatedError
        title={t('Failed to update the order')}
        variant="error"
        data-testid="error-message"
      >
        {t(`${updateError || ''}`)}
      </StyledUpdatedError>
    )
  }

  const Content = () => {
    if (error) {
      return (
        <StyledUpdatedError
          title={t('Failed to update the order')}
          variant="error"
          data-testid="error-message"
        >
          {error}
        </StyledUpdatedError>
      )
    }
    if (orderStore.isLoading || loadingInitial || !order || !serviceNames) {
      return <LoadingMessage />
    }
    return (
      <>
        <UpdateError />
        <EditOrderForm
          order={order}
          services={serviceNames}
          onSubmitHandler={submitOrder}
        />
      </>
    )
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo={redirectUrl}
        title={t('Edit order')}
      />
      <Content />
    </PageTemplate>
  )
}

export default observer(EditOrderPage)
