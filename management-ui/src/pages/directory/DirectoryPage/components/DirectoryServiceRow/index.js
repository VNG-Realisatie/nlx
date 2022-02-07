// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { instanceOf, bool } from 'prop-types'
import Table from '../../../../../components/Table'
import { useConfirmationModal } from '../../../../../components/ConfirmationModal'
import RequestAccessDetails from '../../../RequestAccessDetails'
import getDirectoryServiceAccessUIState, {
  SHOW_ACCESS_REVOKED,
  SHOW_REQUEST_ACCESS,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_REJECTED,
} from '../../../directoryServiceAccessState'
import StateIndicator from '../../../../../components/StateIndicator'
import QuickAccessButton from '../QuickAccessButton'
import AccessMessage from '../AccessMessage'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import { StyledTd, AccessMessageWrapper, StyledTdAccess } from './index.styles'

const DirectoryServiceRow = ({ service, ownService, ...props }) => {
  const { t } = useTranslation()

  const [RequestConfirmationModal, confirmRequest] = useConfirmationModal({
    title: t('Request access'),
    okText: t('Send'),
    children: (
      <RequestAccessDetails
        organization={service.organization}
        serviceName={service.serviceName}
      />
    ),
  })

  const requestAccess = async () => {
    if (await confirmRequest()) {
      service.requestAccess()
    }
  }

  const displayState = getDirectoryServiceAccessUIState(
    service.latestAccessRequest,
    service.latestAccessProof,
  )

  const handleQuickAccessButtonClick = (event) => {
    event.stopPropagation()

    if (displayState === SHOW_REQUEST_FAILED) {
      service.retryRequestAccess()
      return
    }

    requestAccess()
  }

  const showRequestAccessButton = [
    SHOW_REQUEST_ACCESS,
    SHOW_REQUEST_FAILED,
    SHOW_REQUEST_REJECTED,
    SHOW_ACCESS_REVOKED,
  ].includes(displayState)

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
          {showRequestAccessButton && (
            <QuickAccessButton onClick={handleQuickAccessButtonClick} />
          )}
        </AccessMessageWrapper>
      </StyledTdAccess>

      <RequestConfirmationModal />
    </Table.Tr>
  )
}

DirectoryServiceRow.propTypes = {
  service: instanceOf(DirectoryServiceModel),
  selected: bool,
  ownService: bool,
}

export default observer(DirectoryServiceRow)
