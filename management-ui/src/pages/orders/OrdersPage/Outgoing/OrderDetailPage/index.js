// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { string, shape } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { SubTitle } from './index.styles'
import OrderDetailView from './OrderDetailView'

const OrderDetailPage = ({ parentUrl, order }) => {
  const { delegatee, reference } = useParams()
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
          {t('Issued to delegatee', { delegatee: order.delegatee })}
        </SubTitle>
      )}

      <Drawer.Content>
        {!order ? (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the order for delegatee', {
              reference,
              delegatee,
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
  order: shape({
    delegatee: string.isRequired,
    reference: string.isRequired,
  }),
}

OrderDetailPage.defaultProps = {
  parentUrl: '/orders/outgoing',
}

export default OrderDetailPage
