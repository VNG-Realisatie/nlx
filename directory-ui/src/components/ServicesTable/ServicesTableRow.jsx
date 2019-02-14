import React from 'react'
import { TableBodyCell } from '../Table'
import { oneOf, string } from 'prop-types'
import StatusIcon from './Icons/StatusIcon/StatusIcon'
import DocsIcon from './Icons/DocsIcon/DocsIcon'
import LinkIcon from './Icons/LinkIcon/LinkIcon'
import { StyledServiceTableRow, StyledApiTypeLabel } from './ServiceTableRow.styles'

const statusToIconColor = status =>
  status === 'online' ? 'blue' : 'grey'

const ServicesTableRow = ({ status, organization, name, apiType, apiAddress, docsAddress }) =>
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
      <TableBodyCell>
        <a href={apiAddress} style={({ lineHeight: '1rem', borderLeft: '1px solid #F0F2F7' })} target="_blank" rel="noopener noreferrer">
          <LinkIcon color={statusToIconColor(status)} />
        </a>
      </TableBodyCell>
      <TableBodyCell>
        <a href={docsAddress} style={({ lineHeight: '1rem' })} target="_blank" rel="noopener noreferrer">
          <DocsIcon color={statusToIconColor(status)} />
        </a>
      </TableBodyCell>
    </StyledServiceTableRow>

ServicesTableRow.propTypes = {
  status: oneOf(['online', 'offline']).isRequired,
  organization: string.isRequired,
  name: string.isRequired,
  apiAddress: string.isRequired,
  docsAddress: string.isRequired,
}

export default ServicesTableRow
