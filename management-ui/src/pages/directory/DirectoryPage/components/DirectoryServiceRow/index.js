// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { object, shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useConfirmationModal } from '../../../../../components/ConfirmationModal'
import RequestAccessDetails from '../../../RequestAccessDetails'
import getDirectoryServiceAccessUIState, {
  SHOW_REQUEST_ACCESS,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_REJECTED,
  SHOW_ACCESS_REVOKED,
} from '../../../directoryServiceAccessState'
import StateIndicator from '../../../../../components/StateIndicator'
import QuickAccessButton from '../QuickAccessButton'
import AccessMessage from '../AccessMessage'
import { StyledTdAccess, AccessMessageWrapper } from './index.styles'

const DirectoryServiceRow = ({ service, ...props }) => {
  const { t } = useTranslation()
  const {
    organizationName,
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
        organizationName={organizationName}
        serviceName={serviceName}
      />
    ),
  })

  const requestAccess = async () => {
    const isConfirmed = await confirmRequest()
    if (isConfirmed) service.requestAccess()
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
      to={`/directory/${organizationName}/${serviceName}`}
      name={`${organizationName} - ${serviceName}`}
      data-testid="directory-service-row"
      {...props}
    >
      <Table.Td>{organizationName}</Table.Td>
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
    organizationName: string.isRequired,
    serviceName: string.isRequired,
    state: string.isRequired,
    apiSpecificationType: string,
    latestAccessRequest: object,
  }),
}

export default observer(DirectoryServiceRow)
