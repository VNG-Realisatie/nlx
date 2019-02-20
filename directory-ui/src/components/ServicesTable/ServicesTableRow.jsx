import React from 'react'
import { oneOf, string } from 'prop-types'
import { Link } from 'react-router-dom'
import copy from 'copy-text-to-clipboard'

import Table from '../Table'
import StatusIcon from './Icons/StatusIcon/StatusIcon'
import IconButton from '../IconButton/IconButton'
import DocsIcon from './Icons/DocsIcon/DocsIcon'
import LinkIcon from './Icons/LinkIcon/LinkIcon'
import { StyledServiceTableRow, StyledApiTypeLabel } from './ServiceTableRow.styles'

export const apiAddressForService = (name, organization) =>
  `http://{your-outway-address}:12018/${organization}/${name}`

const ServicesTableRow = ({ status, organization, name, apiType }) =>
    <StyledServiceTableRow status={status}>
      <Table.BodyCell align="center" padding="none"><StatusIcon disabled={status === 'offline'} /></Table.BodyCell>
      <Table.BodyCell>{ organization }</Table.BodyCell>
      <Table.BodyCell>{ name }</Table.BodyCell>
      <Table.BodyCell align="right">
        {
          apiType &&
            <StyledApiTypeLabel status={status}>{ apiType }</StyledApiTypeLabel>
        }
      </Table.BodyCell>
      <Table.BodyCell padding="none" border="left">
        <IconButton dataTest="link-icon"
                    onClick={() => copy(apiAddressForService(name, organization))}
          >
            <LinkIcon />
          </IconButton>
      </Table.BodyCell>
      <Table.BodyCell padding="none" border="left">
        {
          apiType ?
            <IconButton as={Link} to={`/documentation/${organization}/${name}`}>
              <DocsIcon />
            </IconButton> :
            <IconButton disabled><DocsIcon /></IconButton>
        }
      </Table.BodyCell>
    </StyledServiceTableRow>

ServicesTableRow.propTypes = {
  status: oneOf(['online', 'offline']).isRequired,
  organization: string.isRequired,
  name: string.isRequired,
}

export default ServicesTableRow
