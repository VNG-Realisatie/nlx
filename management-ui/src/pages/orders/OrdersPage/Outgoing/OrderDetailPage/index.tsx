// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import OutgoingOrderModel from '../../../../../stores/models/OutgoingOrderModel'
import { useConfirmationModal } from '../../../../../components/ConfirmationModal'
import { SubTitle } from './index.styles'
import OrderDetailView from './OrderDetailView'
import EditRevoke from './EditRevoke'

interface OrderDetailPageProps {
  parentUrl: string
  order: OutgoingOrderModel
}

const OrderDetailPage: React.FC<OrderDetailPageProps> = ({
  parentUrl = '/orders/outgoing',
  order,
}) => {
  const { delegatee, reference } =
    useParams<{ delegatee: string; reference: string }>()
  const { showToast } = useContext(ToasterContext)
  const { t } = useTranslation()
  const history = useHistory()

  const close = () => history.push(parentUrl)

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
