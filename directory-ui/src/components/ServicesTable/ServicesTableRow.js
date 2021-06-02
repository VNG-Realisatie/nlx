// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { oneOf, string } from 'prop-types'
import { Link } from 'react-router-dom'
import copy from 'copy-text-to-clipboard'
import Table from '../Table'
import IconButton from '../IconButton'
import StatusIcon from './Icons/StatusIcon/StatusIcon'
import DocsIcon from './Icons/DocsIcon/DocsIcon'
import LinkIcon from './Icons/LinkIcon/LinkIcon'
import {
  StyledServiceTableRow,
  StyledApiTypeLabel,
} from './ServiceTableRow.styles'

export const apiUrlForService = (organization, name) =>
  `http://{your-outway-address}/${organization}/${name}`

const ServicesTableRow = ({
  status,
  organization,
  name,
  apiType,
  ...props
}) => {
  const copyApiUrl = () => {
    copy(apiUrlForService(organization, name))
  }

  return (
    <StyledServiceTableRow
      data-test="service-table-row"
      status={status}
      {...props}
    >
      <Table.BodyCell align="center" padding="none" title={status}>
        <StatusIcon status={status} />
      </Table.BodyCell>
      <Table.BodyCell>{organization}</Table.BodyCell>
      <Table.BodyCell>{name}</Table.BodyCell>
      <Table.BodyCell align="right">
        {apiType && (
          <StyledApiTypeLabel status={status}>{apiType}</StyledApiTypeLabel>
        )}
      </Table.BodyCell>
      <Table.BodyCell padding="none" border="left">
        <IconButton
          rounded="false"
          dataTest="link-icon"
          onClick={copyApiUrl}
          title="Copy API URL"
        >
          <LinkIcon />
        </IconButton>
      </Table.BodyCell>
      <Table.BodyCell padding="none" border="left">
        {apiType ? (
          <IconButton
            as={Link}
            to={`/documentation/${organization}/${name}`}
            rounded="false"
            title="Open API documentation"
          >
            <DocsIcon />
          </IconButton>
        ) : (
          <IconButton disabled>
            <DocsIcon />
          </IconButton>
        )}
      </Table.BodyCell>
    </StyledServiceTableRow>
  )
}

ServicesTableRow.propTypes = {
  status: oneOf(['unknown', 'up', 'degraded', 'down']).isRequired,
  organization: string.isRequired,
  name: string.isRequired,
  apiType: string,
}

export default ServicesTableRow
