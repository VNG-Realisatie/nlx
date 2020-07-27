// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { Button, Spinner } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import DirectoryServiceModel from '../../../../../models/DirectoryServiceModel'
import AccessRequestMessage from '../../../DirectoryPage/components/AccessRequestMessage'
import { SectionGroup } from '../../../../../components/DetailView'

import { IconKey } from '../../../../../icons'

import { StyledAlert, AccessSection, IconItem, StateItem } from './index.styles'

const DirectoryDetailView = ({ service }) => {
  const { t } = useTranslation()
  const { latestAccessRequest, requestAccess } = service

  let icon = Spinner
  if (
    latestAccessRequest &&
    ['FAILED', 'RECEIVED'].includes(latestAccessRequest.state)
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
  service: instanceOf(DirectoryServiceModel),
}

export default observer(DirectoryDetailView)
