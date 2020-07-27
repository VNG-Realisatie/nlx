// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { Table } from '@commonground/design-system'

import DirectoryServiceModel from '../../../../../models/DirectoryServiceModel'
import StateIndicator from '../../../../../components/StateIndicator'
import QuickAccessButton from '../QuickAccessButton'
import AccessRequestMessage from '../AccessRequestMessage'
import { StyledTdAccess } from './index.styles'

const DirectoryServiceRow = ({ service }) => {
  const {
    organizationName,
    serviceName,
    state,
    apiSpecificationType,
    latestAccessRequest,
  } = service

  const requestAccess = (evt) => {
    evt.stopPropagation() // Prevent triggering click on table row
    service.requestAccess()
  }

  return (
    <Table.Tr
      to={`/directory/${organizationName}/${serviceName}`}
      name={`${organizationName} - ${serviceName}`}
      data-testid="directory-service-row"
    >
      <Table.Td>{organizationName}</Table.Td>
      <Table.Td>{serviceName}</Table.Td>
      <Table.Td>
        <StateIndicator state={state} />
      </Table.Td>
      <Table.Td>{apiSpecificationType}</Table.Td>
      <StyledTdAccess>
        {latestAccessRequest ? (
          <AccessRequestMessage latestAccessRequest={latestAccessRequest} />
        ) : (
          <QuickAccessButton onClick={requestAccess} />
        )}
      </StyledTdAccess>
    </Table.Tr>
  )
}

DirectoryServiceRow.propTypes = {
  service: instanceOf(DirectoryServiceModel),
}

export default observer(DirectoryServiceRow)
