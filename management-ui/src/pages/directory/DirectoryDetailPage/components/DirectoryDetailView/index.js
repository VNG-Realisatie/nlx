// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape } from 'prop-types'
import { observer } from 'mobx-react'
import { Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import { ACCESS_REQUEST_STATES } from '../../../../../models/OutgoingAccessRequestModel'
import { SectionGroup } from '../../../../../components/DetailView'

import AccessRequestSection from './AccessRequestSection'
import { StyledAlert } from './index.styles'

const { FAILED } = ACCESS_REQUEST_STATES

const DirectoryDetailView = ({ service }) => {
  const { t } = useTranslation()
  const { organizationName, latestAccessRequest } = service

  const requestAccess = () => {
    const confirmed = window.confirm(
      t('The request will be sent to', { name: organizationName }),
    )

    if (confirmed) {
      service.requestAccess()
    }
  }

  const retryRequestAccess = () => {
    service.retryRequestAccess()
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
          ]}
        />
      )}

      <SectionGroup>
        <AccessRequestSection
          latestAccessRequest={latestAccessRequest}
          requestAccess={requestAccess}
        />
      </SectionGroup>
    </>
  )
}

DirectoryDetailView.propTypes = {
  service: shape(
    pick(directoryServicePropTypes, [
      'organizationName',
      'latestAccessRequest',
      'requestAccess',
    ]),
  ),
}

export default observer(DirectoryDetailView)
