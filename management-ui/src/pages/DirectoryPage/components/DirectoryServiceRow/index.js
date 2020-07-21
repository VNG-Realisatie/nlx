// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { shape, string, object } from 'prop-types'
import { Table } from '@commonground/design-system'

import StateIndicator from '../../../../components/StateIndicator'
import { AccessRequestContext } from '../../index'
import QuickAccessButton from '../QuickAccessButton'
import AccessRequestMessage from '../AccessRequestMessage'
import { StyledTdAccess } from './index.styles'

const DirectoryServiceRow = ({ service }) => {
  const { handleRequestAccess, requestSentTo } = useContext(
    AccessRequestContext,
  )
  const {
    organizationName,
    serviceName,
    state,
    apiSpecificationType,
    latestAccessRequest,
  } = service

  const requestAccess = (evt) => {
    evt.stopPropagation() // Prevent triggering click on table row
    handleRequestAccess({ organizationName, serviceName })
  }

  const isRequestSentForThisService =
    requestSentTo.organizationName === organizationName &&
    requestSentTo.serviceName === serviceName

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
          <QuickAccessButton
            onClick={requestAccess}
            disabled={isRequestSentForThisService}
          />
        )}
      </StyledTdAccess>
    </Table.Tr>
  )
}

DirectoryServiceRow.propTypes = {
  service: shape({
    organizationName: string.isRequired,
    serviceName: string.isRequired,
    state: string.isRequired,
    apiSpecificationType: string,
    latestAccessRequest: object,
  }),
}

export default DirectoryServiceRow
