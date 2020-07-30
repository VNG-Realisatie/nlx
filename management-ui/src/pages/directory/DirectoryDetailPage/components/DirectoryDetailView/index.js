// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape } from 'prop-types'
import { observer } from 'mobx-react'
import { Button, Spinner } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import { ACCESS_REQUEST_STATES } from '../../../../../models/OutgoingAccessRequestModel'
import AccessRequestMessage from '../../../DirectoryPage/components/AccessRequestMessage'
import { SectionGroup } from '../../../../../components/DetailView'
import { IconKey } from '../../../../../icons'
import { StyledAlert, AccessSection, IconItem, StateItem } from './index.styles'

const { FAILED, RECEIVED } = ACCESS_REQUEST_STATES

const DirectoryDetailView = ({ service }) => {
  const { t } = useTranslation()
  const { latestAccessRequest } = service

  const requestAccess = () => {
    const confirmed = window.confirm(
      t('The request will be sent to', { name: service.organizationName }),
    )

    if (confirmed) service.requestAccess()
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
      {latestAccessRequest && latestAccessRequest.state === 'FAILED' && (
        <StyledAlert variant="error" title={t('Request could not be sent')} />
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
              <Button onClick={requestAccess}>{t('Request Access')}</Button>
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
}

export default observer(DirectoryDetailView)
