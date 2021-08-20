// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { arrayOf, instanceOf, shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../../../components/Table'
import { CellServices } from '../index.styles'
import StatusIcon from '../../StatusIcon'
import { Cell, List, Item, OrganizationName, Separator } from './index.styles'

const OrderRow = ({ order }) => {
  const { t } = useTranslation()
  return (
    <Table.Tr
      to={`/orders/incoming/${encodeURIComponent(
        order.delegator,
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
      <Cell>{order.delegator}</Cell>
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
  order: shape({
    delegator: string.isRequired,
    reference: string.isRequired,
    services: arrayOf(
      shape({
        service: string.isRequired,
        organization: string.isRequired,
      }),
    ),
    validFrom: instanceOf(Date).isRequired,
    validUntil: instanceOf(Date).isRequired,
  }),
}

export default observer(OrderRow)
