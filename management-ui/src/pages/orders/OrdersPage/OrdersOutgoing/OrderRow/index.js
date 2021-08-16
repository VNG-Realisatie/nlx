// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { object } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../../../components/Table'
import { CellServices } from '../index.styles'
import StatusIcon from '../../StatusIcon'
import { Cell, List, Item, OrganizationName, Separator } from './index.styles'

const OrderRow = ({ order }) => {
  const { t } = useTranslation()
  return (
    <Table.Tr
      to={`/orders/outgoing/${encodeURIComponent(
        order.delegatee,
      )}/${encodeURIComponent(order.reference)}`}
    >
      <Cell>
        <StatusIcon
          active={
            !(
              order.revokedAt ||
              order.validFrom > new Date() ||
              order.validUntil < new Date()
            )
          }
        />
      </Cell>
      <Cell>{order.description}</Cell>
      <Cell>{order.delegatee}</Cell>
      <CellServices>
        <List>
          {order.services.map((service, i) => (
            <Item
              key={i}
              title={`${service.organization} - ${service.service}`}
            >
              <OrganizationName>{service.organization}</OrganizationName>
              <Separator> - </Separator>
              {service.service}
            </Item>
          ))}
        </List>
      </CellServices>
      <Cell>{t('date', { date: new Date(order.validUntil) })}</Cell>
    </Table.Tr>
  )
}

OrderRow.propTypes = {
  order: object,
}

export default OrderRow
