import React from 'react'
import { TableBodyCell } from '../Table'
import { oneOf } from 'prop-types'
import StatusIcon from './Icons/StatusIcon/StatusIcon'
import DocsIcon from './Icons/DocsIcon/DocsIcon'
import LinkIcon from './Icons/LinkIcon/LinkIcon'
import { StyledServiceTableRow, StyledApiTypeLabel } from './ServiceTableRow.styles'

const statusToIconColor = status =>
  status === 'online' ? 'blue' : 'grey'

const ServicesTableRow = ({ status, organization, name, apiType, apiAddress }) =>
    <StyledServiceTableRow status={status}>
      <TableBodyCell align="center"><StatusIcon status={status} /></TableBodyCell>
      <TableBodyCell>{ organization }</TableBodyCell>
      <TableBodyCell>{ name }</TableBodyCell>
      <TableBodyCell align="right">
        {
          apiType ?
            <StyledApiTypeLabel status={status}>{ apiType }</StyledApiTypeLabel> : '-'
        }
      </TableBodyCell>
      <TableBodyCell align="center">
        <a href={apiAddress} style={({ lineHeight: '1rem' })} target="_blank" rel="noopener noreferrer">
          <LinkIcon color={statusToIconColor(status)} />
        </a>
        <DocsIcon color={statusToIconColor(status)} />
      </TableBodyCell>
    </StyledServiceTableRow>

ServicesTableRow.propTypes = {
  status: oneOf(['online', 'offline']).isRequired
}

export default ServicesTableRow
