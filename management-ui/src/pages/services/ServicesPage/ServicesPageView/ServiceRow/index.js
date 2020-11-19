// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array, bool, shape, string } from 'prop-types'
import { Table } from '@commonground/design-system'
import {
  ServiceVisibilityMessage,
  showServiceVisibilityAlert,
} from '../../../../../components/ServiceVisibilityAlert'
import AmountOfIncomingRequestsLabel from './AmountOfIncomingRequestsLabel'
import { TdAlignRight, StyledTdIncomingRequests } from './index.styles'

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
      <TdAlignRight data-testid="warning-cell">
        {showServiceVisibilityAlert({ internal, inways }) ? (
          <ServiceVisibilityMessage />
        ) : null}
      </TdAlignRight>
      <StyledTdIncomingRequests>
        {service.incomingAccessRequestsCount > 0 ? (
          <AmountOfIncomingRequestsLabel
            count={service.incomingAccessRequestsCount}
          />
        ) : null}
      </StyledTdIncomingRequests>
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

export default ServiceRow
