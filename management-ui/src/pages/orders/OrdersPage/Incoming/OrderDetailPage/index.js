// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { string, instanceOf } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import IncomingOrderModel from '../../../../../stores/models/IncomingOrderModel'
import { SubTitle } from './index.styles'
import OrderDetailView from './OrderDetailView'

const OrderDetailPage = ({ parentUrl, order }) => {
  const { delegator, reference } = useParams()
  const { showToast } = useContext(ToasterContext)
  const { t } = useTranslation()
  const history = useHistory()

  const close = () => history.push(parentUrl)
  const handleRevoke = async (order) => {
    try {
      await order.revoke()
    } catch (err) {
      showToast({
        title: t('Failed to revoke the order'),
        body: err.message,
        variant: 'error',
      })
    }
  }

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
          <OrderDetailView
            order={order}
            revokeHandler={(order) => {
              handleRevoke(order)
            }}
          />
        )}
      </Drawer.Content>
    </Drawer>
  )
}

OrderDetailPage.propTypes = {
  parentUrl: string,
  order: instanceOf(IncomingOrderModel),
}

OrderDetailPage.defaultProps = {
  parentUrl: '/orders',
}

export default OrderDetailPage
