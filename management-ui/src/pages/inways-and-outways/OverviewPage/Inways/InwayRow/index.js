// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape } from 'prop-types'
import Table from '../../../../../components/Table'
import { inwayModelPropTypes } from '../../../../../stores/models/InwayModel'
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
  inway: shape(inwayModelPropTypes),
}

export default InwayRow
