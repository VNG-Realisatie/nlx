// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { useState } from 'react'
import { oneOf, string } from 'prop-types'
import { Link } from 'react-router-dom'
import copy from 'copy-text-to-clipboard'

import Table from '../Table'
import Tooltip from '../Tooltip/Tooltip'
import StatusIcon from './Icons/StatusIcon/StatusIcon'
import IconButton from '../IconButton/IconButton'
import DocsIcon from './Icons/DocsIcon/DocsIcon'
import LinkIcon from './Icons/LinkIcon/LinkIcon'
import { StyledServiceTableRow, StyledApiTypeLabel } from './ServiceTableRow.styles'

export const apiUrlForService = (organization, name) =>
  `http://{your-outway-address}/${organization}/${name}`


const ServicesTableRow = ({ status, organization, name, apiType, ...props }) => {
  const [isCopiedNotifierVisible, setIsCopiedNotifierVisible] = useState(false);

  const showCopiedNotifier = () => {
    setIsCopiedNotifierVisible(true)

    setTimeout(() => {
      setIsCopiedNotifierVisible(false)
    }, 1500);
  }

  const copyApiUrl = () => {
    copy(apiUrlForService(organization, name))
    showCopiedNotifier()
  }

  return (
    <StyledServiceTableRow status={status} {...props}>
      <Table.BodyCell align="center" padding="none" title={status}><StatusIcon status={status} /></Table.BodyCell>
      <Table.BodyCell>{organization}</Table.BodyCell>
      <Table.BodyCell>{name}</Table.BodyCell>
      <Table.BodyCell align="right">
        {
          apiType &&
          <StyledApiTypeLabel status={status}>{apiType}</StyledApiTypeLabel>
        }
      </Table.BodyCell>
      <Table.BodyCell padding="none" border="left">
        <Tooltip content="Copy API URL">
          <div>
            <Tooltip content="URL copied!" isVisible={isCopiedNotifierVisible}>
              <IconButton rounded="false" dataTest="link-icon"
                onClick={copyApiUrl}
              >
                <LinkIcon />
              </IconButton>
            </Tooltip>
          </div>
        </Tooltip>
      </Table.BodyCell>
      <Table.BodyCell padding="none" border="left">
        {
          apiType ?
            <Tooltip content="Open API documentation">
              <div>
                <IconButton as={Link} to={`/documentation/${organization}/${name}`} rounded="false">
                  <DocsIcon />
                </IconButton>
              </div>
            </Tooltip>
            :
            <IconButton disabled><DocsIcon /></IconButton>
        }
      </Table.BodyCell>
    </StyledServiceTableRow>
  )
}

ServicesTableRow.propTypes = {
  status: oneOf(['online', 'offline']).isRequired,
  organization: string.isRequired,
  name: string.isRequired,
}

export default ServicesTableRow
