import React from 'react'
import { oneOf, string } from 'prop-types'
import { Link } from 'react-router-dom'
import copy from 'copy-text-to-clipboard'

import Table from '../Table'
import StatusIcon from './Icons/StatusIcon/StatusIcon'
import DocsIcon from './Icons/DocsIcon/DocsIcon'
import LinkIcon from './Icons/LinkIcon/LinkIcon'
import { StyledServiceTableRow, StyledApiTypeLabel } from './ServiceTableRow.styles'

export const apiAddressForService = (name, organization) =>
  `http://{your-outway-address}:12018/${organization}/${name}`

const ServicesTableRow = ({ status, organization, name, apiType }) =>
    <StyledServiceTableRow status={status}>
      <Table.BodyCell align="center"><StatusIcon status={status} /></Table.BodyCell>
      <Table.BodyCell>{ organization }</Table.BodyCell>
      <Table.BodyCell>{ name }</Table.BodyCell>
      <Table.BodyCell align="right">
        {
          apiType ?
            <StyledApiTypeLabel status={status}>{ apiType }</StyledApiTypeLabel> : '-'
        }
      </Table.BodyCell>
      <Table.BodyCell style={({ borderLeft: '1px solid #F0F2F7' })}>
          <LinkIcon dataTest="link-icon"
                    style={({ lineHeight: '1rem', cursor: 'pointer' })}
                    color="blue"
                    onClick={() => copy(apiAddressForService(name, organization))}
          />
      </Table.BodyCell>
      <Table.BodyCell style={({ borderLeft: '1px solid #F0F2F7' })}>
        {
          apiType ?
            <Link to={`/documentation/${organization}/${name}`} style={({ lineHeight: '1rem' })}>
              <DocsIcon color="blue" />
            </Link> :
            <DocsIcon color="grey" />
        }
      </Table.BodyCell>
    </StyledServiceTableRow>

ServicesTableRow.propTypes = {
  status: oneOf(['online', 'offline']).isRequired,
  organization: string.isRequired,
  name: string.isRequired,
}

export default ServicesTableRow
