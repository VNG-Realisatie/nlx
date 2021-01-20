// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array, bool, shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { Table } from '@commonground/design-system'
import {
  ServiceVisibilityMessage,
  showServiceVisibilityAlert,
} from '../../../../../components/ServiceVisibilityAlert'
import AmountOfIncomingRequestsLabel from './AmountOfIncomingRequestsLabel'
import { WarningCell } from './index.styles'

const ServiceRow = ({ service, ...props }) => {
  const { name, internal, inways } = service

  return (
    <Table.Tr
      to={`/services/${name}`}
      name={name}
      data-testid="service-row"
      {...props}
    >
      <Table.Td>{name}</Table.Td>
      <WarningCell data-testid="warning-cell">
        {showServiceVisibilityAlert({ internal, inways }) ? (
          <ServiceVisibilityMessage />
        ) : service.incomingAccessRequestsCount > 0 ? (
          <AmountOfIncomingRequestsLabel
            count={service.incomingAccessRequestsCount}
          />
        ) : null}
      </WarningCell>
    </Table.Tr>
  )
}

ServiceRow.propTypes = {
  service: shape({
    name: string.isRequired,
    internal: bool.isRequired,
    inways: array.isRequired,
  }),
}

export default observer(ServiceRow)
