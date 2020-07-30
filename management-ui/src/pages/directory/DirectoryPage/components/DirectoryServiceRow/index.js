// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape } from 'prop-types'
import { observer } from 'mobx-react'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import StateIndicator from '../../../../../components/StateIndicator'
import QuickAccessButton from '../QuickAccessButton'
import AccessRequestMessage from '../AccessRequestMessage'
import { StyledTdAccess } from './index.styles'

const DirectoryServiceRow = ({ service }) => {
  const { t } = useTranslation()
  const {
    organizationName,
    serviceName,
    state,
    apiSpecificationType,
    latestAccessRequest,
  } = service

  const requestAccess = (evt) => {
    evt.stopPropagation() // Prevent triggering click on table row

    const confirmed = window.confirm(
      t('The request will be sent to', { name: organizationName }),
    )

    if (confirmed) service.requestAccess()
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
  service: shape(
    pick(directoryServicePropTypes, [
      'organizationName',
      'serviceName',
      'state',
      'apiSpecificationType',
      'latestAccessRequest',
    ]),
  ),
}

export default observer(DirectoryServiceRow)
