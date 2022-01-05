// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useOrderStore } from '../../../../../hooks/use-stores'
import { SubTitle } from './index.styles'
import OrderDetailView from './OrderDetailView'

const OrderDetailPage = () => {
  const { delegator, reference } = useParams()
  const { t } = useTranslation()
  const navigate = useNavigate()
  const orderStore = useOrderStore()

  const order = orderStore.getIncoming(delegator, reference)
  const close = () => navigate('/orders/incoming')

  return (
    <Drawer noMask closeHandler={close}>
      <Drawer.Header
        as="header"
        title={order ? order.description : t('Order not found')}
        closeButtonLabel={t('Close')}
        data-testid="order-name"
      />

      {order && (
        <SubTitle>
          {t('Issued by delegator', { delegator: order.delegator })}
        </SubTitle>
      )}

      <Drawer.Content>
        {!order ? (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the order issued by delegator', {
              reference,
              delegator,
            })}
          </Alert>
        ) : (
          <OrderDetailView order={order} />
        )}
      </Drawer.Content>
    </Drawer>
  )
}

export default OrderDetailPage
