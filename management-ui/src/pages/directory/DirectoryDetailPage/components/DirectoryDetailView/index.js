// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, number, object, shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { Alert, useDrawerStack } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useConfirmationModal } from '../../../../../components/ConfirmationModal'
import RequestAccessDetails from '../../../RequestAccessDetails'
import { ACCESS_REQUEST_STATES } from '../../../../../stores/models/OutgoingAccessRequestModel'
import getDirectoryServiceAccessUIState from '../../../directoryServiceAccessState'
import { SectionGroup } from '../../../../../components/DetailView'
import usePolling from '../../../../../hooks/use-polling'
import CostsSection from '../../../../../components/CostsSection'
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
    oneTimeCosts,
    monthlyCosts,
    requestCosts,
  } = service

  const [RequestConfirmationModal, confirmRequest] = useConfirmationModal({
    title: t('Request access'),
    okText: t('Send'),
    children: (
      <RequestAccessDetails
        organizationName={organizationName}
        serviceName={serviceName}
        oneTimeCosts={oneTimeCosts}
        monthlyCosts={monthlyCosts}
        requestCosts={requestCosts}
      />
    ),
  })

  const [pauseFetchPolling, continueFetchPolling] = usePolling(service.fetch)

  const requestAccess = async () => {
    pauseFetchPolling()

    if (await confirmRequest()) {
      service.requestAccess()
    }

    continueFetchPolling()
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
            <Alert.ActionButton
              key="send-request-access-action-button"
              onClick={retryRequestAccess}
            >
              {t('Retry')}
            </Alert.ActionButton>,

            <Alert.ActionButton
              key="show-trace-access-action-button"
              onClick={showTrace}
            >
              {t('Show stacktrace')}
            </Alert.ActionButton>,
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
        <CostsSection
          oneTimeCosts={service.oneTimeCosts}
          monthlyCosts={service.monthlyCosts}
          requestCosts={service.requestCosts}
        />
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
    oneTimeCosts: number,
    monthlyCosts: number,
    requestCosts: number,
  }),
}

export default observer(DirectoryDetailView)
