// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, object, shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { Alert, useDrawerStack } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useConfirmationModal } from '../../../../../components/ConfirmationModal'
import RequestConfirmation from '../../../RequestConfirmation'
import { ACCESS_REQUEST_STATES } from '../../../../../stores/models/OutgoingAccessRequestModel'
import getDirectoryServiceAccessUIState from '../../../directoryServiceAccessState'
import { SectionGroup } from '../../../../../components/DetailView'
import useInterval from '../../../../../hooks/use-interval'
import StacktraceDrawer from './StacktraceDrawer'
import ExternalLinkSection from './ExternalLinkSection'
import AccessSection from './AccessSection'
import ContactSection from './ContactSection'
import { StyledAlert } from './index.styles'

const { FAILED } = ACCESS_REQUEST_STATES
const stackTraceDrawerId = 'stacktrace'

const DirectoryDetailView = ({ service }) => {
  const { t } = useTranslation()
  const { showDrawer } = useDrawerStack()
  const {
    organizationName,
    serviceName,
    latestAccessRequest,
    latestAccessProof,
  } = service

  const [RequestConfirmationModal, confirmRequest] = useConfirmationModal({
    title: t('Request access'),
    okText: t('Send'),
    children: (
      <RequestConfirmation
        organizationName={organizationName}
        serviceName={serviceName}
      />
    ),
  })

  useInterval(service.fetch)

  const requestAccess = async () => {
    const isConfirmed = await confirmRequest()
    if (isConfirmed) service.requestAccess()
  }

  const retryRequestAccess = () => {
    service.retryRequestAccess()
  }

  const displayState = getDirectoryServiceAccessUIState(
    latestAccessRequest,
    latestAccessProof,
  )

  const showTrace = () => {
    showDrawer(stackTraceDrawerId)
  }

  return (
    <>
      {latestAccessRequest && latestAccessRequest.state === FAILED && (
        <StyledAlert
          variant="error"
          title={t('Request could not be sent')}
          actions={[
            <Alert.ButtonAction
              key="send-request-access-action-button"
              onClick={retryRequestAccess}
            >
              {t('Retry')}
            </Alert.ButtonAction>,

            <Alert.ButtonAction
              key="show-trace-access-action-button"
              onClick={showTrace}
            >
              {t('Show stacktrace')}
            </Alert.ButtonAction>,
          ]}
        >
          {latestAccessRequest.errorDetails.cause}
        </StyledAlert>
      )}

      <ExternalLinkSection service={service} />

      <SectionGroup>
        <AccessSection
          displayState={displayState}
          latestAccessRequest={latestAccessRequest}
          latestAccessProof={latestAccessProof}
          requestAccess={requestAccess}
        />

        <ContactSection service={service} />
      </SectionGroup>

      {latestAccessRequest && latestAccessRequest.errorDetails && (
        <StacktraceDrawer
          id={stackTraceDrawerId}
          parentId="directoryDetail"
          stacktrace={latestAccessRequest.errorDetails.stackTrace}
        />
      )}

      <RequestConfirmationModal />
    </>
  )
}

DirectoryDetailView.propTypes = {
  service: shape({
    organizationName: string.isRequired,
    latestAccessRequest: object,
    latestAccessProof: object,
    requestAccess: func.isRequired,
  }),
}

export default observer(DirectoryDetailView)
