// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { instanceOf, bool } from 'prop-types'
import Table from '../../../../../components/Table'
import getDirectoryServiceAccessUIState from '../../../directoryServiceAccessState'
import StateIndicator from '../../../../../components/StateIndicator'
import AccessMessage from '../AccessMessage'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import { StyledTd, AccessMessageWrapper, StyledTdAccess } from './index.styles'

const DirectoryServiceRow = ({ service, ownService, ...props }) => {
  const displayState = getDirectoryServiceAccessUIState(
    service.latestAccessRequest,
    service.latestAccessProof,
  )

  return (
    <Table.Tr
      to={`/directory/${service.organization.serialNumber}/${service.serviceName}`}
      name={`${service.organization.name} - ${service.serviceName}`}
      data-testid="directory-service-row"
      {...props}
    >
      <StyledTd color={ownService ? '#FFBC2C' : null}>
        {service.organization.name}
      </StyledTd>
      <Table.Td>{service.serviceName}</Table.Td>
      <Table.Td>
        <StateIndicator state={service.state} showText={false} />
      </Table.Td>
      <Table.Td>{service.apiSpecificationType}</Table.Td>
      <StyledTdAccess>
        <AccessMessageWrapper>
          <AccessMessage displayState={displayState} />
        </AccessMessageWrapper>
      </StyledTdAccess>
    </Table.Tr>
  )
}

DirectoryServiceRow.propTypes = {
  service: instanceOf(DirectoryServiceModel),
  selected: bool,
  ownService: bool,
}

export default observer(DirectoryServiceRow)
