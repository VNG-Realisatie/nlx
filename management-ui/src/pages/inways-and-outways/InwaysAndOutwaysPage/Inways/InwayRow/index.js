// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import Table from '../../../../../components/Table'
import InwayModel from '../../../../../stores/models/InwayModel'
import { StyledInwayIcon, StyledIconTd } from './index.styles'

const InwayRow = ({ inway, ...props }) => (
  <Table.Tr
    to={`/inways-and-outways/inways/${inway.name}`}
    name={inway.name}
    data-testid="inway-row"
    {...props}
  >
    <StyledIconTd>
      <StyledInwayIcon />
    </StyledIconTd>
    <Table.Td>{inway.name}</Table.Td>
    <Table.Td>{inway.hostname}</Table.Td>
    <Table.Td>{inway.selfAddress}</Table.Td>
    <Table.Td data-testid="services-count">{inway.services.length}</Table.Td>
    <Table.Td>{inway.version}</Table.Td>
  </Table.Tr>
)

InwayRow.propTypes = {
  inway: instanceOf(InwayModel),
}

export default InwayRow
