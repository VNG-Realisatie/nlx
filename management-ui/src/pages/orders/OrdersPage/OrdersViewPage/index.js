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

const OrdersViewPage = ({ orders }) => {
  const { t } = useTranslation()

  return orders.length ? (
    <Wrapper>
      <Table>
        <thead>
          <Table.TrHead>
            <Table.Th>{t('Order')}</Table.Th>
            <Table.Th>{t('Issued to')}</Table.Th>
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
        <small>{t('There are no active orders')}</small>
      </h3>
    </Centered>
  )
}

OrdersViewPage.propTypes = {
  orders: array,
}

OrdersViewPage.defaultProps = {
  orders: [],
}

export default observer(OrdersViewPage)
