// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Route, Routes } from 'react-router-dom'
import Table from '../../../../components/Table'
import OutgoingOrderModel from '../../../../stores/models/OutgoingOrderModel'
import OrderRow from './OrderRow'
import { Wrapper, CellServices, Centered } from './index.styles'
import OrderDetailPage from './OrderDetailPage'

const Outgoing = ({ orders }) => {
  const { t } = useTranslation()

  return orders.length ? (
    <>
      <Wrapper>
        <Table withLinks>
          <thead>
            <Table.TrHead>
              <Table.Th>{t('Status')}</Table.Th>
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
      <Routes>
        <Route path=":delegatee/:reference" element={<OrderDetailPage />} />
      </Routes>
    </>
  ) : (
    <Centered>
      <h3>
        <small>{t("You don't have any issued orders yet")}</small>
      </h3>
      <p>
        <small>
          {t(
            'Use this to allow other organizations to request certain services on your behalve',
          )}
        </small>
      </p>
    </Centered>
  )
}

Outgoing.propTypes = {
  orders: arrayOf(instanceOf(OutgoingOrderModel)),
}

Outgoing.defaultProps = {
  orders: [],
}

export default observer(Outgoing)
