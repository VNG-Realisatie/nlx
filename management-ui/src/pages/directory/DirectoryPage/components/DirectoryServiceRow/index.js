// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useMemo } from 'react'
import { shape } from 'prop-types'
import { observer } from 'mobx-react'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import getDirectoryServiceAccessUIState, {
  SHOW_REQUEST_ACCESS,
} from '../../../directoryServiceAccessState'
import StateIndicator from '../../../../../components/StateIndicator'
import QuickAccessButton from '../QuickAccessButton'
import AccessMessage from '../AccessMessage'
import { StyledTdAccess } from './index.styles'

const DirectoryServiceRow = ({ service }) => {
  const { t } = useTranslation()
  const {
    organizationName,
    serviceName,
    state,
    apiSpecificationType,
    latestAccessRequest,
    latestAccessProof,
  } = service

  const handleQuickAccessButtonClick = (event) => {
    event.stopPropagation()
    requestAccess()
  }

  const requestAccess = () => {
    const confirmed = window.confirm(
      t('The request will be sent to', { name: organizationName }),
    )

    if (confirmed) {
      service.requestAccess()
    }
  }

  const displayState = useMemo(
    () =>
      getDirectoryServiceAccessUIState(latestAccessRequest, latestAccessProof),
    [latestAccessRequest, latestAccessProof],
  )

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
        {displayState === SHOW_REQUEST_ACCESS ? (
          <QuickAccessButton onClick={handleQuickAccessButtonClick} />
        ) : (
          <AccessMessage displayState={displayState} />
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
