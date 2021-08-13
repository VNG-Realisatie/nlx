// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string, shape } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { SubTitle } from './index.styles'
// import OrderDetailView from './OrderDetailView'

const OrderDetailPage = ({ parentUrl, order }) => {
  const { delegatee, reference } = useParams()
  const { t } = useTranslation()
  const history = useHistory()

  const close = () => history.push(parentUrl)

  return (
    <Drawer noMask closeHandler={close}>
      <Drawer.Header
        as="header"
        title={order ? order.description : t('Order not found')}
        closeButtonLabel={t('Close')}
        data-testid="order-name"
      />
      {order ? (
        <SubTitle>
          {t('Issued to delegatee', { delegatee: order.delegatee })}
        </SubTitle>
      ) : null}

      <Drawer.Content>
        {order ? null : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the order', { reference, delegatee })}
          </Alert>
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
  parentUrl: '/orders',
}

export default OrderDetailPage
