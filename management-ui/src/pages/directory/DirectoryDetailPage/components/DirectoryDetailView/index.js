// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, func } from 'prop-types'
import { observer } from 'mobx-react'
import { Alert, Button, Spinner } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import { ACCESS_REQUEST_STATES } from '../../../../../models/OutgoingAccessRequestModel'
import AccessRequestMessage from '../../../DirectoryPage/components/AccessRequestMessage'
import { SectionGroup } from '../../../../../components/DetailView'
import { IconKey } from '../../../../../icons'
import AccessRequestRepository from '../../../../../domain/access-request-repository'
import { StyledAlert, AccessSection, IconItem, StateItem } from './index.styles'

const { FAILED, RECEIVED } = ACCESS_REQUEST_STATES

const DirectoryDetailView = ({ service, sendAccessRequest }) => {
  const { t } = useTranslation()
  const { organizationName, latestAccessRequest } = service

  const onRequestAccessButtonClick = (event) => {
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

  const sendRequestAccessRequest = () => {
    sendAccessRequest({
      organizationName: service.organizationName,
      serviceName: service.serviceName,
      id: service.latestAccessRequest.id,
    })
  }

  let icon = Spinner
  if (
    latestAccessRequest &&
    [FAILED, RECEIVED].includes(latestAccessRequest.state)
  ) {
    icon = IconKey
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
              onClick={sendRequestAccessRequest}
            >
              {t('Retry')}
            </Alert.ButtonAction>,
          ]}
        />
      )}

      <SectionGroup>
        <AccessSection data-testid="request-access-section">
          {latestAccessRequest ? (
            <>
              <IconItem as={icon} />
              <StateItem>
                <AccessRequestMessage
                  latestAccessRequest={latestAccessRequest}
                  inDetailView
                />
                {icon === Spinner && '...'}
              </StateItem>
            </>
          ) : (
            <>
              <IconItem as={IconKey} />
              <StateItem>{t('You have no access')}</StateItem>
              <Button onClick={onRequestAccessButtonClick}>
                {t('Request Access')}
              </Button>
            </>
          )}
        </AccessSection>
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
  sendAccessRequest: func,
}

DirectoryDetailView.defaultProps = {
  sendAccessRequest: AccessRequestRepository.sendAccessRequest,
}

export default observer(DirectoryDetailView)
