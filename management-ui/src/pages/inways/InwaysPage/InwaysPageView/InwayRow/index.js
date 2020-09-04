// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, array } from 'prop-types'
import { Table } from '@commonground/design-system'

import { StyledInwayIcon, StyledIconTd } from './index.styles'

const InwayRow = ({ inway, ...props }) => {
  const { name, hostname, selfAddress, services, version } = inway
  const servicesCount = services ? services.length : 0

  return (
    <Table.Tr
      to={`/inways/${name}`}
      name={name}
      data-testid="inway-row"
      {...props}
    >
      <StyledIconTd>
        <StyledInwayIcon />
      </StyledIconTd>
      <Table.Td>{name}</Table.Td>
      <Table.Td>{hostname}</Table.Td>
      <Table.Td>{selfAddress}</Table.Td>
      <Table.Td data-testid="services-count">{servicesCount}</Table.Td>
      <Table.Td>{version}</Table.Td>
    </Table.Tr>
  )
}

InwayRow.propTypes = {
  inway: shape({
    name: string,
    hostname: string,
    selfAddress: string,
    services: array,
    version: string,
  }),
}

export default InwayRow
