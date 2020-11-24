// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape } from 'prop-types'
import { observer } from 'mobx-react'
import { Alert, useDrawerStack } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import { ACCESS_REQUEST_STATES } from '../../../../../models/OutgoingAccessRequestModel'
import getDirectoryServiceAccessUIState from '../../../directoryServiceAccessState'
import { SectionGroup } from '../../../../../components/DetailView'
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
  const { organizationName, latestAccessRequest, latestAccessProof } = service

  const requestAccess = () => {
    const confirmed = window.confirm(
      t('The request will be sent to', { name: organizationName }),
    )

    if (!confirmed) {
      return
    }

    service.requestAccess()
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
    </>
  )
}

DirectoryDetailView.propTypes = {
  service: shape(
    pick(directoryServicePropTypes, [
      'organizationName',
      'latestAccessRequest',
      'latestAccessProof',
      'requestAccess',
    ]),
  ),
}

export default observer(DirectoryDetailView)
