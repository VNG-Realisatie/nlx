// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useConfirmationModal } from '../../../../../components/ConfirmationModal'
import { useOrderStore } from '../../../../../hooks/use-stores'
import { SubTitle } from './index.styles'
import OrderDetailView from './OrderDetailView'
import EditRevoke from './EditRevoke'

const OrderDetailPage: React.FC = () => {
  const { delegatee, reference } =
    useParams<{ delegatee: string; reference: string }>()
  const { showToast } = useContext(ToasterContext)
  const { t } = useTranslation()
  const navigate = useNavigate()
  const orderStore = useOrderStore()

  const order = orderStore.getOutgoing(delegatee, reference)

  const close = () => navigate('/orders/outgoing')

  const [ConfirmRevokeModal, confirmRevoke] = useConfirmationModal({
    okText: t('Revoke'),
    children: <p>{t('Do you want to revoke the order?')}</p>,
  })

  const handleRevoke = async () => {
    if (await confirmRevoke()) {
      try {
        await order.revoke()
      } catch (err: any) { // eslint-disable-line
        showToast({
          title: t('Failed to revoke the order'),
          body: err.message,
          variant: 'error',
        })
      }
    }
  }

  const DrawerContent = () => {
    if (!order) {
      return (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the order for delegatee', {
            reference,
            delegatee,
          })}
        </Alert>
      )
    }
    return <OrderDetailView order={order} />
  }

  return (
    <>
      <Drawer noMask closeHandler={close}>
        <Drawer.Header
          as="header"
          title={order ? order.description : t('Order not found')}
          closeButtonLabel={t('Close')}
          data-testid="order-name"
        />

        {order && (
          <SubTitle>
            {t('Issued to delegatee', { delegatee: order.delegatee })}
          </SubTitle>
        )}

        <EditRevoke revokeHandler={handleRevoke} order={order} />

        <Drawer.Content>
          <DrawerContent />
        </Drawer.Content>
      </Drawer>

      <ConfirmRevokeModal />
    </>
  )
}

export default OrderDetailPage
