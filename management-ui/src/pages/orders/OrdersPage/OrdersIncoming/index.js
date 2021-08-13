// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { array } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import Table from '../../../../components/Table'
import OrderRow from './OrderRow'
import { Wrapper, CellServices, Centered } from './index.styles'

const OrdersIncomingView = ({ orders }) => {
  const { t } = useTranslation()

  return orders.length ? (
    <Wrapper>
      <Table>
        <thead>
          <Table.TrHead>
            <Table.Th>{t('Order')}</Table.Th>
            <Table.Th>{t('Issued by')}</Table.Th>
            <CellServices as={Table.Th}>
              {t('Requestable services')}
            </CellServices>
            <Table.Th>{t('Valid until')}</Table.Th>
          </Table.TrHead>
        </thead>
        <tbody>
          {orders.map((order) => (
            <OrderRow key={order.reference} order={order} />
          ))}
        </tbody>
      </Table>
    </Wrapper>
  ) : (
    <Centered>
      <h3>
        <small>{t("You haven't received any orders yet")}</small>
      </h3>
      <p>
        <small>
          {t(
            'Use this to review and accept requests made on behalve of your services',
          )}
        </small>
      </p>
    </Centered>
  )
}

OrdersIncomingView.propTypes = {
  orders: array,
}

OrdersIncomingView.defaultProps = {
  orders: [],
}

export default observer(OrdersIncomingView)
