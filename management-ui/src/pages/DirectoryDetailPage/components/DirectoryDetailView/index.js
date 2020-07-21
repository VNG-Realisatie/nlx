// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { shape, string } from 'prop-types'
import { Button, Spinner } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { AccessRequestContext } from '../../../DirectoryPage'
import AccessRequestMessage from '../../../DirectoryPage/components/AccessRequestMessage'
import { SectionGroup } from '../../../../components/DetailView'

import { IconKey } from '../../../../icons'

import { StyledAlert, AccessSection, IconItem, StateItem } from './index.styles'

const DirectoryDetailView = ({
  organizationName,
  serviceName,
  latestAccessRequest,
}) => {
  const { t } = useTranslation()
  const { handleRequestAccess, requestSentTo } = useContext(
    AccessRequestContext,
  )

  const requestAccess = () =>
    handleRequestAccess({ organizationName, serviceName })

  const isRequestSentForThisService =
    requestSentTo.organizationName === organizationName &&
    requestSentTo.serviceName === serviceName

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
              <Button
                onClick={requestAccess}
                disabled={isRequestSentForThisService}
              >
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
  organizationName: string,
  serviceName: string,
  latestAccessRequest: shape({
    id: string,
    state: string,
    createdAt: string,
    updatedAt: string,
  }),
}

export default DirectoryDetailView
