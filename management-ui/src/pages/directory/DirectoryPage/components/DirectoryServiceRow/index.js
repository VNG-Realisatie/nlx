// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { object, shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
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
import { AccessMessageWrapper, StyledTdAccess } from './index.styles'

const DirectoryServiceRow = ({ service, ...props }) => {
  const { t } = useTranslation()
  const {
    organization,
    serviceName,
    state,
    apiSpecificationType,
    latestAccessRequest,
    latestAccessProof,
  } = service

  const [RequestConfirmationModal, confirmRequest] = useConfirmationModal({
    title: t('Request access'),
    okText: t('Send'),
    children: (
      <RequestAccessDetails
        organization={organization}
        serviceName={serviceName}
      />
    ),
  })

  const requestAccess = async () => {
    if (await confirmRequest()) {
      service.requestAccess()
    }
  }

  const displayState = getDirectoryServiceAccessUIState(
    latestAccessRequest,
    latestAccessProof,
  )

  const handleQuickAccessButtonClick = (evt) => {
    evt.stopPropagation()

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
      to={`/directory/${organization.serialNumber}/${serviceName}`}
      name={`${organization.name} - ${serviceName}`}
      data-testid="directory-service-row"
      {...props}
    >
      <Table.Td>{organization.name}</Table.Td>
      <Table.Td>{serviceName}</Table.Td>
      <Table.Td>
        <StateIndicator state={state} />
      </Table.Td>
      <Table.Td>{apiSpecificationType}</Table.Td>
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
  service: shape({
    organization: shape({
      serialNumber: string.isRequired,
      name: string.isRequired,
    }).isRequired,
    serviceName: string.isRequired,
    state: string.isRequired,
    apiSpecificationType: string,
    latestAccessRequest: object,
  }),
}

export default observer(DirectoryServiceRow)
