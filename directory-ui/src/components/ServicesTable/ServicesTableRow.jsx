import React from 'react'
import { Link } from 'react-router-dom'
import copy from 'copy-text-to-clipboard'
import { TableBodyCell } from '../Table'
import { oneOf, string } from 'prop-types'
import StatusIcon from './Icons/StatusIcon/StatusIcon'
import DocsIcon from './Icons/DocsIcon/DocsIcon'
import LinkIcon from './Icons/LinkIcon/LinkIcon'
import { StyledServiceTableRow, StyledApiTypeLabel } from './ServiceTableRow.styles'

export const apiAddressForService = (name, organization) =>
  `http://{your-outway-address}:12018/${organization}/${name}`

const ServicesTableRow = ({ status, organization, name, apiType }) =>
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
      <TableBodyCell style={({ borderLeft: '1px solid #F0F2F7' })}>
          <LinkIcon dataTest="link-icon"
                    style={({ lineHeight: '1rem', cursor: 'pointer' })}
                    color="blue"
                    onClick={() => copy(apiAddressForService(name, organization))}
          />
      </TableBodyCell>
      <TableBodyCell style={({ borderLeft: '1px solid #F0F2F7' })}>
        {
          apiType ?
            <Link to={`/documentation/${organization}/${name}`} style={({ lineHeight: '1rem' })}>
              <DocsIcon color="blue" />
            </Link> :
            <DocsIcon color="grey" />
        }
      </TableBodyCell>
    </StyledServiceTableRow>

ServicesTableRow.propTypes = {
  status: oneOf(['online', 'offline']).isRequired,
  organization: string.isRequired,
  name: string.isRequired,
}

export default ServicesTableRow
