// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import Table from '../../../../../components/Table'
import OutwayModel from '../../../../../stores/models/OutwayModel'
import { StyledIconTd, StyledOutwayIcon } from './index.styles'

const OutwayRow = ({ outway, ...props }) => (
  <Table.Tr
    to={`/inways-and-outways/outways/${outway.name}`}
    name={outway.name}
    data-testid="outway-row"
    {...props}
  >
    <StyledIconTd>
      <StyledOutwayIcon />
    </StyledIconTd>
    <Table.Td>{outway.name}</Table.Td>
    <Table.Td>{outway.version}</Table.Td>
  </Table.Tr>
)

OutwayRow.propTypes = {
  outway: instanceOf(OutwayModel),
}

export default OutwayRow
