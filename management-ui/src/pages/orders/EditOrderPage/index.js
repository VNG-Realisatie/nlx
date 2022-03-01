// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import useStores from '../../../hooks/use-stores'
import PageTemplate from '../../../components/PageTemplate'
import LoadingMessage from '../../../components/LoadingMessage'
import { StyledUpdatedError } from '../../services/EditServicePage/index.styles'
import OrderForm from '../components/OrderForm'

const EditOrderPage = () => {
  const { t } = useTranslation()
  const { delegatee, reference } = useParams()
  const { directoryServicesStore, orderStore } = useStores()
  const [loadingInitial, setLoadingInitial] = useState(true)
  const [error, setError] = useState(null)
  const [updateError, setUpdatedError] = useState(null)
  const [order, setOrder] = useState(null)
  const navigate = useNavigate()

  useEffect(() => {
    const fetchInitialData = async () => {
      try {
        await Promise.all([
          orderStore.fetchOutgoing(),
          directoryServicesStore.fetchAll(),
        ])

        const orderModel = orderStore.getOutgoing(delegatee, reference)

        if (!orderModel) {
          throw new Error(
            `unable to find outgoing order for delegatee '${delegatee}' with reference '${reference}'`,
          )
        }

        setOrder(orderModel)

        setLoadingInitial(false)
      } catch (err) {
        setError(err.message)
        setLoadingInitial(false)
      }
    }
    fetchInitialData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const servicesWithAccess = directoryServicesStore.servicesWithAccess

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
      ) : orderStore.isLoading || loadingInitial || !order ? (
        <LoadingMessage />
      ) : (
        <>
          {updateError ? (
            <StyledUpdatedError
              title={t('Failed to update the order')}
              variant="error"
              data-testid="error-message"
            >
              {t(`${updateError}`)}
            </StyledUpdatedError>
          ) : null}

          <OrderForm
            isEditMode
            initialValues={{
              description: order.description,
              reference: order.reference,
              delegatee: order.delegatee,
              publicKeyPEM: order.publicKeyPEM,
              validFrom: order.validFrom,
              validUntil: order.validUntil,
              accessProofIds: order.accessProofs.map((model) => model.id),
            }}
            services={servicesWithAccess}
            submitButtonText={t('Update order')}
            onSubmitHandler={submitOrder}
          />
        </>
      )}
    </PageTemplate>
  )
}

export default observer(EditOrderPage)
